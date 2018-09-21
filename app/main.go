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
	// Seed the random number generator
	rand.Seed(time.Now().UTC().UnixNano())

	// Intialize the first solution
	firstValue, secondValue := anne.GenerateNeighbor()
	currentEnergy := anne.CostFunction(firstValue, secondValue)
	best := currentEnergy

	// Input the first solution into the struct
	anne.InputBest(counter)
	anne.AppendSolution(firstValue, secondValue, currentEnergy)

	// Looping
	for anne.KeepGoing() {
		counter++
		// Generate new solution
		firstValue, secondValue := anne.GenerateNeighbor()
		// Generate energy cost
		energy := anne.CostFunction(firstValue, secondValue)
		// Calculate delta energy
		deltaEnergy := energy - currentEnergy

		// If deltaEnergy < 0
		if deltaEnergy < 0 {
			// Set the new generated energy as currentEnergy solution
			currentEnergy = energy
			solutionCounter++
			// Save the solution
			anne.AppendSolution(firstValue, secondValue, energy)

			// If currentEnergy solution better than the best, set the currentEnergy as the best
			if best > currentEnergy {
				best = currentEnergy
				anne.InputBest(solutionCounter)
			}
		} else {
			// Calculate Acceptance Probability
			probability := anne.AcceptanceProbability(currentEnergy, energy)
			randomNumber := rand.Float64()

			// If P >= then random number, set the newly generated as curent
			if probability >= randomNumber {
				currentEnergy = energy
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
