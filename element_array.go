package svg

import (
	"encoding/xml"
	"fmt"
	"log"
)

type ElementArray []any

func (ea ElementArray) MarshalXML(e *xml.Encoder, _ xml.StartElement) error {
	for _, el := range ea {
		if err := e.Encode(el); err != nil {
			log.Print(el)
			return fmt.Errorf(`could not encode child element of type %T: %w`, el, err)
		}
	}
	return nil
}
