package svg

import (
	"encoding/xml"
	"fmt"
)

type ViewBoxAttr struct {
	Left,
	Top,
	Width,
	Height int
}

func (vba ViewBoxAttr) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: fmt.Sprintf(`%d %d %d %d`, vba.Left, vba.Top, vba.Width, vba.Height),
	}, nil
}
