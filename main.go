package main

import (
	"github.com/ft_linear_regression_train/cmd"
	"github.com/spf13/cobra"
)

func main() {

	cmd.Commands(&cobra.Command{
		Short: "The Linear Regression train service",
		Long:  "Train service",
	}).Execute()
}
