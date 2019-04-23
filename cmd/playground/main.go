package main

import (
	"fmt"

	"github.com/MoustaphaSaad/goui/internal/pkg/wingui"
)

func main() {
	fmt.Println("Hello, World!")
	api := wingui.NewAPI()
	win, err := api.CreateWindow("Hello World", 1280, 720)
	if err != nil {
		fmt.Println(err)
		return
	}

	api.Exec(&win)
}
