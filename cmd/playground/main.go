package main

import (
	"fmt"
	"flag"
	"os"
	"log"
	"runtime/pprof"
	"math"

	ui "github.com/MoustaphaSaad/goui/internal/pkg/goui"
	"github.com/MoustaphaSaad/goui/internal/pkg/img"
	"github.com/MoustaphaSaad/goui/internal/pkg/wingui"
)

const (
	windowWidth  = 1280
	windowHeight = 720
)

type Framer struct{
	e *ui.Engine
}

var r = float32(0)
var t = float32(1)
var theta = float32(0)
func (f Framer) Frame() img.Image {
	f.e.FrameBegin()

	origin := ui.V2{windowWidth/2, windowHeight/2}

	if true {
		f.e.DrawCircle(origin.Add(ui.V2{-200, -200}), r, ui.Color{128, 0, 0, 255})
		f.e.DrawCircle(origin.Add(ui.V2{+200, -200}), r, ui.Color{0, 128, 0, 255})
		f.e.DrawCircle(origin.Add(ui.V2{-200, +200}), r, ui.Color{0, 0, 128, 255})
		f.e.DrawCircle(origin.Add(ui.V2{+200, +200}), r, ui.Color{128, 128, 128, 255})

		b := origin.Add(ui.V2{float32(math.Sin(float64(theta))), float32(math.Cos(float64(theta)))}.Scale(200))
		f.e.DrawLine(origin, b, 5, ui.Color{255, 0, 0, 255})

		b = origin.Add(ui.V2{float32(math.Sin(float64(theta/6))), float32(math.Cos(float64(theta/6)))}.Scale(200))
		f.e.DrawLine(origin, b, 5, ui.Color{0, 255, 0, 255})

		b = origin.Add(ui.V2{float32(math.Sin(float64(theta/36))), float32(math.Cos(float64(theta/36)))}.Scale(200))
		f.e.DrawLine(origin, b, 5, ui.Color{0, 0, 255, 255})

		f.e.DrawQuad(origin, ui.V2{+r, +r}, ui.Color{0, 128, 128, 255})
		f.e.DrawQuad(origin, ui.V2{-r, -r}, ui.Color{128, 128, 0, 255})
		f.e.DrawQuad(origin, ui.V2{-r, +r}, ui.Color{128, 0, 128, 255})
		f.e.DrawQuad(origin, ui.V2{+r, -r}, ui.Color{128, 128, 128, 255})

		b = origin.Add(ui.V2{float32(math.Sin(float64(theta))), float32(math.Cos(float64(theta)))}.Scale(100))
		f.e.DrawTriangle(origin, origin.Add(ui.V2{100, 100}), b, ui.Color{255, 255, 128, 255})
	}

	r += t
	if r > 400 || r < 0 { t *= -1 }

	theta += 0.01

	return f.e.FrameEnd()
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")
func main() {
	flag.Parse()
	if *cpuprofile != "" {
		fmt.Println("Hello, World!")
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	f := Framer{ui.NewEngine(windowWidth, windowHeight)}

	win, _ := wingui.CreateWindow("Hello, World!", windowWidth, windowHeight, f)
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
