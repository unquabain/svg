package svg

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
)

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

func (s SVG) Writer() (io.Writer, error) {
	data, err := xml.Marshal(s)
	if err != nil {
		return nil, fmt.Errorf(`unable to serialize SVG: %w`, err)
	}
	return bytes.NewBuffer(data), nil
}
func (s SVG) WriterIndent(prefix, tab string) (io.Writer, error) {
	data, err := xml.MarshalIndent(s, prefix, tab)
	if err != nil {
		return nil, fmt.Errorf(`unable to serialize SVG: %w`, err)
	}
	return bytes.NewBuffer(data), nil
}
