package main

import (
	"encoding/xml"
	"fmt"

	"github.com/Unquabain/svg"
	"github.com/Unquabain/svg/effect"
)

func grad() svg.LinearGradient {
	grad := svg.LinearGradient{X1: 0, Y1: 0, X2: 1, Y2: 1}
	grad.ID = `gradient01`
	grad.AddStops(
		svg.Stop{Offset: 0, StopColor: `black`},
		svg.Stop{Offset: 0.5, StopColor: `darkorchid`},
		svg.Stop{Offset: 1, StopColor: `red`},
	)
	return grad
}

func blur() svg.Filter {
	blur := svg.Filter{}
	blur.ID = `blur`
	blur.Add(effect.Blur{In: `SourceGraphic`, StdDeviation: 0.5})
	return blur
}

func defs() svg.Defs {
	defs := svg.Defs{}
	defs.Add(
		grad(),
		blur(),
	)
	return defs
}

func circle() svg.Circle {
	circle := svg.Circle{CX: 50, CY: 50, R: 50}
	circle.Stroke = `none`
	circle.Fill = `url(#gradient01)`
	return circle
}

func rect() svg.Rect {
	rect := svg.Rect{X: 25, Width: 50, Y: 25, Height: 50}
	rect.Fill = `white`
	rect.StrokeWidth = `2`
	rect.Stroke = `deeppink`
	return rect
}

func path() svg.Path {
	path := svg.Path{}
	path.Fill = `cyan`
	path.FillOpacity = `0.25`
	path.Stroke = `blue`
	path.StrokeWidth = `0.5`
	path.D.MoveTo(10, 10)
	path.D.LineTo(90, 10)
	path.D.QuadraticTo(10, 10, 10, 90)
	path.D.Close()
	path.Filter = `url(#blur)`
	return path
}

func doc() svg.SVG {
	doc := svg.SVG{}
	doc.ViewBox.Width = 100
	doc.ViewBox.Height = 100

	r := rect()
	r.Transform.RotateAbout(45, 50, 50)
	doc.Add(
		defs(),
		circle(),
		r,
		path(),
	)
	return doc
}

func main() {
	data, err := xml.MarshalIndent(doc(), ``, `    `)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
