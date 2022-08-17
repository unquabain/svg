package svg

import "encoding/xml"

type Defs struct {
	XMLName xml.Name `xml:"defs"`
	CoreAttributes
	StylingAttributes
	HasChildren
}
