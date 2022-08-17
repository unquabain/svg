package svg

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type PathCommand struct {
	Type rune
	Args []float64
}

func (pc PathCommand) String() string {
	b := new(strings.Builder)
	b.WriteRune(pc.Type)
	for _, f := range pc.Args {
		fmt.Fprintf(b, ` %f`, f)
	}
	return b.String()
}

type PathAttr struct {
	Commands []PathCommand
}

func (pa PathAttr) Move(dx, dy float64) {
	pa.Commands = append(pa.Commands, PathCommand{'m', []float64{dx, dy}})
}
func (pa PathAttr) MoveTo(x, y float64) {
	pa.Commands = append(pa.Commands, PathCommand{'M', []float64{x, y}})
}
func (pa PathAttr) Line(dx, dy float64) {
	pa.Commands = append(pa.Commands, PathCommand{'l', []float64{dx, dy}})
}
func (pa PathAttr) LineTo(x, y float64) {
	pa.Commands = append(pa.Commands, PathCommand{'L', []float64{x, y}})
}
func (pa PathAttr) HLine(dx float64) {
	pa.Commands = append(pa.Commands, PathCommand{'h', []float64{dx}})
}
func (pa PathAttr) HLineTo(x float64) {
	pa.Commands = append(pa.Commands, PathCommand{'H', []float64{x}})
}
func (pa PathAttr) VLine(dy float64) {
	pa.Commands = append(pa.Commands, PathCommand{'v', []float64{dy}})
}
func (pa PathAttr) VLineTo(y float64) {
	pa.Commands = append(pa.Commands, PathCommand{'V', []float64{y}})
}
func (pa PathAttr) Close() {
	pa.Commands = append(pa.Commands, PathCommand{'z', []float64{}})
}
func (pa PathAttr) Bezier(dx1, dy1, dx2, dy2, dx, dy float64) {
	pa.Commands = append(pa.Commands, PathCommand{'c', []float64{dx1, dy1, dx2, dy2, dx, dy}})
}
func (pa PathAttr) BezierTo(x1, y1, x2, y2, x, y float64) {
	pa.Commands = append(pa.Commands, PathCommand{'C', []float64{x1, y1, x2, y2, x, y}})
}
func (pa PathAttr) SmoothBezier(dx2, dy2, dx, dy float64) {
	pa.Commands = append(pa.Commands, PathCommand{'s', []float64{dx2, dy2, dx, dy}})
}
func (pa PathAttr) SmoothBezierTo(x2, y2, x, y float64) {
	pa.Commands = append(pa.Commands, PathCommand{'S', []float64{x2, y2, x, y}})
}
func (pa PathAttr) Quadratic(dx1, dy1, dx, dy float64) {
	pa.Commands = append(pa.Commands, PathCommand{'q', []float64{dx1, dy1, dx, dy}})
}
func (pa PathAttr) QuadraticTo(x1, y1, x, y float64) {
	pa.Commands = append(pa.Commands, PathCommand{'Q', []float64{x1, y1, x, y}})
}
func (pa PathAttr) SmoothQuadratic(dx, dy float64) {
	pa.Commands = append(pa.Commands, PathCommand{'t', []float64{dx, dy}})
}
func (pa PathAttr) SmoothQuadraticTo(x, y float64) {
	pa.Commands = append(pa.Commands, PathCommand{'T', []float64{x, y}})
}
func (pa PathAttr) Arc(rx, ry, xAxisRotation, largeArcFlag, sweepFlag, dx, dy float64) {
	pa.Commands = append(pa.Commands, PathCommand{'a', []float64{rx, ry, xAxisRotation, largeArcFlag, sweepFlag, dx, dy}})
}
func (pa PathAttr) ArcTo(rx, ry, xAxisRotation, largeArcFlag, sweepFlag, x, y float64) {
	pa.Commands = append(pa.Commands, PathCommand{'A', []float64{rx, ry, xAxisRotation, largeArcFlag, sweepFlag, x, y}})
}
func (pa PathAttr) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	s := make([]string, len(pa.Commands))
	for i, command := range pa.Commands {
		s[i] = command.String()
	}
	return xml.Attr{
		Name:  name,
		Value: strings.Join(s, ` `),
	}, nil
}

type Path struct {
	XMLName xml.Name `xml:"path"`
	CoreAttributes
	StylingAttributes
	PresentationAttributes
	PathLength int      `xml:"pathLength,attr,omitempty"`
	D          PathAttr `xml:"d,attr"`
}
