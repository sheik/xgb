// Package res is the X client API for the X-Resource extension.
package res

/*
	This file was generated by res.xml on May 10 2012 8:04:32pm EDT.
	This file is automatically generated. Edit at your peril!
*/

import (
	"github.com/BurntSushi/xgb"

	"github.com/BurntSushi/xgb/xproto"
)

// Init must be called before using the X-Resource extension.
func Init(c *xgb.Conn) error {
	reply, err := xproto.QueryExtension(c, 10, "X-Resource").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return xgb.Errorf("No extension named X-Resource could be found on on the server.")
	}

	xgb.ExtLock.Lock()
	c.Extensions["X-Resource"] = reply.MajorOpcode
	for evNum, fun := range xgb.NewExtEventFuncs["X-Resource"] {
		xgb.NewEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range xgb.NewExtErrorFuncs["X-Resource"] {
		xgb.NewErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	xgb.ExtLock.Unlock()

	return nil
}

func init() {
	xgb.NewExtEventFuncs["X-Resource"] = make(map[int]xgb.NewEventFun)
	xgb.NewExtErrorFuncs["X-Resource"] = make(map[int]xgb.NewErrorFun)
}

// Skipping definition for base type 'Card32'

// Skipping definition for base type 'Double'

// Skipping definition for base type 'Bool'

// Skipping definition for base type 'Float'

// Skipping definition for base type 'Card8'

// Skipping definition for base type 'Int16'

// Skipping definition for base type 'Int32'

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Byte'

// Skipping definition for base type 'Int8'

// Skipping definition for base type 'Card16'

// Skipping definition for base type 'Char'

// 'Client' struct definition
// Size: 8
type Client struct {
	ResourceBase uint32
	ResourceMask uint32
}

// Struct read Client
func ClientRead(buf []byte, v *Client) int {
	b := 0

	v.ResourceBase = xgb.Get32(buf[b:])
	b += 4

	v.ResourceMask = xgb.Get32(buf[b:])
	b += 4

	return b
}

// Struct list read Client
func ClientReadList(buf []byte, dest []Client) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Client{}
		b += ClientRead(buf[b:], &dest[i])
	}
	return xgb.Pad(b)
}

// Struct write Client
func (v Client) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	xgb.Put32(buf[b:], v.ResourceBase)
	b += 4

	xgb.Put32(buf[b:], v.ResourceMask)
	b += 4

	return buf
}

// Write struct list Client
func ClientListBytes(buf []byte, list []Client) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}
	return b
}

// 'Type' struct definition
// Size: 8
type Type struct {
	ResourceType xproto.Atom
	Count        uint32
}

// Struct read Type
func TypeRead(buf []byte, v *Type) int {
	b := 0

	v.ResourceType = xproto.Atom(xgb.Get32(buf[b:]))
	b += 4

	v.Count = xgb.Get32(buf[b:])
	b += 4

	return b
}

// Struct list read Type
func TypeReadList(buf []byte, dest []Type) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Type{}
		b += TypeRead(buf[b:], &dest[i])
	}
	return xgb.Pad(b)
}

// Struct write Type
func (v Type) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	xgb.Put32(buf[b:], uint32(v.ResourceType))
	b += 4

	xgb.Put32(buf[b:], v.Count)
	b += 4

	return buf
}

// Write struct list Type
func TypeListBytes(buf []byte, list []Type) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}
	return b
}

// Request QueryVersion
// size: 8
type QueryVersionCookie struct {
	*xgb.Cookie
}

func QueryVersion(c *xgb.Conn, ClientMajor byte, ClientMinor byte) QueryVersionCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryVersionRequest(c, ClientMajor, ClientMinor), cookie)
	return QueryVersionCookie{cookie}
}

func QueryVersionUnchecked(c *xgb.Conn, ClientMajor byte, ClientMinor byte) QueryVersionCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryVersionRequest(c, ClientMajor, ClientMinor), cookie)
	return QueryVersionCookie{cookie}
}

// Request reply for QueryVersion
// size: 12
type QueryVersionReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	ServerMajor uint16
	ServerMinor uint16
}

// Waits and reads reply data from request QueryVersion
func (cook QueryVersionCookie) Reply() (*QueryVersionReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryVersionReply(buf), nil
}

// Read reply into structure from buffer for QueryVersion
func queryVersionReply(buf []byte) *QueryVersionReply {
	v := new(QueryVersionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.ServerMajor = xgb.Get16(buf[b:])
	b += 2

	v.ServerMinor = xgb.Get16(buf[b:])
	b += 2

	return v
}

// Write request to wire for QueryVersion
func queryVersionRequest(c *xgb.Conn, ClientMajor byte, ClientMinor byte) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["X-RESOURCE"]
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = ClientMajor
	b += 1

	buf[b] = ClientMinor
	b += 1

	return buf
}

