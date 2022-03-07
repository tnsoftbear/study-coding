package application

import (
	"bufio"
	"flag"
	"os"
	"segmented-sieve-oop/primesearch"
	"segmented-sieve-oop/profiler"
)

const IsProfilingDef = true
const IsResultOutput = true
const AreaSize = 1000001

type App struct {
	reader *bufio.Reader
	writer *bufio.Writer
}

func New() *App {
	return &App{}
}

func (app *App) Run() {
	app.reader = bufio.NewReader(os.Stdin)
	app.writer = bufio.NewWriter(os.Stdout)

	isProfiling := flag.Bool("profiling", IsProfilingDef, "Enable time profiling")
	isEchoResult := flag.Bool("echoResult", IsResultOutput, "Display found prime numbers")
	flag.Parse()

	primeFinder := primesearch.NewPrimeFinder(AreaSize, *isEchoResult)

	inputController := NewInputController(app.reader)
	inputController.Read()

	var prof *profiler.Profiler = profiler.New(*isProfiling)

	for i := 0; i < inputController.CaseCount(); i++ {
		prof.Start()
		primeFinder.RunCase(inputController.RangeByIndex(i), app.writer)
		prof.End()
	}
	prof.PrintMyself(app.writer)
}
