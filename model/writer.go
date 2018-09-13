package model

import (
	"gopkg.in/yaml.v2"
	"fmt"
	"io/ioutil"
	"os"
)

const outFile = "data.yaml"

type Data struct {
	Teth0 float64
	Teth1 float64
	Min   float64
	Max   float64
	Dvi   float64
}

type Writer interface {
	Write(t0 float64, t1 float64, dvi float64) error
}

func (l *Model) Write() error {

	_, err := os.Create(outFile)
	if err != nil {
		fmt.Println(err)
	}

	blob, err := yaml.Marshal(Data{
		Teth0: l.Teth0,
		Teth1: l.Teth1,
		Max:   l.MaxVal,
		Min:   l.MinVal,
		Dvi:   l.Dvi,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(outFile, blob, 777)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
