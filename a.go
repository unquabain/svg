package svg

import "encoding/xml"

type A struct {
	XMLName xml.Name `xml:"a"`
	CoreAttributes
	StylingAttributes
	PresentationAttributes
	HasChildren
	HRef     string `xml:"href,attr"`
	HRefLang string `xml:"hreflang,attr,omitempty"`
	Target   string `xml:"target,attr,omitempty"`
	Type     string `xml:"type,attr,omitempty"`
}
