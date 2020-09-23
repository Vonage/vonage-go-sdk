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

// Package cmd is an example app
package cmd

import (
	"fmt"
	"time"

	"github.com/vonage/vonage-go-sdk/jwt"
	"github.com/spf13/cobra"
)

// jwtCmd represents the jwt command
var jwtCmd = &cobra.Command{
	Use:   "jwt",
	Short: "JWT features for Vonage applications",
	Long:  `Some APIs such as Voice API, Applications API use JWT authentication, so here are some auxiliary features to help integegrate with those endpoints`,
}

// ApplicationId is the app that the key is issued for
var ApplicationId string

// TTL is the validity period of the generated token, in minutes
var Ttl int

// PrivateKeyFile is a path to the private key **TODO** What about env var support?
var PrivateKeyFile string

var jwtGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a JWT to use with an application",
	Long:  `Generate a JWT for an application from private key and application ID, with optional expiry (the default is 15 minutes)`,

	Run: func(cmd *cobra.Command, args []string) {
		g, ferr := jwt.NewGeneratorFromFilename(ApplicationId, PrivateKeyFile)
		if ferr != nil {
			panic(ferr)
		}

		g.TTL = time.Minute * time.Duration(Ttl)
		token, err := g.GenerateToken()
		if err != nil {
			panic(err)
		}

		fmt.Println(token)
	},
}

func init() {
	rootCmd.AddCommand(jwtCmd)
	jwtCmd.AddCommand(jwtGenerateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// smsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// smsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	jwtGenerateCmd.Flags().StringVarP(&ApplicationId, "application-id", "a", "", "Application ID to generate a token for")
	jwtGenerateCmd.MarkFlagRequired("application-id")
	jwtGenerateCmd.Flags().StringVarP(&PrivateKeyFile, "private-key-file", "f", "", "Private key file to sign the key with")
	jwtGenerateCmd.MarkFlagRequired("private-key-file")
	jwtGenerateCmd.Flags().IntVarP(&Ttl, "ttl", "t", 15, "Time to live (TTL) - how long the token should be valid for in minutes (default: 15)")
}
