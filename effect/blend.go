package effect

import (
	"encoding/xml"

	"github.com/Unquabain/svg"
)

type BlendMode string

const (
	BlendModeNormal     BlendMode = "normal"
	BlendModeMultiply   BlendMode = "multiply"
	BlendModeScreen     BlendMode = "screen"
	BlendModeOverlay    BlendMode = "overlay"
	BlendModeDarken     BlendMode = "darken"
	BlendModeLighten    BlendMode = "lighten"
	BlendModeColorDodge BlendMode = "color-dodge"
	BlendModeColorBurn  BlendMode = "color-burn"
	BlendModeHardLight  BlendMode = "hard-light"
	BlendModeSoftLight  BlendMode = "soft-light"
	BlendModeDifference BlendMode = "difference"
	BlendModeExclusion  BlendMode = "exclusion"
	BlendModeHue        BlendMode = "hue"
	BlendModeSaturation BlendMode = "saturation"
	BlendModeColor      BlendMode = "color"
	BlendModeLuminosity BlendMode = "luminosity"
)

func (bm BlendMode) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if bm == `` {
		return xml.Attr{}, nil
	}
	return xml.Attr{
		Name:  name,
		Value: string(bm),
	}, nil
}

type Blend struct {
	XMLName xml.Name `xml:"feBlend"`

	svg.CoreAttributes
	svg.PresentationAttributes
	svg.StylingAttributes
	FilterPrimitive

	In   string    `xml:"in,attr"`
	In2  string    `xml:"in2,attr"`
	Mode BlendMode `xml:"mode,attr"`
}
