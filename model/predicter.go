package model

type Predicter interface {
	Predict(val float64) float64
}

func (l *Model) Predict(val float64) float64 {
	normVal := l.NormalizeIT(l.MinVal, l.MaxVal, val)
	return l.Teth0 + l.Teth1*normVal
}
