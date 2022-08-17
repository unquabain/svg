package svg

import "encoding/xml"

type Line struct {
	XMLName xml.Name `xml:"line"`
	CoreAttributes
	StylingAttributes
	PresentationAttributes
	X1         float64 `xml:"x1,attr"`
	X2         float64 `xml:"x2,attr"`
	Y1         float64 `xml:"y1,attr"`
	Y2         float64 `xml:"y2,attr"`
	PathLength float64 `xml:"pathLength,attr,omitempty"`
}
