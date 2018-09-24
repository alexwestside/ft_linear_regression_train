package regression

import (
	"gonum.org/v1/plot"
	"log"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"strconv"
	"fmt"
	"errors"
)

const pngFile = "stdout/regression.png"

func (m *Model) View() error {

	if len(m.DataFrame.DF) == 0 {
		return errors.New("DataFrame is empty...you need read some for view Model")
	}

	pts := make(plotter.XYs, len(m.DataFrame.DF))
	ptsPred := make(plotter.XYs, len(m.DataFrame.DF))

	p, err := plot.New()
	if err != nil {
		fmt.Println(err)
	}

	for i := range m.DataFrame.DF {
		val1, val2, err := parseFloat(m.DataFrame.DF[i][0], m.DataFrame.DF[i][1])
		if err != nil {
			log.Fatal(err)
		}
		pts[i].X = val1
		pts[i].Y = val2

		ptsPred[i].X = val1
		ptsPred[i].Y = m.Predict(val1)
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
	if err := p.Save(15*vg.Centimeter, 15*vg.Centimeter, pngFile); err != nil {
		fmt.Println(err)
	}

	fmt.Println(fmt.Sprintf("SUCCESS: Train model and get Graph in %s", pngFile))

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
