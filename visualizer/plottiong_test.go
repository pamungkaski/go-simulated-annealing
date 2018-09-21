package visualizer

import "testing"

func TestPlotting_Visualize(t *testing.T) {
	type args struct {
		XList []float64
		YList []float64
		ZList []float64
	}
	tests := []struct {
		name string
		p    *Plotting
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Visualize(tt.args.XList, tt.args.YList, tt.args.ZList)
		})
	}
}
