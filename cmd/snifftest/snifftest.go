package main

import (
	"encoding/xml"
	"fmt"

	svg "github.com/Unquabain/svglib"
)

func main() {
	doc := svg.SVG{}
	doc.ViewBox.Width = 100
	doc.ViewBox.Height = 100
	circle := svg.Circle{CX: 50, CY: 50, R: 50}
	circle.Stroke = `none`
	circle.Fill = `black`
	rect := svg.Rect{X: 25, Width: 50, Y: 25, Height: 50}
	rect.Fill = `white`
	rect.StrokeWidth = `2`
	rect.Stroke = `red`
	doc.Add(
		circle,
		rect,
	)
	data, err := xml.MarshalIndent(doc, ``, `    `)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
