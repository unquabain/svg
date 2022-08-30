package svg

import (
	"encoding/xml"
)

type Style struct {
	XMLName xml.Name `xml:"style"`
	CoreAttributes
	StylingAttributes
	Type    string `xml:"type,attr,omitempty"`
	Media   string `xml:"media,attr,omitempty"`
	Title   string `xml:"title,attr,omitempty"`
	Content string `xml:",cdata"`
}
