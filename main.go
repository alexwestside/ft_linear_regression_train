package main

import (
	"github.com/ft_linear_regression_train/model"
	"os"
)

func main() {

	var err error
	lnReg := model.NewLnReg()

	err = lnReg.NewCommand().Execute()
	lnReg.ErrorHandler(err)

	os.Exit(0)
}
