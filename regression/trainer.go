package regression

import (
	"math"
	"errors"
)

type Trainer interface {
	Train(df [][]string) (error)
}

func (m *Model) Train() (error) {

	if len(m.DataFrame.DF) == 0 {
		return errors.New("DataFrame is empty...you need read some fo train Model")
	}

	normVals, err := m.Normalize(m.DataFrame.DF)
	if err != nil {
		return err
	}

	var t0, t1 float64
	lRate := 0.1
	for {
		tmpT0, tmpT1 := nextIteration(t0, t1, normVals, lRate)
		if tmpT0 == t0 && tmpT1 == t1 {
			m.DataResult.Teth0 = t0
			m.DataResult.Teth1 = t1
			m.DataResult.Dvi = deviation(t0, t1, normVals)
			return nil
		} else {
			t0, t1 = tmpT0, tmpT1
		}
	}

	return err
}

func nextIteration(t0 float64, t1 float64, df [][]float64, lRate float64) (float64, float64) {

	tmpT0 := 0.0
	tmpT1 := 0.0

	for _, record := range df {
		tmpT0 += (t0 + t1*record[0]) - record[1]
		tmpT1 += ((t0 + t1*record[0]) - record[1]) * record[0]
	}

	t0 -= (lRate * tmpT0) / float64(len(df))
	t1 -= (lRate * tmpT1) / float64(len(df))

	return t0, t1
}

func deviation(t0 float64, t1 float64, df [][]float64) float64 {
	var dvi float64

	for _, record := range df {
		dvi += math.Pow((t0+t1*record[0])-record[1], 2)
	}

	dvi = math.Sqrt(dvi) / float64(len(df)-1)

	return dvi
}
