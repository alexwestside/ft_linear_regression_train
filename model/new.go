package model

type Model struct {
	File    string
	Column  int
	Headers []string
	Teth0   float64
	Teth1   float64
	Dvi     float64
	MinVal  float64
	MaxVal  float64
	Commander
	Reader
	Normalizer
	Trainer
	Viewer
	Writer
	Predicter
}

func NewLnReg(pathTodataSet string) *Model {
	return &Model{
		File: pathTodataSet,
	}
}
