package svg

import (
	"encoding/xml"
)

type Script struct {
	XMLName xml.Name `xml:"script"`
	CoreAttributes
	StylingAttributes
	CrossOrigin string `xml:"crossorigin,attr,omitempty"`
	HRef        string `xml:"href,attr,omitempty"`
	Type        string `xml:"type,attr,omitempty"`
	Content     string `xml:",cdata"`
}
