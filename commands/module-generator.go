package commands

import (
	"github.com/spf13/cobra"
	"log"
	"module-generator/view"
)

func InitCli() {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(initModuleGenerator())

	err := rootCmd.Execute()

	if err != nil {
		log.Fatal("Error on execute CMD EXECUTE")
	}
}

func initModuleGenerator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "generate a new android module with project",
		Run: func(cmd *cobra.Command, args []string) {
			view.IGeneratorView.InitView(&view.PromptuiView{})
		},
	}

	return cmd
}
