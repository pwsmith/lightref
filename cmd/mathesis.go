/*
Copyright © 2020 Peter W. Smith <peter.w.smith16@gmail.com>

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
	bibinfo "lightref/bibentries"

	"github.com/spf13/cobra"
)

// mathesisCmd represents the mathesis command
var mathesisCmd = &cobra.Command{
	Use:   "mastersthesis",
	Short: "Adds a mastersthesis entry to the bibliography",
	Long: `The mastersthesis command asks for the following fields:
    - citekey
    - author
    - title
    - school
    - year
    - type
    - address
    - month
    - note`,
	Run: func(cmd *cobra.Command, args []string) {
		bibinfo.AddMAThesis()
	},
}

func init() {
	addCmd.AddCommand(mathesisCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mathesisCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mathesisCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
