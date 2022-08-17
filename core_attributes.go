package svg

type CoreAttributes struct {
	ID       string `xml:"id,attr,omitempty"`
	Lang     string `xml:"lang,attr,omitempty"`
	TabIndex int    `xml:"tabindex,attr,omitempty"`
	XMLBase  string `xml:"xml:base,attr,omitempty"`
	XMLLang  string `xml:"xml:lang,attr,omitempty"`
	XMLSpace string `xml:"xml:space,attr,omitempty"`
}
