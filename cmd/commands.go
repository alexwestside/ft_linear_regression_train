package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ft_linear_regression_train/regression"
)

func Commands(rootCmd *cobra.Command) *cobra.Command {
	rootCmd.PersistentFlags().String("t", "", "read target")
	rootCmd.AddCommand(train(), view())
	return rootCmd
}

func train() (cmd *cobra.Command) {
	return &cobra.Command{
		Use:   "train",
		Short: "train",
		Run: func(cmd *cobra.Command, args []string) {
			var path string
			if path = cmd.Flag("t").Value.String(); path == "" {
				handler("target")
			} else {
				regression.ErrorHandler(regression.NewModel().Calculation(path).Write())
			}
		},
	}
}

func view() (cmd *cobra.Command) {
	return &cobra.Command{
		Use:   "view",
		Short: "view",
		Run: func(cmd *cobra.Command, args []string) {
			var path string
			if path = cmd.Flag("t").Value.String(); path == "" {
				handler("target")
			} else {
				regression.ErrorHandler(regression.NewModel().Calculation(path).View())
			}
		},
	}
}
