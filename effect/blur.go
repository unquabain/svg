package effect

import (
	"encoding/xml"

	"github.com/Unquabain/svg"
)

type BlurEdgeMode string

const (
	BlurEdgeModeDuplicate BlurEdgeMode = `duplicate`
	BlurEdgeModeWrap      BlurEdgeMode = `wrap`
	BlurEdgeModeNone      BlurEdgeMode = `none`
)

func (bem BlurEdgeMode) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if bem == `` {
		return xml.Attr{}, nil
	}
	return xml.Attr{
		Name:  name,
		Value: string(bem),
	}, nil
}

type Blur struct {
	XMLName xml.Name `xml:"feGaussianBlur"`

	svg.CoreAttributes
	svg.PresentationAttributes
	svg.StylingAttributes
	FilterPrimitive

	In           string       `xml:"in,attr"`
	StdDeviation float64      `xml:"stdDeviation,attr"`
	EdgeMode     BlurEdgeMode `xml:"edgeMode,attr,omitempty"`
}
