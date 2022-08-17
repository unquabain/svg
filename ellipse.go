package svg

import "encoding/xml"

type Ellipse struct {
	XMLName xml.Name `xml:"ellipse"`
	CoreAttributes
	StylingAttributes
	PresentationAttributes
	CX         float64 `xml:"cx,attr"`
	CY         float64 `xml:"cy,attr"`
	RX         float64 `xml:"rx,attr"`
	RY         float64 `xml:"ry,attr"`
	PathLength float64 `xml:"pathLength,attr,omitempty"`
}
