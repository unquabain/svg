package svg

import "encoding/xml"

type ClipPath struct {
	XMLName xml.Name `xml:"clipPath"`
	CoreAttributes
	StylingAttributes
	PresentationAttributes
	HasChildren
	ClipPathUnits string `xml:"clipPathUnits,attr,omitempty"`
}
