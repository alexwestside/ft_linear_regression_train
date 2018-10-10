package regression

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"fmt"
)

const outFile = "data/train_out.yaml"

func (m *Model) Write() error {

	_, err := os.Create(outFile)
	ErrorHandler(err)

	blob, err := yaml.Marshal(m.DataResult)
	ErrorHandler(err)

	ErrorHandler(ioutil.WriteFile(outFile, blob, 777))

	//fmt.Println(fmt.Sprintf("SUCCESS: Train model and get Result in %s", outFile))
	output()

	return nil
}


func output() {
	fmt.Println("------------------------------------------------------------------")
	fmt.Print("Liner regression -> TRAIN MODE")
	fmt.Println("------------------------------------------------------------------")
	fmt.Println(fmt.Sprintf("SUCCESS: Train model and get Result in %s", outFile))
	fmt.Println("------------------------------------------------------------------")
}