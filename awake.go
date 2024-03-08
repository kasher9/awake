package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		robotgo.Move(0, 0)
		robotgo.DragSmooth(1, 1, 0.5, 1.0)
		time.Sleep(60 * time.Second)
		fmt.Print(time.Now())
	}
}
