package application

import (
	"bufio"
	"os"
	"flag"
	"segmented-sieve-oop/primesearch"
	"segmented-sieve-oop/profiler"
)

const IS_PROFILING_DEF = true
const IS_RESULT_OUTPUT = true
const AREA_SIZE = 1000001

type App struct{
	reader *bufio.Reader
	writer *bufio.Writer
}

func (app *App) Run() {
	app.reader = bufio.NewReader(os.Stdin)
	app.writer = bufio.NewWriter(os.Stdout)

	isProfiling := flag.Bool("profiling", IS_PROFILING_DEF, "Enable time profiling")
	isEchoResult := flag.Bool("echoResult", IS_RESULT_OUTPUT, "Display found prime numbers")
	flag.Parse()

	var primeFinder *primesearch.PrimeFinder = &primesearch.PrimeFinder{}
	primeFinder.Construct(AREA_SIZE, *isEchoResult)

	var inputController *InputController = &InputController{}
	inputController.Construct(app.reader)
	inputController.Read()
	
	var profiler *profiler.Profiler = &profiler.Profiler{}
	profiler.Construct(*isProfiling)

	for i := 0; i < inputController.CaseCount(); i++ {
		profiler.Start()
		primeFinder.RunCase(inputController.RangeByIndex(i), i, app.writer)
		profiler.End()
	}
	profiler.PrintMyself(app.writer)
}
