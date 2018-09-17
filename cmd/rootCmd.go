package cmd

import (
	"github.com/spf13/cobra"
)

func RootCmd() (cmd *cobra.Command) {
	rootCmd := &cobra.Command{
		Use:   "PROJECT: ft_linear_regression (train regression)",
		Short: "The Linear Regression train service",
		Long:  "Train service",
	}
	//Commands(rootCmd)
	return rootCmd
}

