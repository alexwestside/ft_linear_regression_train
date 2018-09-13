package main

import (
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	"github.com/ft_linear_regression_train/model"
	"fmt"
)

var rootCmd = &cobra.Command{
	Use:  "mycli-full",
	Long: `SDK Use Demo`,
}


const pathToDataSrt = "dataset/data.csv"

func main() {

	var err error
	lnReg := model.NewLnReg(pathToDataSrt)

	rootCmd.Flags().Int("port", 1138, "Port to run Application server on")
	err = viper.BindPFlag("port", rootCmd.Flags().Lookup("port"))
	lnReg.ErrorHandler(err)

	fmt.Println(viper.Get("port"))



	//
	//err = lnReg.NewCommand()
	//lnReg.ErrorHandler(err)
	//
	//df, err := lnReg.Read()
	//lnReg.ErrorHandler(err)
	//
	//err = lnReg.Train(df)
	//lnReg.ErrorHandler(err)
	//
	//err = lnReg.Write()
	//lnReg.ErrorHandler(err)
	//
	//err = lnReg.View(df)
	//lnReg.ErrorHandler(err)
}
