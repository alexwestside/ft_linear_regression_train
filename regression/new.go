package regression

import "github.com/spf13/cobra"

type ModelData struct {
	File    string
	Column  int
	Headers []string
	DF      [][]string
}

type ModelResult struct {
	Teta0  float64
	Teta1  float64
	Dvi    float64
	MinVal float64
	MaxVal float64
}

type Model struct {
	DataFrame  ModelData
	DataResult ModelResult
}

type Commander interface {
	NewCommand() (cmd *cobra.Command)
}

func NewModel() *Model {
	return &Model{
		DataFrame:  ModelData{},
		DataResult: ModelResult{},
	}
}
