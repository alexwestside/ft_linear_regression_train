package model

import (
	"io"
)

type Model interface {
	Reader() ([][]string, error)
	Writer(t0 float64, t1 float64, dvi float64)
	Normalizer(df [][]string) ([][]float64, error)
	Viewer(df [][]string)
	Trainer(df [][]string) (float64, float64, float64, error)
	Closer(w io.Closer)
	Logger()
}

type Lnreg struct {
	File    string
	Column  int
	Headers []string
	Teth0   float64
	Teth1   float64
	MaxVal  float64
	MinVal  float64
}

func NewLnReg(pathTodataSet string) *Lnreg {
	return &Lnreg{
		File: pathTodataSet,
	}
}
