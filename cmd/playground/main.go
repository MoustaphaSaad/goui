package main

import (
	"fmt"
	"math/rand"
	"flag"
	"os"
	"log"
	"runtime/pprof"

	"github.com/MoustaphaSaad/goui/internal/pkg/wingui"
	g "github.com/MoustaphaSaad/goui/internal/pkg/geometry"
	ui "github.com/MoustaphaSaad/goui/internal/pkg/goui"
)

var shader = ui.NewCircleShader()

func render(w *wingui.Window) {
	buffer := w.Chain.Back()
	buffer.Clear()
	shader.Circle(ui.Circle{
		Center: g.Vec2{X: 1280/2, Y: 720/2},
		Radius: 200/2,
	})
	for i := 0; i < 0; i++ {
		shader.Circle(ui.Circle{
			Center: g.Vec2{X: float32(rand.Uint32() % w.Width), Y: float32(rand.Uint32() % w.Height)},
			Radius: float32(rand.Uint32() % 100),
		})
	}
	shader.Render(buffer)
	w.Chain.Swap()
	// fmt.Println("Rendering")
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")
func main() {
	fmt.Println("Hello, World!")
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	win, err := wingui.CreateWindow("Hello World", 1280, 720, render)
	if err != nil {
		fmt.Println(err)
		return
	}

	for win.Running {
		win.Poll()
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}
}
