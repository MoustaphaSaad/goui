package main

import (
	"math/rand"
	"fmt"
	"flag"
	"os"
	"log"
	"runtime/pprof"

	ui "github.com/MoustaphaSaad/goui/internal/pkg/goui"
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

	for i := 0; i < 1000; i++ {
		win.Point(
			float32(rand.Uint32() % 1280),
			float32(rand.Uint32() % 720),
			float32(rand.Uint32() % 50),
		)
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
