package regression

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const outFile = "data.yaml"

//type Data struct {
//	Teth0 float64
//	Teth1 float64
//	Min   float64
//	Max   float64
//	Dvi   float64
//}

type Writer interface {
	Write(t0 float64, t1 float64, dvi float64) error
}

func (m *Model) Write() error {

	_, err := os.Create(outFile)
	ErrorHandler(err)

	blob, err := yaml.Marshal(m.DataResult)
	ErrorHandler(err)

	ErrorHandler(ioutil.WriteFile(outFile, blob, 777))

	return nil
}
