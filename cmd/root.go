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
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lightref",
	Short: "",
	Long: `
Lightref is a small, cross platform CLI for adding entries to a bibtex bibliography.
It is designed as a command line based alternative to the GUI bibliograph management alternatives such as Jabref and Bibdesk.
Lightref is **not** intended as a serious competitor to these: they are wonderful programs with vastly more options than Lightref offers.
However, for a simple task such as adding an entry to a bibtex file, opening a full GUI is often overkill, hence Lightref.

Usage:

lightref add <ENTRYTYPE> : add the relevant entry to the bibliography:

    lightref add article ## to add an article
    lightref add book ## add a book entry
    lightref add booklet ## add a booklet entry
    lightref add inbook ## to add an inbook entry
    lightref add incollection ## add an incollection entry
    lightref add inproceedings ## add an inproceedings entry
    lightref add manual ## add a manual entry
    lightref add mastersthesis ## add a mastersthesis entry
    lightref add misc ## add a misc entry
    lightref add phdthesis ## add a phdthesis entry
    lightref add proceedings ## add a proceedings entry
    lightref add techreport ## add a techreport entry
    lightref add unpublished ## add an unpublished entry

lightref print : view a printout of the bibliography

    lightref print
    lightref print -l/--long
    lightref -r/--raw

Configuration:

A default bibliography can be specified in the configuration file, by default $/.lightref.yaml. You can also specify a bibliography per command using the -b/--bibliography 

For full instructions, see https://github.com/pwsmith/lightref.

`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	var Bib string
	rootCmd.PersistentFlags().StringVarP(&Bib, "bibliography", "b", "", "specify an alternate bibliography")

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lightref)")
	//	rootCmd.PersistentFlags().BoolP("bibliography", "b", false, "Specify a different bibiography from the default")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		fmt.Println(cfgFile)
		//		fmt.Println("Config found")
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".lightref" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".lightref")
		//		fmt.Println("Config found")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		bib, _ := rootCmd.Flags().GetString("bibliography")
		if bib != "" {
			fmt.Println("Using the specified bibliography:", bib)
			viper.Set("bibliography", bib)
		} else {
			fmt.Println("Using the default bibliography")
		}
	}

}
