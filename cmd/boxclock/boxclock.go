package main

import (
	_ "embed"
	"encoding/base64"
	"fmt"
	"image/color"

	"github.com/Unquabain/svg"
)

//go:generate sh -c "GOOS=js GOARCH=wasm go build -o generated/script.wasm wasm/boxclock.go"

//go:embed style.css
var style string

//go:embed script.js
var script string

//go:embed wasm_exec.js
var wasmExec string

//go:embed generated/script.wasm
var wasm []byte

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
	return svg.Style{Content: fmt.Sprintf(style, svg.FromColor(ColorLight))}
}

func drawScript() svg.Script {
	data := base64.StdEncoding.EncodeToString(wasm)
	return svg.Script{Content: fmt.Sprintf(script, data)}
}

func main() {
	doc := svg.SVG{}
	doc.ViewBox.Width = 900
	doc.ViewBox.Height = 900
	doc.Style = fmt.Sprintf(`background-color: %s`, svg.FromColor(ColorLight))
	doc.Add(
		drawStyle(),
		drawBlocks(),
		svg.Script{Content: wasmExec},
		drawScript(),
	)
	writer, err := doc.WriterIndent(``, `  `)
	if err != nil {
		panic(err)
	}
	fmt.Println(writer)
}
