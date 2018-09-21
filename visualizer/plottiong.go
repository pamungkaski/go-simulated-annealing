package visualizer

import (
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/palette"
	"gonum.org/v1/plot"
)

type Plotting struct {
	row int
}

func (p *Plotting)Visualize(XList, YList, ZList []float64) {
	list := append(XList, append(YList, ZList...)...)
	matrix := unitGrid{
		mat.NewDense(
			p.row,
			len(XList),
			list,
		),
	}
	levels := []float64{1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5, 9.5, 10.5, 11.5}

	c := plotter.NewContour(matrix, levels, palette.Rainbow(10, palette.Blue, palette.Red, 1, 1, 1))
	c.LineStyles[0].Width *= 5

	plt, _ := plot.New()

	//plt.Add(h)
	plt.Add(c)
	plt.Add(plotter.NewGlyphBoxes())

	plt.X.Padding = 0
	plt.Y.Padding = 0
	plt.X.Max = 3.5
	plt.Y.Max = 2.5
	plt.Save(2000, 2000, "visualization/contour.png")
}
