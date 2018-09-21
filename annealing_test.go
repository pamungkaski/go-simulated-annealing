// Package SA is package that will model simulated annealing.
// It include all the main function of Simulated Annealing Algorithm.
package Simulated

import (
	"reflect"
	"testing"
)

func TestAnnealing_CostFunction(t *testing.T) {
	type fields struct {
		temperature float64
		alpha       float64
		minValue    float64
		maxValue    float64
		Solutions   []Solution
	}
	type args struct {
		solution Solution
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
				temperature: tt.fields.temperature,
				alpha:       tt.fields.alpha,
				minValue:    tt.fields.minValue,
				maxValue:    tt.fields.maxValue,
				Solutions:   tt.fields.Solutions,
			}
			if got := a.CostFunction(tt.args.solution); got != tt.want {
				t.Errorf("Annealing.CostFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnnealing_AcceptanceProbability(t *testing.T) {
	type fields struct {
		temperature float64
		alpha       float64
		minValue    float64
		maxValue    float64
		Solutions   []Solution
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
				temperature: tt.fields.temperature,
				alpha:       tt.fields.alpha,
				minValue:    tt.fields.minValue,
				maxValue:    tt.fields.maxValue,
				Solutions:   tt.fields.Solutions,
			}
			if got := a.AcceptanceProbability(tt.args.oldCost, tt.args.newCost); got != tt.want {
				t.Errorf("Annealing.AcceptanceProbability() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnnealing_GenerateNeighbor(t *testing.T) {
	type fields struct {
		temperature float64
		alpha       float64
		minValue    float64
		maxValue    float64
		Solutions   []Solution
	}
	tests := []struct {
		name   string
		fields fields
		want   Solution
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Annealing{
				temperature: tt.fields.temperature,
				alpha:       tt.fields.alpha,
				minValue:    tt.fields.minValue,
				maxValue:    tt.fields.maxValue,
				Solutions:   tt.fields.Solutions,
			}
			if got := a.GenerateNeighbor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Annealing.GenerateNeighbor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnnealing_Cooling(t *testing.T) {
	type fields struct {
		temperature float64
		alpha       float64
		minValue    float64
		maxValue    float64
		Solutions   []Solution
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
				temperature: tt.fields.temperature,
				alpha:       tt.fields.alpha,
				minValue:    tt.fields.minValue,
				maxValue:    tt.fields.maxValue,
				Solutions:   tt.fields.Solutions,
			}
			a.Cooling()
		})
	}
}

func TestAnnealing_Visualizing(t *testing.T) {
	type fields struct {
		temperature float64
		alpha       float64
		minValue    float64
		maxValue    float64
		Solutions   []Solution
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
				temperature: tt.fields.temperature,
				alpha:       tt.fields.alpha,
				minValue:    tt.fields.minValue,
				maxValue:    tt.fields.maxValue,
				Solutions:   tt.fields.Solutions,
			}
			a.Visualizing()
		})
	}
}
