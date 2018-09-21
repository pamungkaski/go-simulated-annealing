// Package SA is package that will model simulated annealing.
// It include all the main function of Simulated Annealing Algorithm.
package Simulated


import (
	"math"
	"math/rand"
)

// SimulatedAnnealing is the main interface that holds all of the SimulatedAnnealing Business Logic.
type SimulatedAnnealing interface {
	// Cost function is the function that being simulated.
	CostFunction(firstValue, secondValue float64) float64
	// AcceptanceProbability is the function that will calculate which value that is more likely to be the answer.
	AcceptanceProbability(oldCost, newCost float64) float64
	// GenerateNeighbor is a function that will continuesly create another solution based on its current 
	GenerateNeighbor() (float64, float64)
	// Cooling is the function to decrease temperature by alpha percent.
	Cooling()
	//
	Visualizing()
}

type Visualizer interface {
	Visualize(XList, YList, ZList []float64)
}

// Annealing is the struct that implement SimulatedAnnealing interface.
// This struct will hold the temperature, alpha, and value range.
type Annealing struct {
	temperature float64
	alpha float64
	minValue float64
	maxValue float64
	solutionXList []float64
	solutionYList []float64
	solutionZList []float64
	visualizer Visualizer
}

// Cost function is the function that being simulated.
func (a *Annealing) CostFunction(firstValue, secondValue float64) float64 {
	var result float64
	fVSquare := firstValue * firstValue
	sVSquare := secondValue * secondValue

	fVSin := math.Sin(firstValue)
	sVCos := math.Cos(secondValue)

	subtrahend := math.Exp(fVSquare + sVSquare) / math.Pi
	insideExponential := math.Abs(1 - subtrahend)
	exponent := math.Exp(insideExponential)

	result = -1 * math.Abs(fVSin * sVCos * exponent)

	return result
}

// AcceptanceProbability is the function that will calculate which value that is more likely to be the answer.
func (a *Annealing) AcceptanceProbability(oldCost, newCost float64) float64 {
	var result float64

	deltaCost := newCost - oldCost
	result = math.Exp(-1.0 * deltaCost / a.temperature)

	return result
}

// GenerateNeighbor is a function that will continuesly create another solution based on its current 
func (a *Annealing) GenerateNeighbor() (float64, float64) {
	randomNumber := rand.Float64()
	firstValue := randomNumber * (a.maxValue - a.minValue) + a.minValue
	randomNumber = rand.Float64()
	secondValue := randomNumber * (a.maxValue - a.minValue) + a.minValue

	return firstValue, secondValue
}

// Cooling is the function to decrease temperature by alpha percent.
// The idea of this function to limit / stop the neighbor generation.
func (a *Annealing) Cooling() {
	a.temperature *= a.alpha
}

func (a *Annealing) Visualizing() {
	a.visualizer.Visualize(a.solutionXList, a.solutionYList, a.solutionZList)
}
