/*
Copyright © 2020 Lorna Mitchell <lorna@lornajane.net>

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

// Package cmd is an example app
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Key is the user's API key
var Key string

// Secret is the user's API secret
var Secret string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vongo",
	Short: "A simple demo of a golang SDK for Vonage APIs",
	Long: `This application is a work in progress, it uses an under-construction
	version of the Vonage golang SDK.`,
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
	// cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&Key, "key", "", "API key")
	rootCmd.PersistentFlags().StringVar(&Secret, "secret", "", "API secret")

	rootCmd.Version = "0.0.1"
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	/*
		if cfgFile != "" {
			// Use config file from the flag.
			viper.SetConfigFile(cfgFile)
		} else {
			// Find home directory.
			home, err := homedir.Dir()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Search config in home directory with name ".cli" (without extension).
			viper.AddConfigPath(home)
			viper.SetConfigName(".cli")
		}

		viper.AutomaticEnv() // read in environment variables that match

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	*/
}
