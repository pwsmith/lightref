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
	Short: "Displays the bibliography in the terminal window.",
	Long:  `Print displays the bibliography in the terminal window. In its default use, it will display a short overview of each entry, with a highlighted citekey, the title of the entry, the author of the entry, and the year of publication. This option is useful for quick reference of the bibliography, particularly when combined with the pagination tool of your operating system, such as 'less'.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("print called")
		raw_value, _ := cmd.Flags().GetBool("raw")
		tstatus, _ := cmd.Flags().GetBool("long")
		readFile, err := os.Open(viper.GetString("bibliography"))
		if err != nil {
			log.Fatalf("failed to open file: %s", err)
		}

		show_green := color.New(color.FgGreen, color.Bold)
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		var fileTextLines []string

		for fileScanner.Scan() {
			fileTextLines = append(fileTextLines, fileScanner.Text())
		}

		defer readFile.Close()
		if raw_value {
			for _, eachline := range fileTextLines {
				fmt.Println(eachline)
			}
		} else if tstatus {
			for _, eachline := range fileTextLines {
				//eachline = strings.Replace(eachline, "@article", "citekey = ", -1)
				eachline = bibinfo.RemoveType(eachline)
				eachline = strings.Replace(eachline, "{", "", -1)
				eachline = strings.Replace(eachline, "}", "", -1)
				eachline = strings.Replace(eachline, " =", ":", -1)
				eachline = strings.TrimSuffix(eachline, ",")
				if strings.Contains(eachline, "Citekey") {
					show_green.Printf("\n" + eachline + "\n")
				} else {
					fmt.Println(eachline)
				}

			}
		} else {
			for _, eachline := range fileTextLines {
				if strings.Contains(eachline, "Title") || strings.Contains(eachline, "title") || strings.Contains(eachline, "Year") || strings.Contains(eachline, "Year") || strings.Contains(eachline, "Author") || strings.Contains(eachline, "author") || strings.Contains(eachline, "@") {
					//eachline = strings.Replace(eachline, "@article", "citekey = ", -1)
					eachline = bibinfo.RemoveType(eachline)
					eachline = strings.Replace(eachline, "{", "", -1)
					eachline = strings.Replace(eachline, "}", "", -1)
					eachline = strings.Replace(eachline, " =", ":", -1)
					eachline = strings.TrimSuffix(eachline, ",")
					if strings.Contains(eachline, "Citekey") {
						show_green.Printf("\n" + eachline + "\n")
					} else {
						fmt.Println(eachline)
					}
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
	printCmd.Flags().BoolP("long", "l", false, "Show the bibliography in longer, but formatted form")
	printCmd.Flags().BoolP("raw", "r", false, "Show the bibliography in raw, unedited form")
}
