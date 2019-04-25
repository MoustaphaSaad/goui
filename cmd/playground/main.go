package main

import (
	"fmt"

	"github.com/MoustaphaSaad/goui/internal/pkg/wingui"
)

func main() {
	fmt.Println("Hello, World!")
	win, err := wingui.CreateWindow("Hello World", 1280, 720)
	if err != nil {
		fmt.Println(err)
		return
	}

	for win.Running {
		win.Poll()
	}
}
