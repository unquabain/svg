package svg

import "encoding/xml"

type G struct {
	XMLName xml.Name `xml:"g"`
	CoreAttributes
	StylingAttributes
	PresentationAttributes
	HasChildren
}
