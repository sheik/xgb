package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"sort"
)

// Context represents the protocol we're converting to Go, and a writer
// buffer to write the Go source to.
type Context struct {
	protocol *Protocol
	out      *bytes.Buffer
}

func newContext() *Context {
	return &Context{
		out: bytes.NewBuffer([]byte{}),
	}
}

// Putln calls put and adds a new line to the end of 'format'.
func (c *Context) Putln(format string, v ...interface{}) {
	c.Put(format+"\n", v...)
}

// Put is a short alias to write to 'out'.
func (c *Context) Put(format string, v ...interface{}) {
	_, err := c.out.WriteString(fmt.Sprintf(format, v...))
	if err != nil {
		log.Fatalf("There was an error writing to context buffer: %s", err)
	}
}

// Morph is the big daddy of them all. It takes in an XML byte slice,
// parse it, transforms the XML types into more usable types,
// and writes Go code to the 'out' buffer.
func (c *Context) Morph(xmlBytes []byte) {
	parsedXml := &XML{}
	err := xml.Unmarshal(xmlBytes, parsedXml)
	if err != nil {
		log.Fatal(err)
	}

	// Parse all imports
	parsedXml.Imports.Eval()

	// Translate XML types to nice types
	c.protocol = parsedXml.Translate(nil)
	
	// For backwards compatibility we patch the type of the send_event field of
	// PutImage to be byte
	if c.protocol.Name == "shm" {
		for _, req := range c.protocol.Requests {
			if req.xmlName != "PutImage" {
				continue
			}
			for _, ifield := range req.Fields {
				field, ok := ifield.(*SingleField)
				if !ok || field.xmlName != "send_event" {
					continue
				}
				field.Type = &Base{ srcName: "byte", xmlName: "CARD8", size: newFixedSize(1, true) }
			}
		}
	}

	// Start with Go header.
	c.Putln("// Package %s is the X client API for the %s extension.",
		c.protocol.PkgName(), c.protocol.ExtXName)
	c.Putln("package %s", c.protocol.PkgName())
	c.Putln("")
	c.Putln("// This file is automatically generated from %s.xml. "+
		"Edit at your peril!", c.protocol.Name)
	c.Putln("")

	// Write imports. We always need to import at least xgb.
	// We also need to import xproto if it's an extension.
	c.Putln("import (")
	c.Putln("\"github.com/sheik/xgb\"")
	c.Putln("")
	if c.protocol.isExt() {
		c.Putln("\"github.com/sheik/xgb/xproto\"")
	}

	sort.Sort(Protocols(c.protocol.Imports))
	for _, imp := range c.protocol.Imports {
		// We always import xproto, so skip it if it's explicitly imported
		if imp.Name == "xproto" {
			continue
		}
		c.Putln("\"github.com/sheik/xgb/%s\"", imp.Name)
	}
	c.Putln(")")
	c.Putln("")

	// If this is an extension, create a function to initialize the extension
	// before it can be used.
	if c.protocol.isExt() {
		xname := c.protocol.ExtXName

		c.Putln("// Init must be called before using the %s extension.",
			xname)
		c.Putln("func Init(c *xgb.Conn) error {")
		c.Putln("reply, err := xproto.QueryExtension(c, %d, \"%s\").Reply()",
			len(xname), xname)
		c.Putln("switch {")
		c.Putln("case err != nil:")
		c.Putln("return err")
		c.Putln("case !reply.Present:")
		c.Putln("return xgb.Errorf(\"No extension named %s could be found on "+
			"on the server.\")", xname)
		c.Putln("}")
		c.Putln("")
		c.Putln("c.ExtLock.Lock()")
		c.Putln("c.Extensions[\"%s\"] = reply.MajorOpcode", xname)
		c.Putln("c.ExtLock.Unlock()")
		c.Putln("for evNum, fun := range xgb.NewExtEventFuncs[\"%s\"] {",
			xname)
		c.Putln("xgb.NewEventFuncs[int(reply.FirstEvent) + evNum] = fun")
		c.Putln("}")
		c.Putln("for errNum, fun := range xgb.NewExtErrorFuncs[\"%s\"] {",
			xname)
		c.Putln("xgb.NewErrorFuncs[int(reply.FirstError) + errNum] = fun")
		c.Putln("}")
		c.Putln("return nil")
		c.Putln("}")
		c.Putln("")

		// Make sure newExtEventFuncs["EXT_NAME"] map is initialized.
		// Same deal for newExtErrorFuncs["EXT_NAME"]
		c.Putln("func init() {")
		c.Putln("xgb.NewExtEventFuncs[\"%s\"] = make(map[int]xgb.NewEventFun)",
			xname)
		c.Putln("xgb.NewExtErrorFuncs[\"%s\"] = make(map[int]xgb.NewErrorFun)",
			xname)
		c.Putln("}")
		c.Putln("")
	} else {
		// In the xproto package, we must provide a Setup function that uses
		// SetupBytes in xgb.Conn to return a SetupInfo structure.
		c.Putln("// Setup parses the setup bytes retrieved when")
		c.Putln("// connecting into a SetupInfo struct.")
		c.Putln("func Setup(c *xgb.Conn) *SetupInfo {")
		c.Putln("setup := new(SetupInfo)")
		c.Putln("SetupInfoRead(c.SetupBytes, setup)")
		c.Putln("return setup")
		c.Putln("}")
		c.Putln("")
		c.Putln("// DefaultScreen gets the default screen info from SetupInfo.")
		c.Putln("func (s *SetupInfo) DefaultScreen(c *xgb.Conn) *ScreenInfo {")
		c.Putln("return &s.Roots[c.DefaultScreen]")
		c.Putln("}")
		c.Putln("")
	}

	// Now write Go source code
	sort.Sort(Types(c.protocol.Types))
	sort.Sort(Requests(c.protocol.Requests))
	for _, typ := range c.protocol.Types {
		typ.Define(c)
	}
	for _, req := range c.protocol.Requests {
		req.Define(c)
	}
}
