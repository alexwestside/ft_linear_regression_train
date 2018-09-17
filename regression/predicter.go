package regression

type Predicter interface {
	Predict(val float64) float64
}

func (m *Model) Predict(val float64) float64 {
	normVal := m.NormalizeIT(m.DataResult.MinVal, m.DataResult.MaxVal, val)
	return m.DataResult.Teth0 + m.DataResult.Teth1*normVal
}
