package effect

import (
	"encoding/xml"

	"github.com/Unquabain/svg"
)

type CompositeOperator string

const (
	CompositeOperatorOver       CompositeOperator = "over"
	CompositeOperatorIn         CompositeOperator = "in"
	CompositeOperatorOut        CompositeOperator = "out"
	CompositeOperatorAtop       CompositeOperator = "atop"
	CompositeOperatorXor        CompositeOperator = "xor"
	CompositeOperatorLighter    CompositeOperator = "lighter"
	CompositeOperatorArithmetic CompositeOperator = "arithmetic"
)

func (co CompositeOperator) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if co == `` {
		return xml.Attr{}, nil
	}
	return xml.Attr{
		Name:  name,
		Value: string(co),
	}, nil
}

type Composite struct {
	XMLName xml.Name `xml:"feComposite"`

	svg.CoreAttributes
	svg.StylingAttributes
	svg.PresentationAttributes
	FilterPrimitive

	In       string            `xml:"in,attr"`
	In2      string            `xml:"in2,attr"`
	Operator CompositeOperator `xml:"operator,attr"`

	K1 float64 `xml:"k1,attr,omitempty"`
	K2 float64 `xml:"k2,attr,omitempty"`
	K3 float64 `xml:"k3,attr,omitempty"`
	K4 float64 `xml:"k4,attr,omitempty"`
}
