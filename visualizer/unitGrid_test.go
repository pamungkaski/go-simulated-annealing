package visualizer

import "testing"

func Test_unitGrid_Dims(t *testing.T) {
	tests := []struct {
		name  string
		g     unitGrid
		wantC int
		wantR int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotC, gotR := tt.g.Dims()
			if gotC != tt.wantC {
				t.Errorf("unitGrid.Dims() gotC = %v, want %v", gotC, tt.wantC)
			}
			if gotR != tt.wantR {
				t.Errorf("unitGrid.Dims() gotR = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func Test_unitGrid_Z(t *testing.T) {
	type args struct {
		c int
		r int
	}
	tests := []struct {
		name string
		g    unitGrid
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.Z(tt.args.c, tt.args.r); got != tt.want {
				t.Errorf("unitGrid.Z() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unitGrid_X(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		g    unitGrid
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.X(tt.args.c); got != tt.want {
				t.Errorf("unitGrid.X() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unitGrid_Y(t *testing.T) {
	type args struct {
		r int
	}
	tests := []struct {
		name string
		g    unitGrid
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.Y(tt.args.r); got != tt.want {
				t.Errorf("unitGrid.Y() = %v, want %v", got, tt.want)
			}
		})
	}
}
