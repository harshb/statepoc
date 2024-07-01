package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"wf/pkg"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("set called")
		SetTokens()
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

}

func SetTokens() error {
	err := pkg.SetTokens()
	if err != nil {
		fmt.Println("Error completing tasks:", err)
	}
	return nil
}
