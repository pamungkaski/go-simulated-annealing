package main

import (
	"math/rand"
	"time"

	"github.com/pamungkaski/go-simulated-annealing"
)

func main() {
	// Creating Simulated Annealing instance.
	anne := Simulated.NewAnnealing(20, 0.95, -10, 10, -19.2085)
	counter := 0
	solutionCounter := 0
	minimumIterations := 1000

	// Initialization
	rand.Seed(time.Now().UTC().UnixNano())

	firstValue, secondValue := anne.GenerateNeighbor()
	current := anne.CostFunction(firstValue, secondValue)
	best := current

	anne.InputBest(counter)
	anne.AppendSolution(firstValue, secondValue, current)

	// Looping
	for anne.KeepGoing() {
		counter++
		firstValue, secondValue := anne.GenerateNeighbor()
		energy := anne.CostFunction(firstValue, secondValue)
		changes := energy - current

		if changes < 0 {
			current = energy
			solutionCounter++
			anne.AppendSolution(firstValue, secondValue, energy)

			if best > current {
				best = current
				anne.InputBest(solutionCounter)
			}
		} else {
			P := anne.AcceptanceProbability(current, energy)
			R := rand.Float64()
			if P >= R {
				current = energy
			}
		}

		if counter > minimumIterations {
			anne.Cooling()
		}
	}
	anne.BestSolution()
	anne.Visualizing()
}
