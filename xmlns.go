package svg

import "encoding/xml"

type XMLNS struct{}

func (_ XMLNS) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: `http://www.w3.org/2000/svg`,
	}, nil
}
