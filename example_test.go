package svg_test

import (
	"fmt"
	"image/color"

	"github.com/Unquabain/svg"
)

func ExampleSVG() {
	rectColor := color.RGBA{0, 0xFF, 0, 0xFF}
	circleColor := color.RGBA{0xFF, 0, 0, 0xFF}
	doc := svg.SVG{}
	doc.ViewBox.Height = 100
	doc.ViewBox.Width = 100

	g := svg.G{}
	rect := svg.Rect{X: 0, Y: 0, Width: 100, Height: 100}
	rect.Fill = svg.FromColor(rectColor)

	circle := svg.Circle{CX: 50, CY: 50, R: 50}
	circle.Fill = svg.FromColor(circleColor)

	g.Add(rect, circle)
	doc.Add(g)

	w, err := doc.WriterIndent(``, `  `)
	if err != nil {
		panic(err)
	}
	fmt.Print(w)

	// Output:
	// <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100">
	//   <g>
	//     <rect fill="rgba(0, 255, 0, 255)" x="0" y="0" width="100" height="100"></rect>
	//     <circle fill="rgba(255, 0, 0, 255)" cx="50" cy="50" r="50"></circle>
	//   </g>
	// </svg>
}
