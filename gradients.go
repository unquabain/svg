package svg

import "encoding/xml"

type Stop struct {
	XMLName xml.Name `xml:"stop"`

	CoreAttributes
	StylingAttributes
	PresentationAttributes

	Offset      Percentage `xml:"offset,attr"`
	StopColor   ColorAttr  `xml:"stop-color,attr"`
	StopOpacity float64    `xml:"stop-opacity,attr,omitempty"`
}

type LinearGradient struct {
	XMLName xml.Name `xml:"linearGradient"`

	CoreAttributes
	StylingAttributes
	PresentationAttributes

	GradientUnits     string     `xml:"gradientUnits,attr,omitempty"`
	GradientTransform string     `xml:"gradientTransform,attr,omitempty"`
	HRef              string     `xml:"href,attr,omitempty"`
	SpreadMethod      string     `xml:"spreadMethod,attr,omitempty"`
	X1                Percentage `xml:"x1,attr"`
	X2                Percentage `xml:"x2,attr"`
	Y1                Percentage `xml:"y1,attr"`
	Y2                Percentage `xml:"y2,attr"`
	Stops             []Stop     `xml:"stop"`
}

func (lg *LinearGradient) AddStops(stops ...Stop) {
	lg.Stops = append(lg.Stops, stops...)
}

type RadialGradient struct {
	XMLName xml.Name `xml:"radialGradient"`

	CoreAttributes
	StylingAttributes
	PresentationAttributes

	GradientUnits     string  `xml:"gradientUnits,attr,omitempty"`
	GradientTransform string  `xml:"gradientTransform,attr,omitempty"`
	HRef              string  `xml:"href,attr,omitempty"`
	SpreadMethod      string  `xml:"spreadMethod,attr,omitempty"`
	CX                float64 `xml:"cx,attr"`
	CY                float64 `xml:"cy,attr"`
	FX                float64 `xml:"fx,attr"`
	FY                float64 `xml:"fy,attr"`
	FR                float64 `xml:"fr,attr"`
	R                 float64 `xml:"r,attr"`
	Stops             []Stop  `xml:"stop"`
}

func (rg *RadialGradient) AddStops(stops ...Stop) {
	rg.Stops = append(rg.Stops, stops...)
}
