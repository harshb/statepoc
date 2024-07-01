/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"wf/pkg"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts a new workflow",
	Long:  `Starts a new workflow by creating a new execution of the State Machine`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
		err := StartWorkFlow()
		if err != nil {
			fmt.Println(err)

		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

}
func StartWorkFlow() error {

	fmt.Println("Setting up state machine")
	err := pkg.CreateActivityStateMachine()
	if err != nil {
		return err
	}
	fmt.Println("Created state machine")

	return nil
}
