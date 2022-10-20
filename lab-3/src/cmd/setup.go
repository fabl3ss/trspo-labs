package main

import (
	"awesomeProject1/src/config"
	"flag"
	"fmt"
	"log"
)

func setupNumbers(n int) []int {
	numbers := make([]int, 0, n)

	for i := 1; i < n; i++ {
		numbers = append(numbers, i)
	}

	return numbers
}

func isPlotEnable() bool {
	usePlot := flag.Bool(
		config.PlotFlagName,
		config.PlotFlagDefaultValue,
		config.PlotFlagDescription,
	)

	flag.Parse()
	log.Printf("Display output into .png file==[%t]\n", *usePlot)

	return *usePlot
}

func getNumbersAmount() int {
	var value int

	_, err := fmt.Sscan(flag.Arg(0), &value)
	if err != nil {
		log.Fatalln(err)
	}

	return value
}
