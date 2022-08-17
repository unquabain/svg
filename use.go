package svg

import "encoding/xml"

type Use struct {
	XMLName xml.Name `xml:"use"`

	CoreAttributes
	StylingAttributes
	PresentationAttributes

	HRef   string  `xml:"href,attr"`
	X      float64 `xml:"x,attr"`
	Y      float64 `xml:"y,attr"`
	Width  float64 `xml:"width,attr,omitempty"`
	Height float64 `xml:"height,attr,omitempty"`
}
