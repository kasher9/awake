package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		x0, y0 := robotgo.Location()
		robotgo.MoveSmooth(0, 0)
		robotgo.MoveSmooth(x0, y0)
		time.Sleep(55 * time.Second)
		fmt.Print(time.Now())
	}
}
