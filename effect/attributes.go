package effect

type FilterPrimitive struct {
	Height float64 `xml:"height,attr,omitempty"`
	Width  float64 `xml:"width,attr,omitempty"`
	X      float64 `xml:"x,attr,omitempty"`
	Y      float64 `xml:"y,attr,omitempty"`
	Result string  `xml:"result,attr,omitempty"`
}
