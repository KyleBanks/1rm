package main

import (
	"flag"
	"fmt"
	"os"
)

const dampeningFactor = 0.033

var (
	weight float64
	reps   int
)

func init() {
	flag.Float64Var(&weight, "weight", 0, "The weight lifted.")
	flag.IntVar(&reps, "reps", 0, "The number of reps completed.")
	flag.Parse()

	if weight <= 0 || reps <= 0 {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	e1rm := weight
	if reps > 1 {
		e1rm = (weight * float64(reps) * dampeningFactor) + weight
	}
	fmt.Println(e1rm)
}
