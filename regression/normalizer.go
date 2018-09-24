package regression

import (
	"strconv"
)

func (m *Model) Normalize(df [][]string) ([][]float64, error) {

	dfNew, err := castToFloat(df)
	if err != nil {
		return nil, err
	}

	m.DataResult.MinVal, m.DataResult.MaxVal = calcMinMaxVal(dfNew)

	for i, val := range dfNew {
		dfNew[i][0] = m.NormalizeIT(m.DataResult.MinVal, m.DataResult.MaxVal, val[0])
	}

	return dfNew, nil
}

func castToFloat(df [][]string) ([][]float64, error) {
	var err error
	dfNew := make([][]float64, len(df))
	for i := range dfNew {
		dfNew[i] = make([]float64, 2)
	}

	for i, record := range df {
		for j, val := range record {
			dfNew[i][j], err = strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, err
			}
		}
	}

	return dfNew, nil

}

func (l *Model) NormalizeIT(minVal float64, maxVal float64, val float64) float64 {
	return (val - minVal) / (maxVal - minVal)
}

func calcMinMaxVal(records [][]float64) (float64, float64) {

	minVal := records[0][0]
	maxVal := records[0][0]

	for _, record := range records {
		if record[0] < minVal {
			minVal = record[0]
		}
		if record[0] > maxVal {
			maxVal = record[0]
		}
	}

	return minVal, maxVal
}
