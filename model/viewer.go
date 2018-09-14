package model

import (
	"gonum.org/v1/plot"
	"log"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"strconv"
	"fmt"
)

type Viewer interface {
	View(df [][]string) error
}

func (l *Model) View(df [][]string) error {

	pts := make(plotter.XYs, len(df))
	ptsPred := make(plotter.XYs, len(df))

	p, err := plot.New()
	if err != nil {
		fmt.Println(err)
	}

	for i := range df {
		val1, val2, err := parseFloat(df[i][0], df[i][1])
		if err != nil {
			log.Fatal(err)
		}
		pts[i].X = val1
		pts[i].Y = val2

		ptsPred[i].X = val1
		ptsPred[i].Y = l.Predict(val1)
	}

	p.X.Label.Text = "km"
	p.Y.Label.Text = "price"

	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(pts)
	if err != nil {
		fmt.Println(err)
	}
	s.GlyphStyle.Radius = vg.Points(2)
	s.GlyphStyle.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}

	//Add the line plot points for the predictions.
	line, e := plotter.NewLine(ptsPred)
	if e != nil {
		fmt.Println(err)
	}
	line.LineStyle.Width = vg.Points(0.5)
	line.LineStyle.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
	line.LineStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}

	//// Save the plot to a PNG file.
	p.Add(s, line)
	if err := p.Save(15*vg.Centimeter, 15*vg.Centimeter, "./graphs/regression.png"); err != nil {
		fmt.Println(err)
	}

	return nil
}

func parseFloat(p1 string, p2 string) (float64, float64, error) {
	var val1 float64
	var val2 float64
	var err error

	val1, err = strconv.ParseFloat(p1, 64)
	if err != nil {
		return val1, val2, err
	}

	val2, err = strconv.ParseFloat(p2, 64)
	if err != nil {
		return val1, val2, err
	}

	return val1, val2, nil
}
