package regression

func (m *Model) Calculation(path string) *Model {
	ErrorHandler(m.Read(path))
	ErrorHandler(m.Train())
	return m
}
