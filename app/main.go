// Package main is the main package to run simulated annealing package.
// It create new instance of Simulated Annealing and configure the options of it.
package main

import (
	"math/rand"
	"time"

	"github.com/pamungkaski/go-simulated-annealing"
)

func main() {
	// Creating Simulated Annealing instance.
	anne := Simulated.NewAnnealing(20, 0.95, -10, 10, -19.2085, 0.0001)
	counter := 0
	solutionCounter := 0
	minimumIterations := 1000

	// Initialization
	rand.Seed(time.Now().UTC().UnixNano())

	firstValue, secondValue := anne.GenerateNeighbor()
	current := anne.CostFunction(firstValue, secondValue)
	best := current

	// Input
	anne.InputBest(counter)
	anne.AppendSolution(firstValue, secondValue, current)

	// Looping
	for anne.KeepGoing() {
		counter++
		// Generate new solution
		firstValue, secondValue := anne.GenerateNeighbor()
		// Generate energy cost
		energy := anne.CostFunction(firstValue, secondValue)
		// Calculate delta energy
		changes := energy - current

		// If changes < 0
		if changes < 0 {
			// Set the new generated energy as current solution
			current = energy
			solutionCounter++
			// Save the solution
			anne.AppendSolution(firstValue, secondValue, energy)

			// If current solution better than the best, set the current as the best
			if best > current {
				best = current
				anne.InputBest(solutionCounter)
			}
		} else {
			// Calculate Acceptance Probability
			P := anne.AcceptanceProbability(current, energy)
			R := rand.Float64()

			// If P >= then random number, set the newly generated as curent
			if P >= R {
				current = energy
			}
		}

		// Decrease the Temperature if minimumIterations already reached.
		if counter > minimumIterations {
			anne.Cooling()
		}
	}
	// Print the Best solution
	anne.BestSolution()
	// Create Visualization of the model.
	anne.Visualizing()
}
