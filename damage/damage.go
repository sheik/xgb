// Package damage is the X client API for the DAMAGE extension.
package damage

/*
	This file was generated by damage.xml on May 10 2012 8:04:31pm EDT.
	This file is automatically generated. Edit at your peril!
*/

import (
	"github.com/BurntSushi/xgb"

	"github.com/BurntSushi/xgb/xfixes"
	"github.com/BurntSushi/xgb/xproto"
)

// Init must be called before using the DAMAGE extension.
func Init(c *xgb.Conn) error {
	reply, err := xproto.QueryExtension(c, 6, "DAMAGE").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return xgb.Errorf("No extension named DAMAGE could be found on on the server.")
	}

	xgb.ExtLock.Lock()
	c.Extensions["DAMAGE"] = reply.MajorOpcode
	for evNum, fun := range xgb.NewExtEventFuncs["DAMAGE"] {
		xgb.NewEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range xgb.NewExtErrorFuncs["DAMAGE"] {
		xgb.NewErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	xgb.ExtLock.Unlock()

	return nil
}

func init() {
	xgb.NewExtEventFuncs["DAMAGE"] = make(map[int]xgb.NewEventFun)
	xgb.NewExtErrorFuncs["DAMAGE"] = make(map[int]xgb.NewErrorFun)
}

// Skipping definition for base type 'Int32'

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Byte'

// Skipping definition for base type 'Int8'

// Skipping definition for base type 'Card16'

// Skipping definition for base type 'Char'

// Skipping definition for base type 'Card32'

// Skipping definition for base type 'Double'

// Skipping definition for base type 'Bool'

// Skipping definition for base type 'Float'

// Skipping definition for base type 'Card8'

// Skipping definition for base type 'Int16'

const (
	ReportLevelRawRectangles   = 0
	ReportLevelDeltaRectangles = 1
	ReportLevelBoundingBox     = 2
	ReportLevelNonEmpty        = 3
)

type Damage uint32

func NewDamageId(c *xgb.Conn) (Damage, error) {
	id, err := c.NewId()
	if err != nil {
		return 0, err
	}
	return Damage(id), nil
}

// Event definition Notify (0)
// Size: 32

const Notify = 0

type NotifyEvent struct {
	Sequence  uint16
	Level     byte
	Drawable  xproto.Drawable
	Damage    Damage
	Timestamp xproto.Timestamp
	Area      xproto.Rectangle
	Geometry  xproto.Rectangle
}

// Event read Notify
func NotifyEventNew(buf []byte) xgb.Event {
	v := NotifyEvent{}
	b := 1 // don't read event number

	v.Level = buf[b]
	b += 1

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Drawable = xproto.Drawable(xgb.Get32(buf[b:]))
	b += 4

	v.Damage = Damage(xgb.Get32(buf[b:]))
	b += 4

	v.Timestamp = xproto.Timestamp(xgb.Get32(buf[b:]))
	b += 4

	v.Area = xproto.Rectangle{}
	b += xproto.RectangleRead(buf[b:], &v.Area)

	v.Geometry = xproto.Rectangle{}
	b += xproto.RectangleRead(buf[b:], &v.Geometry)

	return v
}

// Event write Notify
func (v NotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 0
	b += 1

	buf[b] = v.Level
	b += 1

	b += 2 // skip sequence number

	xgb.Put32(buf[b:], uint32(v.Drawable))
	b += 4

	xgb.Put32(buf[b:], uint32(v.Damage))
	b += 4

	xgb.Put32(buf[b:], uint32(v.Timestamp))
	b += 4

	{
		structBytes := v.Area.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}

	{
		structBytes := v.Geometry.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}

	return buf
}

func (v NotifyEvent) ImplementsEvent() {}

func (v NotifyEvent) SequenceId() uint16 {
	return v.Sequence
}

func (v NotifyEvent) String() string {
	fieldVals := make([]string, 0, 6)
	fieldVals = append(fieldVals, xgb.Sprintf("Sequence: %d", v.Sequence))
	fieldVals = append(fieldVals, xgb.Sprintf("Level: %d", v.Level))
	fieldVals = append(fieldVals, xgb.Sprintf("Drawable: %d", v.Drawable))
	fieldVals = append(fieldVals, xgb.Sprintf("Damage: %d", v.Damage))
	fieldVals = append(fieldVals, xgb.Sprintf("Timestamp: %d", v.Timestamp))
	return "Notify {" + xgb.StringsJoin(fieldVals, ", ") + "}"
}

func init() {
	xgb.NewExtEventFuncs["DAMAGE"][0] = NotifyEventNew
}

// Error definition BadDamage (0)
// Size: 32

const BadBadDamage = 0

type BadDamageError struct {
	Sequence uint16
	NiceName string
}

// Error read BadDamage
func BadDamageErrorNew(buf []byte) xgb.Error {
	v := BadDamageError{}
	v.NiceName = "BadDamage"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	return v
}

func (err BadDamageError) ImplementsError() {}

func (err BadDamageError) SequenceId() uint16 {
	return err.Sequence
}

func (err BadDamageError) BadId() uint32 {
	return 0
}

func (err BadDamageError) Error() string {
	fieldVals := make([]string, 0, 0)
	fieldVals = append(fieldVals, "NiceName: "+err.NiceName)
	fieldVals = append(fieldVals, xgb.Sprintf("Sequence: %d", err.Sequence))
	return "BadBadDamage {" + xgb.StringsJoin(fieldVals, ", ") + "}"
}

func init() {
	xgb.NewExtErrorFuncs["DAMAGE"][0] = BadDamageErrorNew
}

// Request QueryVersion
// size: 12
type QueryVersionCookie struct {
	*xgb.Cookie
}

func QueryVersion(c *xgb.Conn, ClientMajorVersion uint32, ClientMinorVersion uint32) QueryVersionCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryVersionRequest(c, ClientMajorVersion, ClientMinorVersion), cookie)
	return QueryVersionCookie{cookie}
}

