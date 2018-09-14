package model

import (
	"github.com/spf13/cobra"
	"fmt"
)

const (
	targetFlag      = "target"
	targetFlagShort = "t"
	viewFlag        = "view"
	viewFlagShort   = "v"
	resultFlag      = "result"
	resultFlagShort = "r"
)

type Commander interface {
	NewCommand() error
}

func (l *Model) NewCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "PROJECT: ft_linear_regression (train model)",
		Short: "The Linear Regression train service",
		Long:  "Train service",
		Run:   l.serve,
	}

	flags := cmd.Flags()
	flags.StringP(targetFlag, targetFlagShort, "", "File target with data-set")
	flags.StringP(viewFlag, viewFlagShort, "", "Target folder for graph representation")
	flags.StringP(resultFlag, resultFlagShort, "", "Target folder for write result data")

	return
}

func (l *Model) serve(cmd *cobra.Command, args []string) {
	fmt.Println(cmd.Use)
	fmt.Println(cmd.Short, "\n")

	var df [][]string
	var err error

	if path := cmd.Flag(targetFlag).Value.String(); path != "" {
		df, err = l.Read(path)
		l.ErrorHandler(err)
		fmt.Println(" -> Read dataset from:", path)

		err = l.Train(df)
		if err != nil {
			fmt.Println("ERROR in train process")
			l.ErrorHandler(err)
		} else {
			fmt.Println(" -> Train model with success")

			if res := cmd.Flag(resultFlag).Value.String(); res != "" {
				err = l.Write()
				l.ErrorHandler(err)
				fmt.Println(" -> Write result to:", res)
			} else {
				fmt.Println("Write target not defined")
			}

			if view := cmd.Flag(viewFlag).Value.String(); view != "" {
				err = l.View(df)
				l.ErrorHandler(err)
				fmt.Println(" -> Create view into:", view)
			} else {
				fmt.Println("View target not defined")
			}
		}
	} else {
		fmt.Println("Read target not defined")
	}

}
