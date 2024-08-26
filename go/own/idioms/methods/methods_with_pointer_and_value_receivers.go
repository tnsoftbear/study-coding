package main

import (
	"fmt"
)

type Runnable interface {
	Run()
}

type Jogger struct {
	Name string
}

func (j *Jogger) Run() {
	fmt.Printf("Jogger %s is running\n", j.Name)
}


type FootballPlayer struct {
	Name string
}

func (f FootballPlayer) Run() {
	fmt.Printf("FootballPlayer %s is running\n", f.Name)
}

func (f FootballPlayer) Rename() {
	f.Name = "xxx"
}

func runRunner(runner Runnable) {
	runner.Run()
}

func main() {
	footballPlayerRonaldo := &FootballPlayer{Name: "Krishjano Ronaldo"}	
	footballPlayerMessi := FootballPlayer{Name: "Leonel Messi"}	
	footballPlayerRonaldo.Rename()			// не виляет на объект, потому что в методах используется value receiver
	// Можно передавать объекты и указатели на объекты, потому что метод Run() в FootballPlayer ожидает value receiver
	runRunner(footballPlayerRonaldo)
	runRunner(footballPlayerMessi)
	runRunner(footballPlayerRonaldo)
	runRunner(&footballPlayerMessi)
	
	joggerVasja := &Jogger{Name: "Vasja"}
	joggerPetja := Jogger{Name: "Petja"}
	runRunner(joggerVasja)
	runRunner(&joggerPetja)
	// cannot use joggerPetja (variable of type Jogger) as Runnable value in argument to runRunner: Jogger does not implement Runnable (method Run has pointer receiver)
	// runRunner(joggerPetja)				// нельзя передать обычный тип, нужен указатель, потому что Jogger работает через pointer receiver
}