func QueryVersionUnchecked(c *xgb.Conn, ClientMajorVersion uint32, ClientMinorVersion uint32) QueryVersionCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryVersionRequest(c, ClientMajorVersion, ClientMinorVersion), cookie)
	return QueryVersionCookie{cookie}
}

// Request reply for QueryVersion
// size: 32
type QueryVersionReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	MajorVersion uint32
	MinorVersion uint32
	// padding: 16 bytes
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

	v.MajorVersion = xgb.Get32(buf[b:])
	b += 4

	v.MinorVersion = xgb.Get32(buf[b:])
	b += 4

	b += 16 // padding

	return v
}

// Write request to wire for QueryVersion
func queryVersionRequest(c *xgb.Conn, ClientMajorVersion uint32, ClientMinorVersion uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DAMAGE"]
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], ClientMajorVersion)
	b += 4

	xgb.Put32(buf[b:], ClientMinorVersion)
	b += 4

	return buf
}

// Request Create
// size: 16
type CreateCookie struct {
	*xgb.Cookie
}

// Write request to wire for Create
func Create(c *xgb.Conn, Damage Damage, Drawable xproto.Drawable, Level byte) CreateCookie {
	cookie := c.NewCookie(false, false)
	c.NewRequest(createRequest(c, Damage, Drawable, Level), cookie)
	return CreateCookie{cookie}
}

func CreateChecked(c *xgb.Conn, Damage Damage, Drawable xproto.Drawable, Level byte) CreateCookie {
	cookie := c.NewCookie(true, false)
	c.NewRequest(createRequest(c, Damage, Drawable, Level), cookie)
	return CreateCookie{cookie}
}

func (cook CreateCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for Create
func createRequest(c *xgb.Conn, Damage Damage, Drawable xproto.Drawable, Level byte) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DAMAGE"]
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Damage))
	b += 4

	xgb.Put32(buf[b:], uint32(Drawable))
	b += 4

	buf[b] = Level
	b += 1

	b += 3 // padding

	return buf
}

// Request Destroy
// size: 8
type DestroyCookie struct {
	*xgb.Cookie
}

// Write request to wire for Destroy
func Destroy(c *xgb.Conn, Damage Damage) DestroyCookie {
	cookie := c.NewCookie(false, false)
	c.NewRequest(destroyRequest(c, Damage), cookie)
	return DestroyCookie{cookie}
}

func DestroyChecked(c *xgb.Conn, Damage Damage) DestroyCookie {
	cookie := c.NewCookie(true, false)
	c.NewRequest(destroyRequest(c, Damage), cookie)
	return DestroyCookie{cookie}
}

func (cook DestroyCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for Destroy
func destroyRequest(c *xgb.Conn, Damage Damage) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DAMAGE"]
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Damage))
	b += 4

	return buf
}

// Request Subtract
// size: 16
type SubtractCookie struct {
	*xgb.Cookie
}

// Write request to wire for Subtract
func Subtract(c *xgb.Conn, Damage Damage, Repair xfixes.Region, Parts xfixes.Region) SubtractCookie {
	cookie := c.NewCookie(false, false)
	c.NewRequest(subtractRequest(c, Damage, Repair, Parts), cookie)
	return SubtractCookie{cookie}
}

func SubtractChecked(c *xgb.Conn, Damage Damage, Repair xfixes.Region, Parts xfixes.Region) SubtractCookie {
	cookie := c.NewCookie(true, false)
	c.NewRequest(subtractRequest(c, Damage, Repair, Parts), cookie)
	return SubtractCookie{cookie}
}

func (cook SubtractCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for Subtract
func subtractRequest(c *xgb.Conn, Damage Damage, Repair xfixes.Region, Parts xfixes.Region) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DAMAGE"]
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Damage))
	b += 4

	xgb.Put32(buf[b:], uint32(Repair))
	b += 4

	xgb.Put32(buf[b:], uint32(Parts))
	b += 4

	return buf
}

// Request Add
// size: 12
type AddCookie struct {
	*xgb.Cookie
}

// Write request to wire for Add
func Add(c *xgb.Conn, Drawable xproto.Drawable, Region xfixes.Region) AddCookie {
	cookie := c.NewCookie(false, false)
	c.NewRequest(addRequest(c, Drawable, Region), cookie)
	return AddCookie{cookie}
}

func AddChecked(c *xgb.Conn, Drawable xproto.Drawable, Region xfixes.Region) AddCookie {
	cookie := c.NewCookie(true, false)
	c.NewRequest(addRequest(c, Drawable, Region), cookie)
	return AddCookie{cookie}
}

func (cook AddCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for Add
func addRequest(c *xgb.Conn, Drawable xproto.Drawable, Region xfixes.Region) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DAMAGE"]
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Drawable))
	b += 4

	xgb.Put32(buf[b:], uint32(Region))
	b += 4

	return buf
}