// Package dpms is the X client API for the DPMS extension.
package dpms

/*
	This file was generated by dpms.xml on May 10 2012 8:04:31pm EDT.
	This file is automatically generated. Edit at your peril!
*/

import (
	"github.com/BurntSushi/xgb"

	"github.com/BurntSushi/xgb/xproto"
)

// Init must be called before using the DPMS extension.
func Init(c *xgb.Conn) error {
	reply, err := xproto.QueryExtension(c, 4, "DPMS").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return xgb.Errorf("No extension named DPMS could be found on on the server.")
	}

	xgb.ExtLock.Lock()
	c.Extensions["DPMS"] = reply.MajorOpcode
	for evNum, fun := range xgb.NewExtEventFuncs["DPMS"] {
		xgb.NewEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range xgb.NewExtErrorFuncs["DPMS"] {
		xgb.NewErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	xgb.ExtLock.Unlock()

	return nil
}

func init() {
	xgb.NewExtEventFuncs["DPMS"] = make(map[int]xgb.NewEventFun)
	xgb.NewExtErrorFuncs["DPMS"] = make(map[int]xgb.NewErrorFun)
}

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

// Skipping definition for base type 'Int32'

// Skipping definition for base type 'Void'

const (
	DPMSModeOn      = 0
	DPMSModeStandby = 1
	DPMSModeSuspend = 2
	DPMSModeOff     = 3
)

// Request GetVersion
// size: 8
type GetVersionCookie struct {
	*xgb.Cookie
}

func GetVersion(c *xgb.Conn, ClientMajorVersion uint16, ClientMinorVersion uint16) GetVersionCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(getVersionRequest(c, ClientMajorVersion, ClientMinorVersion), cookie)
	return GetVersionCookie{cookie}
}

func GetVersionUnchecked(c *xgb.Conn, ClientMajorVersion uint16, ClientMinorVersion uint16) GetVersionCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(getVersionRequest(c, ClientMajorVersion, ClientMinorVersion), cookie)
	return GetVersionCookie{cookie}
}

// Request reply for GetVersion
// size: 12
type GetVersionReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	ServerMajorVersion uint16
	ServerMinorVersion uint16
}

// Waits and reads reply data from request GetVersion
func (cook GetVersionCookie) Reply() (*GetVersionReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return getVersionReply(buf), nil
}

// Read reply into structure from buffer for GetVersion
func getVersionReply(buf []byte) *GetVersionReply {
	v := new(GetVersionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.ServerMajorVersion = xgb.Get16(buf[b:])
	b += 2

	v.ServerMinorVersion = xgb.Get16(buf[b:])
	b += 2

	return v
}

// Write request to wire for GetVersion
func getVersionRequest(c *xgb.Conn, ClientMajorVersion uint16, ClientMinorVersion uint16) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DPMS"]
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put16(buf[b:], ClientMajorVersion)
	b += 2

	xgb.Put16(buf[b:], ClientMinorVersion)
	b += 2

	return buf
}

// Request Capable
// size: 4
type CapableCookie struct {
	*xgb.Cookie
}

func Capable(c *xgb.Conn) CapableCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(capableRequest(c), cookie)
	return CapableCookie{cookie}
}

func CapableUnchecked(c *xgb.Conn) CapableCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(capableRequest(c), cookie)
	return CapableCookie{cookie}
}

// Request reply for Capable
// size: 32
type CapableReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	Capable bool
	// padding: 23 bytes
}

// Waits and reads reply data from request Capable
func (cook CapableCookie) Reply() (*CapableReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return capableReply(buf), nil
}

// Read reply into structure from buffer for Capable
func capableReply(buf []byte) *CapableReply {
	v := new(CapableReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	if buf[b] == 1 {
		v.Capable = true
	} else {
		v.Capable = false
	}
	b += 1

	b += 23 // padding

	return v
}

// Write request to wire for Capable
func capableRequest(c *xgb.Conn) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DPMS"]
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// Request GetTimeouts
// size: 4
type GetTimeoutsCookie struct {
	*xgb.Cookie
}

func GetTimeouts(c *xgb.Conn) GetTimeoutsCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(getTimeoutsRequest(c), cookie)
	return GetTimeoutsCookie{cookie}
}

func GetTimeoutsUnchecked(c *xgb.Conn) GetTimeoutsCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(getTimeoutsRequest(c), cookie)
	return GetTimeoutsCookie{cookie}
}

// Request reply for GetTimeouts
// size: 32
type GetTimeoutsReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	StandbyTimeout uint16
	SuspendTimeout uint16
	OffTimeout     uint16
	// padding: 18 bytes
}

// Waits and reads reply data from request GetTimeouts
func (cook GetTimeoutsCookie) Reply() (*GetTimeoutsReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return getTimeoutsReply(buf), nil
}

