package main

import (
	"fmt"
	_ "math/rand"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var barXs = []int{750, 890, 1050, 1140} // full window
// var barXs = []int{140, 230, 330, 400} // small window

var barY = 400
var counter = 1

func main() {
	robotgo.MouseSleep = 0
	processBot()
}

func processBot() {
	quitCh := registerHooks()

	for {
		found := findBar()
		if found != -1 {
			clickOnBar(found)
			counter++
		}

		select {
		case <-quitCh:
			goto doneLoop
		default:
		}
	}

doneLoop:
	fmt.Printf("Total bars: %d\n", counter)
}

func findBar() int {
	for i, x := range barXs {
		if isTargetBar(x, barY) {
			return i
		}
	}
	return -1
}

func isTargetBar(x, y int) bool {
	color := robotgo.GetPixelColor(x, y)
	return color == "000000"
}

func isHittedBar(x, y int) bool {
	color := robotgo.GetPixelColor(x, y)
	if color[0] == '6' {
		return true
	}
	return false
}

func clickOnBar(column int) {
	x := barXs[column]
	y := barY
	column++
	fmt.Printf("%d) Move & click (%d: %d, %d)\n", counter, column, x, y)
	robotgo.MoveClick(x, y, "left")
	if isTargetBar(x, y+50) {
		robotgo.MoveClick(x, y+50, "left", true)
	}
}

func clickOnBar2(column int) {
	x := barXs[column]
	y := barY
	column++
	fmt.Printf("%d) Move & click (%d: %d, %d)\n", counter, column, x, y)
	robotgo.MoveClick(x, y, "left")
	inc := 50

	for i := 1; i <= 3; i++ {
		y += inc
		xr := x // + rand.Intn(10) - 5
		robotgo.Move(xr, y)
		var state int = getState(xr, y)
		if state == 1 {
			robotgo.Click("left", i%2 == 1)
			fmt.Printf("%d) Repeat-%d (%d: %d, %d)\n", counter, i, column, xr, y)
		} else if state == 2 {
			fmt.Printf("%d) Hitted found (%d: %d, %d)\n", counter, column, xr, y)
			return
		} else {
			fmt.Printf("%d) Empty found (%d: %d, %d)\n", counter, column, xr, y)
		}

		if getState(xr, y-inc) == 2 {
			fmt.Printf("%d) Re-check previous position and found hitted (%d: %d, %d)\n", counter, column, xr, y-75)
			return
		}
	}
}

func getState(x, y int) int {
	if isTargetBar(x, y) {
		return 1
	}
	if isHittedBar(x, y) {
		return 2
	}
	return 0
}

func registerHooks() (out chan bool) {
	fmt.Println("--- Please press q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q"}, func(e hook.Event) {
		fmt.Println("quit..")
		hook.End()
	})

	// hook.Register(hook.MouseDown, []string{}, func(e hook.Event) {
	// 	fmt.Println(robotgo.GetLocationColor())
	// })

	s := hook.Start()
	return hook.Process(s)
}
