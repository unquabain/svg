package svg

import "encoding/xml"

type Circle struct {
	XMLName xml.Name `xml:"circle"`
	CoreAttributes
	StylingAttributes
	PresentationAttributes
	CX         float64 `xml:"cx,attr"`
	CY         float64 `xml:"cy,attr"`
	R          float64 `xml:"r,attr"`
	PathLength float64 `xml:"pathLength,attr,omitempty"`
}
