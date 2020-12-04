/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
)

// inbookCmd represents the inbook command
var inbookCmd = &cobra.Command{
	Use:   "inbook",
	Short: "Adds an inbook entry to the bibliography.",
	Long: `The inbook command asks for the following fields:
    - author
    - editor
    - title
    - chapter
    - pages
    - publisher
    - year
    - volume
    - number
    - series
    - type
    - address
    - edition
    - month
    - note`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("inbook called")
	},
}

func init() {
	addCmd.AddCommand(inbookCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inbookCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inbookCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
