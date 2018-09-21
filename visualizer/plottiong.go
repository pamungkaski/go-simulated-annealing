package visualizer

import (
	"image/color"
	"math"
	"sort"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type Plotting struct{}
type Data struct {
	X float64
	Y float64
	Z float64
}

func (p *Plotting) Visualize(XList, YList, ZList []float64) {
	scatterData := make(plotter.XYZs, len(XList))
	for i := range XList {
		scatterData[i].X = XList[i]
		scatterData[i].Y = YList[i]
		scatterData[i].Z = ZList[i] * -1.0
	}
	sort.Slice(scatterData, func(i, j int) bool {
		return scatterData[i].Z < scatterData[j].Z
	})

	minZ, maxZ := math.Inf(1), math.Inf(-1)
	for _, xyz := range scatterData {
		if xyz.Z > maxZ {
			maxZ = xyz.Z
		}
		if xyz.Z < minZ {
			minZ = xyz.Z
		}
	}

	plt, _ := plot.New()
	plt.Title.Text = "Simulated Annealing"
	plt.X.Label.Text = "X"
	plt.Y.Label.Text = "Y"
	plt.Legend.Add("The bigger the circle, the smaller the value")

	sc, err := plotter.NewScatter(scatterData)
	if err != nil {
		panic(err)
	}
	// Specify style for individual points.
	sc.GlyphStyleFunc = func(i int) draw.GlyphStyle {
		x, y, z := scatterData.XYZ(i)
		c := color.RGBA{R: 190 + uint8(x*20), G: 128 + uint8(y*20), B: 0 + uint8(z*20), A: 255}
		var minRadius, maxRadius = vg.Points(-10), vg.Points(10)
		rng := maxRadius - minRadius
		d := (z - minZ) / (maxZ - minZ)
		r := vg.Length(d) * rng * 2
		return draw.GlyphStyle{Color: c, Radius: r, Shape: draw.CircleGlyph{}}
	}
	plt.Add(sc)

	plt.Save(10*vg.Inch, 10*vg.Inch, "visualization/contour.svg")
}
