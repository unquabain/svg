package svg

import "encoding/xml"

type TextContent string

func (tc TextContent) MarshalText() ([]byte, error) {
	return []byte(tc), nil
}

type Text struct {
	XMLName xml.Name `xml:"text"`
	CoreAttributes
	StylingAttributes
	PresentationAttributes
	X            float64     `xml:"x,attr"`
	Y            float64     `xml:"y,attr"`
	DX           float64     `xml:"dx,attr,omitempty"`
	DY           float64     `xml:"dy,attr,omitempty"`
	Rotate       float64     `xml:"rotate,attr,omitempty"`
	LengthAdjust string      `xml:"lengthAdjust,attr,omitempty"`
	TextLength   string      `xml:"textLength,attr,omitempty"`
	Content      TextContent `xml:",chardata"`
}
