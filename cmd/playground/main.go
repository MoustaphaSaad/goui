package main

import (
	"math/rand"
	"fmt"
	"flag"
	"os"
	"log"
	"runtime/pprof"

	ui "github.com/MoustaphaSaad/goui/internal/pkg/goui"
	"github.com/MoustaphaSaad/goui/internal/pkg/geometry"
)

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

	win, err := ui.NewWindow(1280, 720, "Hello World")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 100; i++ {
		win.ChildAdd(ui.NewCircle(geometry.Vec2{
			X: float32(rand.Uint32()%1280),
			Y: float32(rand.Uint32()%720),
		},
		float32(rand.Uint32() % 100)))
	}

	win.Exec()

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
