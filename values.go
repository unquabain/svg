package svg

import (
	"encoding/xml"
	"fmt"
	"image/color"
	"strings"

	"golang.org/x/image/colornames"
)

type Percentage float64

func (p Percentage) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if p == 0 {
		return xml.Attr{}, nil
	}
	return xml.Attr{
		Name:  name,
		Value: fmt.Sprintf(`%f%%`, p*100),
	}, nil
}

type URL string

func (u URL) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if u == `` {
		return xml.Attr{}, nil
	}
	return xml.Attr{
		Name:  name,
		Value: fmt.Sprintf(`url(%q)`, u),
	}, nil
}

type Transformation struct {
	Command string
	Args    []float64
}

func (t Transformation) String() string {
	sargs := make([]string, len(t.Args))
	for i, f := range t.Args {
		sargs[i] = fmt.Sprint(f)
	}
	return t.Command + `(` + strings.Join(sargs, `,`) + `)`
}

type TransformAttr struct {
	Transformations []Transformation
}

func (ta TransformAttr) String() string {
	strans := make([]string, len(ta.Transformations))
	for i, t := range ta.Transformations {
		strans[i] = t.String()
	}
	return strings.Join(strans, ` `)
}

func (ta TransformAttr) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if len(ta.Transformations) == 0 {
		return xml.Attr{}, nil
	}
	return xml.Attr{
		Name:  name,
		Value: ta.String(),
	}, nil
}

func (ta *TransformAttr) Matrix(a, b, c, d, e, f float64) {
	ta.Transformations = append(ta.Transformations, Transformation{`matrix`, []float64{a, b, c, d, e, f}})
}
func (ta *TransformAttr) TranslateX(x float64) {
	ta.Transformations = append(ta.Transformations, Transformation{`translate`, []float64{x}})
}
func (ta *TransformAttr) Translate(x, y float64) {
	ta.Transformations = append(ta.Transformations, Transformation{`translate`, []float64{x, y}})
}
func (ta *TransformAttr) ScaleUniform(s float64) {
	ta.Transformations = append(ta.Transformations, Transformation{`scale`, []float64{s}})
}
func (ta *TransformAttr) Scale(x, y float64) {
	ta.Transformations = append(ta.Transformations, Transformation{`scale`, []float64{x, y}})
}
func (ta *TransformAttr) Rotate(r float64) {
	ta.Transformations = append(ta.Transformations, Transformation{`rotate`, []float64{r}})
}
func (ta *TransformAttr) RotateAbout(r, x, y float64) {
	ta.Transformations = append(ta.Transformations, Transformation{`rotate`, []float64{r, x, y}})
}
func (ta *TransformAttr) Skew(r float64) {
	ta.Transformations = append(ta.Transformations, Transformation{`skew`, []float64{r}})
}
func (ta *TransformAttr) SkewY(r float64) {
	ta.Transformations = append(ta.Transformations, Transformation{`skewY`, []float64{r}})
}

type ColorAttr string

func (ca *ColorAttr) FromColor(c color.Color) {
	r, g, b, a := c.RGBA()
	*ca = ColorAttr(fmt.Sprintf(`rgba(%d, %d, %d, %d)`, r/0x100, g/0x100, b/0x100, a/0x100))
}

func (ca *ColorAttr) FromNamedColor(name string) {
	rgba, ok := colornames.Map[name]
	if ok {
		ca.FromColor(rgba)
		return
	}
	*ca = ColorAttr(name)
}

func (ca ColorAttr) String() string {
	return string(ca)
}

func (ca ColorAttr) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if ca == `` {
		return xml.Attr{}, nil
	}
	return xml.Attr{
		Name:  name,
		Value: ca.String(),
	}, nil
}
