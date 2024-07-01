package cmd

import (
	"fmt"
	"wf/pkg"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops a running State Machine workflow",
	Long:  `stops a running State Machine workflow by stopping the execution of the State Machine. `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stop called")

		err := pkg.PurgeSQSQueue()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("purged queue")

		err = pkg.DeleteStateMachineAndActivities()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)

}
