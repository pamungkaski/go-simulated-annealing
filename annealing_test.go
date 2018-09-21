// Package SA is package that will model simulated annealing.
// It include all the main function of Simulated Annealing Algorithm.
package Simulated

import (
	"reflect"
	"testing"
)

func TestNewAnnealing(t *testing.T) {
	type args struct {
		temperature        float64
		alpha              float64
		minValue           float64
		maxValue           float64
		goal               float64
		minimumTemperature float64
	}
	tests := []struct {
		name string
		args args
		want SimulatedAnnealing
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAnnealing(tt.args.temperature, tt.args.alpha, tt.args.minValue, tt.args.maxValue, tt.args.goal, tt.args.minimumTemperature); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAnnealing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnnealing_KeepGoing(t *testing.T) {
	type fields struct {
		temperature        float64
		alpha              float64
		minValue           float64
		maxValue           float64
		goal               float64
		minimumTemperature float64
		solutionXList      []float64
		solutionYList      []float64
		solutionZList      []float64
		best               int
		visualizer         Visualizer
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Annealing{
				temperature:        tt.fields.temperature,
				alpha:              tt.fields.alpha,
				minValue:           tt.fields.minValue,
				maxValue:           tt.fields.maxValue,
				goal:               tt.fields.goal,
				minimumTemperature: tt.fields.minimumTemperature,
				solutionXList:      tt.fields.solutionXList,
				solutionYList:      tt.fields.solutionYList,
				solutionZList:      tt.fields.solutionZList,
				best:               tt.fields.best,
				visualizer:         tt.fields.visualizer,
			}
			if got := a.KeepGoing(); got != tt.want {
				t.Errorf("Annealing.KeepGoing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnnealing_CostFunction(t *testing.T) {
	type fields struct {
		temperature        float64
		alpha              float64
		minValue           float64
		maxValue           float64
		goal               float64
		minimumTemperature float64
		solutionXList      []float64
		solutionYList      []float64
		solutionZList      []float64
		best               int
		visualizer         Visualizer
	}
	type args struct {
		firstValue  float64
		secondValue float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Annealing{
				temperature:        tt.fields.temperature,
				alpha:              tt.fields.alpha,
				minValue:           tt.fields.minValue,
				maxValue:           tt.fields.maxValue,
				goal:               tt.fields.goal,
				minimumTemperature: tt.fields.minimumTemperature,
				solutionXList:      tt.fields.solutionXList,
				solutionYList:      tt.fields.solutionYList,
				solutionZList:      tt.fields.solutionZList,
				best:               tt.fields.best,
				visualizer:         tt.fields.visualizer,
			}
			if got := a.CostFunction(tt.args.firstValue, tt.args.secondValue); got != tt.want {
				t.Errorf("Annealing.CostFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnnealing_AcceptanceProbability(t *testing.T) {
	type fields struct {
		temperature        float64
		alpha              float64
		minValue           float64
		maxValue           float64
		goal               float64
		minimumTemperature float64
		solutionXList      []float64
		solutionYList      []float64
		solutionZList      []float64
		best               int
		visualizer         Visualizer
	}
	type args struct {
		oldCost float64
		newCost float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Annealing{
				temperature:        tt.fields.temperature,
				alpha:              tt.fields.alpha,
				minValue:           tt.fields.minValue,
				maxValue:           tt.fields.maxValue,
				goal:               tt.fields.goal,
				minimumTemperature: tt.fields.minimumTemperature,
				solutionXList:      tt.fields.solutionXList,
				solutionYList:      tt.fields.solutionYList,
				solutionZList:      tt.fields.solutionZList,
				best:               tt.fields.best,
				visualizer:         tt.fields.visualizer,
			}
			if got := a.AcceptanceProbability(tt.args.oldCost, tt.args.newCost); got != tt.want {
				t.Errorf("Annealing.AcceptanceProbability() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnnealing_GenerateNeighbor(t *testing.T) {
	type fields struct {
		temperature        float64
		alpha              float64
		minValue           float64
		maxValue           float64
		goal               float64
		minimumTemperature float64
		solutionXList      []float64
		solutionYList      []float64
		solutionZList      []float64
		best               int
		visualizer         Visualizer
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
		want1  float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Annealing{
				temperature:        tt.fields.temperature,
				alpha:              tt.fields.alpha,
				minValue:           tt.fields.minValue,
				maxValue:           tt.fields.maxValue,
				goal:               tt.fields.goal,
				minimumTemperature: tt.fields.minimumTemperature,
				solutionXList:      tt.fields.solutionXList,
				solutionYList:      tt.fields.solutionYList,
				solutionZList:      tt.fields.solutionZList,
				best:               tt.fields.best,
				visualizer:         tt.fields.visualizer,
			}
			got, got1 := a.GenerateNeighbor()
			if got != tt.want {
				t.Errorf("Annealing.GenerateNeighbor() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Annealing.GenerateNeighbor() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAnnealing_Cooling(t *testing.T) {
	type fields struct {
		temperature        float64
		alpha              float64
		minValue           float64
		maxValue           float64
		goal               float64
		minimumTemperature float64
		solutionXList      []float64
		solutionYList      []float64
		solutionZList      []float64
		best               int
		visualizer         Visualizer
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Annealing{
				temperature:        tt.fields.temperature,
				alpha:              tt.fields.alpha,
				minValue:           tt.fields.minValue,
				maxValue:           tt.fields.maxValue,
				goal:               tt.fields.goal,
				minimumTemperature: tt.fields.minimumTemperature,
				solutionXList:      tt.fields.solutionXList,
				solutionYList:      tt.fields.solutionYList,
				solutionZList:      tt.fields.solutionZList,
				best:               tt.fields.best,
				visualizer:         tt.fields.visualizer,
			}
			a.Cooling()
		})
	}
}

func TestAnnealing_Visualizing(t *testing.T) {
	type fields struct {
		temperature        float64
		alpha              float64
		minValue           float64
		maxValue           float64
		goal               float64
		minimumTemperature float64
		solutionXList      []float64
		solutionYList      []float64
		solutionZList      []float64
		best               int
		visualizer         Visualizer
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Annealing{
				temperature:        tt.fields.temperature,
				alpha:              tt.fields.alpha,
				minValue:           tt.fields.minValue,
				maxValue:           tt.fields.maxValue,
				goal:               tt.fields.goal,
				minimumTemperature: tt.fields.minimumTemperature,
				solutionXList:      tt.fields.solutionXList,
				solutionYList:      tt.fields.solutionYList,
				solutionZList:      tt.fields.solutionZList,
				best:               tt.fields.best,
				visualizer:         tt.fields.visualizer,
			}
			a.Visualizing()
		})
	}
}

func TestAnnealing_AppendSolution(t *testing.T) {
	type fields struct {
		temperature        float64
		alpha              float64
		minValue           float64
		maxValue           float64
		goal               float64
		minimumTemperature float64
		solutionXList      []float64
		solutionYList      []float64
		solutionZList      []float64
		best               int
		visualizer         Visualizer
	}
	type args struct {
		firstValue  float64
		secondValue float64
		energy      float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Annealing{
				temperature:        tt.fields.temperature,
				alpha:              tt.fields.alpha,
				minValue:           tt.fields.minValue,
				maxValue:           tt.fields.maxValue,
				goal:               tt.fields.goal,
				minimumTemperature: tt.fields.minimumTemperature,
				solutionXList:      tt.fields.solutionXList,
				solutionYList:      tt.fields.solutionYList,
				solutionZList:      tt.fields.solutionZList,
				best:               tt.fields.best,
				visualizer:         tt.fields.visualizer,
			}
			a.AppendSolution(tt.args.firstValue, tt.args.secondValue, tt.args.energy)
		})
	}
}

func TestAnnealing_InputBest(t *testing.T) {
	type fields struct {
		temperature        float64
		alpha              float64
		minValue           float64
		maxValue           float64
		goal               float64
		minimumTemperature float64
		solutionXList      []float64
		solutionYList      []float64
		solutionZList      []float64
		best               int
		visualizer         Visualizer
	}
	type args struct {
		it int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Annealing{
				temperature:        tt.fields.temperature,
				alpha:              tt.fields.alpha,
				minValue:           tt.fields.minValue,
				maxValue:           tt.fields.maxValue,
				goal:               tt.fields.goal,
				minimumTemperature: tt.fields.minimumTemperature,
				solutionXList:      tt.fields.solutionXList,
				solutionYList:      tt.fields.solutionYList,
				solutionZList:      tt.fields.solutionZList,
				best:               tt.fields.best,
				visualizer:         tt.fields.visualizer,
			}
			a.InputBest(tt.args.it)
		})
	}
}

func TestAnnealing_BestSolution(t *testing.T) {
	type fields struct {
		temperature        float64
		alpha              float64
		minValue           float64
		maxValue           float64
		goal               float64
		minimumTemperature float64
		solutionXList      []float64
		solutionYList      []float64
		solutionZList      []float64
		best               int
		visualizer         Visualizer
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Annealing{
				temperature:        tt.fields.temperature,
				alpha:              tt.fields.alpha,
				minValue:           tt.fields.minValue,
				maxValue:           tt.fields.maxValue,
				goal:               tt.fields.goal,
				minimumTemperature: tt.fields.minimumTemperature,
				solutionXList:      tt.fields.solutionXList,
				solutionYList:      tt.fields.solutionYList,
				solutionZList:      tt.fields.solutionZList,
				best:               tt.fields.best,
				visualizer:         tt.fields.visualizer,
			}
			a.BestSolution()
		})
	}
}
