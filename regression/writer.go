package regression

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const outFile = "data/train_out.yaml"

func (m *Model) Write() error {

	_, err := os.Create(outFile)
	ErrorHandler(err)

	blob, err := yaml.Marshal(m.DataResult)
	ErrorHandler(err)

	ErrorHandler(ioutil.WriteFile(outFile, blob, 777))

	output("TRAIN MODE", "Train model and get Result", outFile)

	return nil
}


