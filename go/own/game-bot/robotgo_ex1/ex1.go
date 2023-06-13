package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
	fmt.Print("Hello World!\n")
	robotgo.MouseSleep = 300
	processBot()
	systemstack(func() { println("Hello, World!") })
}

func moveMouse() {
	x := 380
	y := 180
	fmt.Printf("Move (%d, %d)\n", x, y)
	robotgo.Move(x, y)
	robotgo.Click("left", true)
	fmt.Println(robotgo.GetLocationColor())
	fmt.Println(robotgo.GetPixelColor(x+300, y))
}

func processBot() {
	quitCh := registerHooks()
	//registerMouseClick()

	for {
		// Выполняйте здесь вашу работу
		fmt.Println("Выполняется работа...")

		// Добавьте задержку, чтобы не производить работу слишком быстро
		time.Sleep(time.Second)

		// Проверяем, сработал ли хук
		select {
		case <-quitCh:
			// Хук сработал, завершаем цикл
			goto doneLoop
		default:
			// Хук не сработал, продолжаем цикл
		}
	}

doneLoop:
}

func registerHooks() (out chan bool) {
	fmt.Println("--- Please press q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q"}, func(e hook.Event) {
		fmt.Println("quit..")
		hook.End()
	})

	hook.Register(hook.MouseDown, nil, func(e hook.Event) {
		fmt.Println(robotgo.Location())
	})

	s := hook.Start()
	return hook.Process(s)
}
