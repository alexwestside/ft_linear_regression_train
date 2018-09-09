package main

import (
	"github.com/ft_linear_regression_train/model"
	"fmt"
)

func main() {

	pathToDataSrt := "dataset/data.csv"

	lnReg := model.NewLnReg(pathToDataSrt)

	df, errRead := lnReg.Reader()
	if errRead != nil {
		fmt.Println(errRead)
	}

	t0, t1, divi, errTrain := lnReg.Trainer(df)
	if errRead != nil {
		fmt.Println(errTrain)
	}

	fmt.Println(t0, t1, divi)

}
