package main

import (
	"awesomeProject1/src/pkg/collatz"
	plot2 "awesomeProject1/src/pkg/plot"
	"log"
	"sync"

	"gonum.org/v1/plot"
)

func main() {
	isPlotOutput := isPlotEnable()
	numbersAmount := getNumbersAmount()
	numbers := setupNumbers(numbersAmount)

	var p *plot.Plot
	if isPlotOutput {
		p = plot2.CreatePlot()
	}

	wg := sync.WaitGroup{}
	c := make(chan []int, numbersAmount)
	for _, v := range numbers {
		go logic(v, c, &wg)
		steps := <-c

		log.Printf("Collatz for [%d] took [%d] steps\n", v, len(steps))
		if isPlotOutput {
			go plot2.PrintPlotValues(&wg, p, steps)
		}
	}
	wg.Wait()

	if isPlotOutput {
		plot2.SavePlot(p)
	}
}

func logic(newValue int, c chan []int, wg *sync.WaitGroup) {
	wg.Add(1)
	c <- collatz.Collatz(newValue)
	wg.Done()
}
