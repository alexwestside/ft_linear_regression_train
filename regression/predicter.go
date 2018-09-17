package regression

type Predicter interface {
	Predict(val float64) float64
}

func (m *Model) Predict(val float64) float64 {
	normVal := m.NormalizeIT(m.DataResult.MinVal, m.DataResult.MaxVal, val)
	return m.DataResult.Teta0 + m.DataResult.Teta1*normVal
}
