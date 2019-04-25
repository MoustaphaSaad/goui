package main

import (
	"fmt"
	"math/rand"

	"github.com/MoustaphaSaad/goui/internal/pkg/wingui"
	g "github.com/MoustaphaSaad/goui/internal/pkg/geometry"
	ui "github.com/MoustaphaSaad/goui/internal/pkg/goui"
)

var shader = ui.NewCircleShader()

func render(w *wingui.Window) {
	buffer := w.Chain.Back()
	buffer.Clear()
	for i := 0; i < 100; i++ {
		shader.Circle(ui.Circle{
			Center: g.Vec2{X: rand.Uint32() % w.Width, Y: rand.Uint32() % w.Height},
			Radius: rand.Uint32() % 10,
		})
	}
	shader.Render(buffer)
	w.Chain.Swap()
	// fmt.Println("Rendering")
}

func main() {
	fmt.Println("Hello, World!")
	win, err := wingui.CreateWindow("Hello World", 1280, 720, render)
	if err != nil {
		fmt.Println(err)
		return
	}

	for win.Running {
		win.Poll()
	}
}
