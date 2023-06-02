package main

import (
	"fmt"
	_ "math/rand"

	//	"math/rand"

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
		// Проверяем, сработал ли хук
		found := findBar()
		if found != -1 {
			//fmt.Printf("Found bar# %d in column %d\n", counter, found+1)
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

func clickOnBar(i int) {
	x := barXs[i]
	y := barY
	fmt.Printf("%d) Move & click (%d, %d)\n", counter, x, y)
	robotgo.MoveClick(x, y, "left")

	y += 75
	var state int = getState(x, y)
	if state == 1 {
		fmt.Printf("%d) Repeat-1 (%d, %d)\n", counter, x, y)
		robotgo.MoveClick(x, y, "left")
	} else if state == 2 {
		return
	}

	y += 75
	state = getState(x, y)
	if state == 1 {
		fmt.Printf("%d) Repeat-2 (%d, %d)\n", counter, x, y)
		robotgo.MoveClick(x, y, "left")
	} else if state == 2 {
		return
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
