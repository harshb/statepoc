package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"wf/pkg"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get tokens",
	Long:  `get tokens. This command is used to get the list of outstanding tokens for a workflow execution`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
		GetTokens()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

}
func GetTokens() error {
	messages, err := pkg.GetAllMessagesFromSQSQueue()
	if err != nil {
		fmt.Println("Error getting messages:", err)
		return err
	}

	for _, message := range messages {
		fmt.Println("Task Name:", message.TaskName)
		fmt.Println("Task Token:", message.TaskToken)
	}
	return nil
}
