package svg

import "encoding/xml"

type Filter struct {
	XMLName xml.Name `xml:"filter"`

	CoreAttributes
	PresentationAttributes
	StylingAttributes
	HasChildren

	X              float64 `xml:"x,attr,omitempty"`
	Y              float64 `xml:"y,attr,omitempty"`
	Width          float64 `xml:"width,attr,omitempty"`
	Height         float64 `xml:"height,attr,omitempty"`
	FilterUnits    string  `xml:"filterUnits,attr,omitempty"`
	PrimitiveUnits string  `xml:"primitiveUnits,attr,omitempty"`
}
