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
	"github.com/spf13/viper"
	//	"github.com/spf13/viper"
	//	"io/ioutil"
	//	"log"
)

func Message() {
	fmt.Println("Hello world")
}

func MessageDef() {
	fmt.Println("This is the default message")
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		fmt.Println("####### This is the landing page for add #######")
		//	var bibliography string = viper.GetString("bibliography")
		//	fmt.Println(bibliography)
		//	data, err := ioutil.ReadFile(bibliography)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//	fmt.Println(string(data))

		tstatus, _ := cmd.Flags().GetBool("toggle")
		if tstatus {
			Message()
		} else {
			MessageDef()
		}
		fmt.Println(cmd.Flags().GetString("bibliography"))
		bib, _ := cmd.Flags().GetString("bibliography")
		if bib != "" {
			fmt.Println("Bibliography =", bib)
			viper.Set("bibliography", bib)
		}
		fmt.Println("File chosen by viper is", viper.GetString("bibliography"))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
