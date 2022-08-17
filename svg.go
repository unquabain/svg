package svg

import "encoding/xml"

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	XMLNS   XMLNS    `xml:"xmlns,attr"`
	CoreAttributes
	StylingAttributes
	PresentationAttributes
	HasChildren
	ViewBox             ViewBoxAttr `xml:"viewBox,attr"`
	PreserveAspectRatio string      `xml:"preserveAspectRatio,attr,omitempty"`
	X                   float64     `xml:"x,attr,omitempty"`
	Y                   float64     `xml:"y,attr,omitempty"`
}
