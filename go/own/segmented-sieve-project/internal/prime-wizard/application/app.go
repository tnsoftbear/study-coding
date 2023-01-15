package application

import (
	"bufio"
	"flag"
	"os"
	"segmented-sieve-project/internal/pkg/profiler"
	"segmented-sieve-project/internal/prime-wizard/domain/primesearch"
	"segmented-sieve-project/internal/prime-wizard/domain/render"
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

	primeFinder := primesearch.NewPrimeFinder(AreaSize)
	basicRenderer := render.NewBasicRenderer(app.writer)

	inputController := NewInputController(app.reader)
	inputController.Read()

	var prof = profiler.New(*isProfiling)

	for i := 0; i < inputController.CaseCount(); i++ {
		prof.Start()
		primes := primeFinder.DetectPrimes(inputController.RangeByIndex(i))
		if *isEchoResult {
			basicRenderer.PrintPrimes(primes)
		}
		prof.End()
	}
	prof.PrintMyself(app.writer)
}