// Read reply into structure from buffer for GetTimeouts
func getTimeoutsReply(buf []byte) *GetTimeoutsReply {
	v := new(GetTimeoutsReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.StandbyTimeout = xgb.Get16(buf[b:])
	b += 2

	v.SuspendTimeout = xgb.Get16(buf[b:])
	b += 2

	v.OffTimeout = xgb.Get16(buf[b:])
	b += 2

	b += 18 // padding

	return v
}

// Write request to wire for GetTimeouts
func getTimeoutsRequest(c *xgb.Conn) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DPMS"]
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// Request SetTimeouts
// size: 12
type SetTimeoutsCookie struct {
	*xgb.Cookie
}

// Write request to wire for SetTimeouts
func SetTimeouts(c *xgb.Conn, StandbyTimeout uint16, SuspendTimeout uint16, OffTimeout uint16) SetTimeoutsCookie {
	cookie := c.NewCookie(false, false)
	c.NewRequest(setTimeoutsRequest(c, StandbyTimeout, SuspendTimeout, OffTimeout), cookie)
	return SetTimeoutsCookie{cookie}
}

func SetTimeoutsChecked(c *xgb.Conn, StandbyTimeout uint16, SuspendTimeout uint16, OffTimeout uint16) SetTimeoutsCookie {
	cookie := c.NewCookie(true, false)
	c.NewRequest(setTimeoutsRequest(c, StandbyTimeout, SuspendTimeout, OffTimeout), cookie)
	return SetTimeoutsCookie{cookie}
}

func (cook SetTimeoutsCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for SetTimeouts
func setTimeoutsRequest(c *xgb.Conn, StandbyTimeout uint16, SuspendTimeout uint16, OffTimeout uint16) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DPMS"]
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put16(buf[b:], StandbyTimeout)
	b += 2

	xgb.Put16(buf[b:], SuspendTimeout)
	b += 2

	xgb.Put16(buf[b:], OffTimeout)
	b += 2

	return buf
}

// Request Enable
// size: 4
type EnableCookie struct {
	*xgb.Cookie
}

// Write request to wire for Enable
func Enable(c *xgb.Conn) EnableCookie {
	cookie := c.NewCookie(false, false)
	c.NewRequest(enableRequest(c), cookie)
	return EnableCookie{cookie}
}

func EnableChecked(c *xgb.Conn) EnableCookie {
	cookie := c.NewCookie(true, false)
	c.NewRequest(enableRequest(c), cookie)
	return EnableCookie{cookie}
}

func (cook EnableCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for Enable
func enableRequest(c *xgb.Conn) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DPMS"]
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// Request Disable
// size: 4
type DisableCookie struct {
	*xgb.Cookie
}

// Write request to wire for Disable
func Disable(c *xgb.Conn) DisableCookie {
	cookie := c.NewCookie(false, false)
	c.NewRequest(disableRequest(c), cookie)
	return DisableCookie{cookie}
}

func DisableChecked(c *xgb.Conn) DisableCookie {
	cookie := c.NewCookie(true, false)
	c.NewRequest(disableRequest(c), cookie)
	return DisableCookie{cookie}
}

func (cook DisableCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for Disable
func disableRequest(c *xgb.Conn) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DPMS"]
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// Request ForceLevel
// size: 8
type ForceLevelCookie struct {
	*xgb.Cookie
}

// Write request to wire for ForceLevel
func ForceLevel(c *xgb.Conn, PowerLevel uint16) ForceLevelCookie {
	cookie := c.NewCookie(false, false)
	c.NewRequest(forceLevelRequest(c, PowerLevel), cookie)
	return ForceLevelCookie{cookie}
}

func ForceLevelChecked(c *xgb.Conn, PowerLevel uint16) ForceLevelCookie {
	cookie := c.NewCookie(true, false)
	c.NewRequest(forceLevelRequest(c, PowerLevel), cookie)
	return ForceLevelCookie{cookie}
}

func (cook ForceLevelCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for ForceLevel
func forceLevelRequest(c *xgb.Conn, PowerLevel uint16) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DPMS"]
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put16(buf[b:], PowerLevel)
	b += 2

	return buf
}

// Request Info
// size: 4
type InfoCookie struct {
	*xgb.Cookie
}

func Info(c *xgb.Conn) InfoCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(infoRequest(c), cookie)
	return InfoCookie{cookie}
}

func InfoUnchecked(c *xgb.Conn) InfoCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(infoRequest(c), cookie)
	return InfoCookie{cookie}
}

// Request reply for Info
// size: 32
type InfoReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	PowerLevel uint16
	State      bool
	// padding: 21 bytes
}

// Waits and reads reply data from request Info
func (cook InfoCookie) Reply() (*InfoReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return infoReply(buf), nil
}

// Read reply into structure from buffer for Info
func infoReply(buf []byte) *InfoReply {
	v := new(InfoReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.PowerLevel = xgb.Get16(buf[b:])
	b += 2

	if buf[b] == 1 {
		v.State = true
	} else {
		v.State = false
	}
	b += 1

	b += 21 // padding

	return v
}

// Write request to wire for Info
func infoRequest(c *xgb.Conn) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DPMS"]
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}