package cmd

import (
	"fmt"
	"wf/pkg"

	"github.com/spf13/cobra"
)

// purgeCmd represents the purge command
var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "purges sqs queue",
	Long:  `purges sqs queue. This command is used to purge the sqs queue of all messages.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("purge called")
		err := pkg.PurgeSQSQueue()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("purged queue")
	},
}

func init() {
	rootCmd.AddCommand(purgeCmd)

}
