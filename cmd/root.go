/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"os"

	"github.com/GuilhermeJSilva/advent-of-code-2021/solutions"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

var allSolutions []func() = []func(){
	solutions.SolveDay1,
	solutions.SolveDay2,
	solutions.SolveDay3,
	solutions.SolveDay4,
	solutions.SolveDay5,
	solutions.SolveDay6,
	solutions.SolveDay7,
	solutions.SolveDay8,
	solutions.SolveDay9,
	solutions.SolveDay10,
	solutions.SolveDay11,
	solutions.SolveDay12,
	solutions.SolveDay13,
	solutions.SolveDay14,
	solutions.SolveDay15,
	solutions.SolveDay16,
	solutions.SolveDay17,
	solutions.SolveDay18,
	solutions.SolveDay19,
	solutions.SolveDay20,
	solutions.SolveDay21,
	solutions.SolveDay22,
	solutions.SolveDay23,
	solutions.SolveDay24,
	solutions.SolveDay25,
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "advent-of-code-2021",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Default")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func createCommand(day int, solution func()) *cobra.Command {
	return &cobra.Command{
		Use: fmt.Sprintf("day%02d", day+1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Solving Day %d\n", day+1)
			solution()
		},
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.advent-of-code-2021.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	for day, solution := range allSolutions {
		rootCmd.AddCommand(createCommand(day, solution))
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".advent-of-code-2021" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".advent-of-code-2021")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
