package model

import (
	"math"
)

type Trainer interface {
	Train(df [][]string) (error)
}

func (l *Model) Train(df [][]string) (error) {

	normVals, err := l.Normalize(df)
	if err != nil {
		return err
	}

	var t0, t1 float64
	lRate := 0.1
	for {
		tmpT0, tmpT1 := nextIteration(t0, t1, normVals, lRate)
		if tmpT0 == t0 && tmpT1 == t1 {
			l.Teth0 = t0
			l.Teth1 = t1
			l.Dvi = deviation(t0, t1, normVals)
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
