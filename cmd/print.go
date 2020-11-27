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
	"bufio"
	"fmt"
	bibinfo "lightref/bibentries"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("print called")
		readFile, err := os.Open(viper.GetString("bibliography"))
		if err != nil {
			log.Fatalf("failed to open file: %s", err)
		}

		d := color.New(color.FgGreen, color.Bold)
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		var fileTextLines []string

		for fileScanner.Scan() {
			fileTextLines = append(fileTextLines, fileScanner.Text())
		}

		readFile.Close()

		for _, eachline := range fileTextLines {
			if strings.Contains(eachline, "Title") || strings.Contains(eachline, "Year") || strings.Contains(eachline, "Author") || strings.Contains(eachline, "@") {
				//eachline = strings.Replace(eachline, "@article", "citekey = ", -1)
				eachline = bibinfo.RemoveType(eachline)
				eachline = strings.Replace(eachline, "{", "", -1)
				eachline = strings.Replace(eachline, "}", "", -1)
				if strings.Contains(eachline, "Citekey") {
					d.Printf("\n" + eachline + "\n")
				} else {
					fmt.Println(eachline)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(printCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
