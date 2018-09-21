package visualizer

import "gonum.org/v1/gonum/mat"

type unitGrid struct{ mat.Matrix }

func (g unitGrid) Dims() (c, r int) {
	r, c = g.Matrix.Dims()
	return c, r
}

func (g unitGrid) Z(c, r int) float64 {
	return g.Matrix.At(r, c)
}

func (g unitGrid) X(c int) float64 {
	_, n := g.Matrix.Dims()
	if c < 0 || c >= n {
		panic("index out of range")
	}
	return float64(c)
}

func (g unitGrid) Y(r int) float64 {
	m, _ := g.Matrix.Dims()
	if r < 0 || r >= m {
		panic("index out of range")
	}
	return float64(r)
}
