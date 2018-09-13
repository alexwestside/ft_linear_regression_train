package main

import (
	"github.com/ft_linear_regression_train/model"
	"fmt"
)

const pathToDataSrt = "dataset/data.csv"

func main() {

	lnReg := model.NewLnReg(pathToDataSrt)

	errCMD := lnReg.NewCommand()
	if errCMD != nil {
		fmt.Println(errCMD)
	}

	df, errRead := lnReg.Read()
	if errRead != nil {
		fmt.Println(errRead)
	}

	errTrain := lnReg.Train(df)
	if errRead != nil {
		fmt.Println(errTrain)
	}

	errWrite := lnReg.Write()
	if errWrite != nil {
		fmt.Println(errTrain)
	}

	errView := lnReg.View(df)
	if errView != nil {
		fmt.Println(errTrain)
	}

}
