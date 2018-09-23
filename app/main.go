// Package main is the main package to run simulated annealing package.
// It create new instance of Simulated Annealing and configure the options of it.
package main

import (
	"math/rand"
	"time"

	"github.com/joho/godotenv"

	"github.com/pamungkaski/go-simulated-annealing"
	"os"
	"strconv"
	"fmt"
)

func main() {
	// Input the params from env
	godotenv.Load()
	tempString := os.Getenv("TEMPERATURE")
	alphaString := os.Getenv("ALPHA")
	minIterationString := os.Getenv("MIN_ITERATION")
	minValueString := os.Getenv("MIN_VALUE")
	maxValueString := os.Getenv("MAX_VALUE")
	goalString := os.Getenv("GOAL")
	minTempString := os.Getenv("MIN_TEMPERATURE")
	loopString := os.Getenv("LOOP")

	// Parse env into float
	temp, err := strconv.ParseFloat(tempString, 64)
	if err != nil {
		fmt.Errorf("%s", err.Error())
	}
	alpha, err := strconv.ParseFloat(alphaString, 64)
	if err != nil {
		fmt.Errorf("%s", err.Error())
	}
	minIteration, err := strconv.ParseFloat(minIterationString, 64)
	if err != nil {
		fmt.Errorf("%s", err.Error())
	}
	minValue, err := strconv.ParseFloat(minValueString, 64)
	if err != nil {
		fmt.Errorf("%s", err.Error())
	}
	maxValue, err := strconv.ParseFloat(maxValueString, 64)
	if err != nil {
		fmt.Errorf("%s", err.Error())
	}
	goal, err := strconv.ParseFloat(goalString, 64)
	if err != nil {
		fmt.Errorf("%s", err.Error())
	}
	minTemp, err := strconv.ParseFloat(minTempString, 64)
	if err != nil {
		fmt.Errorf("%s", err.Error())
	}
	loop, err := strconv.Atoi(loopString)
	if err != nil {
		fmt.Errorf("%s", err.Error())
	}
	// Print Params
	fmt.Printf("Temperature: %.2f\n", temp)
	fmt.Printf("Alpha: %.2f\n", alpha)
	fmt.Printf("Minimum Iteration: %.2f\n", minIteration)
	fmt.Printf("Minimum Value: %.2f\n", minValue)
	fmt.Printf("Maximum Value: %.2f\n", maxValue)
	fmt.Printf("Goal: %.6f\n", goal)
	fmt.Printf("Minimum Temperature: %.4f\n", minTemp)
	fmt.Printf("Loop: %d\n", loop)

	for i := 0; i < loop; i++{
		// Creating Simulated Annealing instance.
		anne := Simulated.NewAnnealing(temp, alpha, minValue, maxValue, goal, minTemp)
		counter := 0
		solutionCounter := 0
		minimumIterations := minIteration

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
}
