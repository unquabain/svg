package svg

import "encoding/xml"

type Rect struct {
	XMLName xml.Name `xml:"rect"`
	CoreAttributes
	StylingAttributes
	PresentationAttributes
	X          float64 `xml:"x,attr"`
	Y          float64 `xml:"y,attr"`
	Width      float64 `xml:"width,attr"`
	Height     float64 `xml:"height,attr"`
	RX         float64 `xml:"rx,attr,omitempty"`
	RY         float64 `xml:"ry,attr,omitempty"`
	PathLength float64 `xml:"pathLength,attr,omitempty"`
}