// Request QueryClients
// size: 4
type QueryClientsCookie struct {
	*xgb.Cookie
}

func QueryClients(c *xgb.Conn) QueryClientsCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryClientsRequest(c), cookie)
	return QueryClientsCookie{cookie}
}

func QueryClientsUnchecked(c *xgb.Conn) QueryClientsCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryClientsRequest(c), cookie)
	return QueryClientsCookie{cookie}
}

// Request reply for QueryClients
// size: (32 + xgb.Pad((int(NumClients) * 8)))
type QueryClientsReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	NumClients uint32
	// padding: 20 bytes
	Clients []Client // size: xgb.Pad((int(NumClients) * 8))
}

// Waits and reads reply data from request QueryClients
func (cook QueryClientsCookie) Reply() (*QueryClientsReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryClientsReply(buf), nil
}

// Read reply into structure from buffer for QueryClients
func queryClientsReply(buf []byte) *QueryClientsReply {
	v := new(QueryClientsReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.NumClients = xgb.Get32(buf[b:])
	b += 4

	b += 20 // padding

	v.Clients = make([]Client, v.NumClients)
	b += ClientReadList(buf[b:], v.Clients)

	return v
}

// Write request to wire for QueryClients
func queryClientsRequest(c *xgb.Conn) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["X-RESOURCE"]
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// Request QueryClientResources
// size: 8
type QueryClientResourcesCookie struct {
	*xgb.Cookie
}

func QueryClientResources(c *xgb.Conn, Xid uint32) QueryClientResourcesCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryClientResourcesRequest(c, Xid), cookie)
	return QueryClientResourcesCookie{cookie}
}

func QueryClientResourcesUnchecked(c *xgb.Conn, Xid uint32) QueryClientResourcesCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryClientResourcesRequest(c, Xid), cookie)
	return QueryClientResourcesCookie{cookie}
}

// Request reply for QueryClientResources
// size: (32 + xgb.Pad((int(NumTypes) * 8)))
type QueryClientResourcesReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	NumTypes uint32
	// padding: 20 bytes
	Types []Type // size: xgb.Pad((int(NumTypes) * 8))
}

// Waits and reads reply data from request QueryClientResources
func (cook QueryClientResourcesCookie) Reply() (*QueryClientResourcesReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryClientResourcesReply(buf), nil
}

// Read reply into structure from buffer for QueryClientResources
func queryClientResourcesReply(buf []byte) *QueryClientResourcesReply {
	v := new(QueryClientResourcesReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.NumTypes = xgb.Get32(buf[b:])
	b += 4

	b += 20 // padding

	v.Types = make([]Type, v.NumTypes)
	b += TypeReadList(buf[b:], v.Types)

	return v
}

// Write request to wire for QueryClientResources
func queryClientResourcesRequest(c *xgb.Conn, Xid uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["X-RESOURCE"]
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], Xid)
	b += 4

	return buf
}

// Request QueryClientPixmapBytes
// size: 8
type QueryClientPixmapBytesCookie struct {
	*xgb.Cookie
}

func QueryClientPixmapBytes(c *xgb.Conn, Xid uint32) QueryClientPixmapBytesCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryClientPixmapBytesRequest(c, Xid), cookie)
	return QueryClientPixmapBytesCookie{cookie}
}

func QueryClientPixmapBytesUnchecked(c *xgb.Conn, Xid uint32) QueryClientPixmapBytesCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryClientPixmapBytesRequest(c, Xid), cookie)
	return QueryClientPixmapBytesCookie{cookie}
}

// Request reply for QueryClientPixmapBytes
// size: 16
type QueryClientPixmapBytesReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	Bytes         uint32
	BytesOverflow uint32
}

// Waits and reads reply data from request QueryClientPixmapBytes
func (cook QueryClientPixmapBytesCookie) Reply() (*QueryClientPixmapBytesReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryClientPixmapBytesReply(buf), nil
}

// Read reply into structure from buffer for QueryClientPixmapBytes
func queryClientPixmapBytesReply(buf []byte) *QueryClientPixmapBytesReply {
	v := new(QueryClientPixmapBytesReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.Bytes = xgb.Get32(buf[b:])
	b += 4

	v.BytesOverflow = xgb.Get32(buf[b:])
	b += 4

	return v
}

// Write request to wire for QueryClientPixmapBytes
func queryClientPixmapBytesRequest(c *xgb.Conn, Xid uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["X-RESOURCE"]
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], Xid)
	b += 4

	return buf
}