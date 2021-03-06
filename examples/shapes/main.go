// Example shapes shows how to draw basic shapes into a window.
// It can be considered the Go aequivalent of
// https://x.org/releases/X11R7.5/doc/libxcb/tutorial/#drawingprim
// Four points, a single polyline, two line segments,
// two rectangle and two arcs are drawn.
package main

import (
	"fmt"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

func main() {
	X, err := xgb.NewConn()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer X.Close()

	setup := xproto.Setup(X)
	screen := setup.DefaultScreen(X)
	wid, _ := xproto.NewWindowId(X)
	draw := xproto.Drawable(wid)

	// Create the window
	xproto.CreateWindow(X, screen.RootDepth, wid, screen.Root,
		0, 0, 500, 500, 8, // X, Y, width, height, *border width*
		xproto.WindowClassInputOutput, screen.RootVisual,
		xproto.CwBackPixel | xproto.CwEventMask,
		[]uint32{ screen.WhitePixel, xproto.EventMaskStructureNotify | xproto.EventMaskExposure })

	// Map the window on the screen
	xproto.MapWindow(X, wid)

	// Up to here everything is the same as in the `create-window` example.
	// We opened a connection, created and mapped the window.
	// Note how this time the border width is set to 8 instead of 0.
	//
	// But this time we'll be drawing some basic shapes:

	// First of all we need to create a context to draw with.
	// The graphics context combines all properties (e.g. color, line width, font, fill style, ...)
	// that should be used to draw something. All available properties
	//
	// These properties can be set by or'ing their keys (xproto.Gc*)
	// and adding the value to the end of the values array.
	// The order in which the values have to be given corresponds to the order that they defined
	// mentioned in `xproto`.
	//
	// Here we create a new graphics context
	// which only has the foreground (color) value set to black:
	foreground, _ := xproto.NewGcontextId(X)
	mask := uint32(xproto.GcForeground)
	values := []uint32{ screen.BlackPixel }
	xproto.CreateGC(X, foreground, draw, mask, values)

	// It is possible to set the foreground value to something different.
	// In production, this should use xorg color maps instead for compatibility
	// but for demonstration setting the color directly also works:
	red, _ := xproto.NewGcontextId(X)
	mask = uint32(xproto.GcForeground)
	values = []uint32{ 0xff0000 }
	xproto.CreateGC(X, red, draw, mask, values)

	// We'll create another graphics context that draws thick lines:
	thick, _ := xproto.NewGcontextId(X)
	mask = uint32(xproto.GcLineWidth)
	values = []uint32{ 10 }
	xproto.CreateGC(X, thick, draw, mask, values)


	points := []xproto.Point{
		{X: 10, Y: 10},
		{X: 20, Y: 10},
		{X: 30, Y: 10},
		{X: 40, Y: 10},
	}

	// A polyline is essientially a line with multiple points.
	// The first point is placed absolutely on the screen.
	// The other points are relative to the one before them
	polyline := []xproto.Point{
		{X: 50, Y: 10},
		{X: 5, Y: 20}, // move 5 to the right, 20 down
		{X: 25, Y: -20}, // move 25 to the right, 20 up - notice how this point is level again with the first point
		{X: 10, Y: 10}, // move 10 to the right, 10 down
	}

	segments := []xproto.Segment{
		{X1: 100, Y1: 10, X2: 140, Y2: 30},
		{X1: 110, Y1: 25, X2: 130, Y2: 60},
	}

	// Rectangles have a start coordinate (upper left) and width and height.
	rectangles := []xproto.Rectangle{
		{X: 10, Y: 50, Width: 40, Height: 20},
		{X: 80, Y: 50, Width: 10, Height: 40},
	}

	arcs := []xproto.Arc{
		{X: 10, Y: 100, Width: 60, Height: 40, Angle1: 0, Angle2: 90 << 6},
		{X: 90, Y: 100, Width: 55, Height: 40, Angle1: 0, Angle2: 270 << 6},
	}

	for {
		evt, err := X.WaitForEvent()
		switch evt.(type) {
		case xproto.ExposeEvent:
			/* We draw the points */
			xproto.PolyPoint(X, xproto.CoordModeOrigin, draw, foreground, points)

			/* We draw the polygonal line */
			xproto.PolyLine(X, xproto.CoordModePrevious, draw, red, polyline)

			/* We draw the segments */
			xproto.PolySegment(X, draw, thick, segments)

			/* We draw the rectangles */
			xproto.PolyRectangle(X, draw, red, rectangles)

			/* We draw the arcs */
			xproto.PolyArc(X, draw, foreground, arcs)

		case xproto.DestroyNotifyEvent:
			return
		}

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
