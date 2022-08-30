package main

import (
	"fmt"
	"image/color"

	"github.com/Unquabain/svg"
)

var ColorLight = color.Gray{0xEE}
var ColorDark = color.Gray{0x44}

func drawRect(id, label string, x, y, width, height float64) svg.G {
	g := svg.G{}
	rect := svg.Rect{X: x, Y: y, Width: width, Height: height}
	rect.ID = id
	rect.Fill = svg.FromColor(ColorDark)

	text := svg.Text{
		X:       x + (width / 2),
		Y:       y + (height / 2),
		Content: svg.TextContent(label),
	}
	g.Add(rect, text)
	return g
}

func drawBlock(blockIDPrefix string, multiplier, count int, x, y, width, height float64) svg.G {
	g := svg.G{}
	bottom := y + height
	blockHeight := height / float64(count-1)
	// From 1 through count - 1
	for i := 1; i < count; i++ {
		id := fmt.Sprintf(`%s-%d-%d`, blockIDPrefix, multiplier, i*multiplier)
		label := fmt.Sprintf(`%d`, i*multiplier)
		b := drawRect(id, label, x, bottom-blockHeight, width, blockHeight)
		g.Add(b)
		bottom -= blockHeight
	}
	return g
}

func drawBlocks() svg.G {
	g := svg.G{}
	g.Add(
		drawBlock(`h`, 12, 2, 0, 0, 300, 300),
		drawBlock(`h`, 3, 4, 300, 0, 300, 300),
		drawBlock(`h`, 1, 3, 600, 0, 300, 300),

		drawBlock(`m`, 15, 4, 0, 300, 300, 300),
		drawBlock(`m`, 5, 3, 300, 300, 300, 300),
		drawBlock(`m`, 1, 5, 600, 300, 300, 300),

		drawBlock(`s`, 15, 4, 0, 600, 300, 300),
		drawBlock(`s`, 5, 3, 300, 600, 300, 300),
		drawBlock(`s`, 1, 5, 600, 600, 300, 300),
	)
	return g
}

func drawStyle() svg.Style {
	return svg.Style{
		Content: fmt.Sprintf(`
text {
	font-family: sans-serif;
	font-size: 60px;
	fill: %s;
	text-anchor: middle;
	dominant-baseline: middle;
}`, svg.FromColor(ColorLight))}
}

func drawScript() svg.Script {
	return svg.Script{Content: `
function display(unit, multiplier, max, value) {
	const idPrefix = '' + unit + '-' + multiplier + '-'
	for (var i = max; i > 0; i -= multiplier) {
			const id = idPrefix + i
			if (value >= i) {
				document.getElementById(id).style.display = 'block'
			} else {
				document.getElementById(id).style.display = 'none'
			}
	}
}
setInterval(function() {
	const now = new Date()
	const t = Math.floor(now.getTime() / 1000 - now.getTimezoneOffset()*60)
	const s = t % 60
	const m = Math.floor(t / 60) % 60 
	const h = Math.floor(t / 3600) % 24

	display('s', 1, 4, s % 5)
	display('s', 5, 10, s % 15)
	display('s', 15, 45, s % 60)

	display('m', 1, 4, m % 5)
	display('m', 5, 10, m % 15)
	display('m', 15, 45, m % 60)

	display('h', 1, 2, h % 3)
	display('h', 3, 9, h % 12)
	display('h', 12, 12, h % 24)
}, 500)
`}
}

func main() {
	doc := svg.SVG{}
	doc.ViewBox.Width = 900
	doc.ViewBox.Height = 900
	doc.Style = fmt.Sprintf(`background-color: %s`, svg.FromColor(ColorLight))
	doc.Add(
		drawStyle(),
		drawBlocks(),
		drawScript(),
	)
	writer, err := doc.WriterIndent(``, `  `)
	if err != nil {
		panic(err)
	}
	fmt.Println(writer)
}
