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

	"github.com/nexmo-community/nexmo-go"
	"github.com/spf13/cobra"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify the user owns a phone number",
	Long: `Verify is a great way to implement 2FA or step-up
auth in your applications`,
}

// RequestId identifies the verify request once it has been created
var RequestId string

// Number is the phone number to confirm ownership of or access to, in E.164 format
var Number string

// Brand name to say who the message is from
var Brand string

// Code supplied by the user
var Code string

var verifyRequestCmd = &cobra.Command{
	Use:   "start",
	Short: "Verify a user's phone number",
	Long:  `Start the verification request for the user's phone number by sending them a pin`,

	Run: func(cmd *cobra.Command, args []string) {
		auth := nexmo.CreateAuthFromKeySecret(Key, Secret)
		verifyClient := nexmo.NewVerifyClient(auth)

		response, respErr, err := verifyClient.Request(Number, Brand, nexmo.VerifyOpts{})

		if err != nil {
			panic(err)
		}

		if respErr.ErrorText != "" {
			fmt.Println("Error status " + respErr.Status + ": " + respErr.ErrorText)
			if respErr.RequestId != "" {
				// the concurrent requests error returns the in-progress ID
				fmt.Println("Request ID: " + respErr.RequestId)
			}
		} else {
			fmt.Println("Request ID: " + response.RequestId)
		}

	},
}

var verifyCheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Check the PIN code sent by the user",
	Long:  `Check the PIN code that the user sent matches what Nexmo sent`,

	Run: func(cmd *cobra.Command, args []string) {
		auth := nexmo.CreateAuthFromKeySecret(Key, Secret)
		verifyClient := nexmo.NewVerifyClient(auth)

		response, respErr, err := verifyClient.Check(RequestId, Code)

		if err != nil {
			panic(err)
		}

		if respErr.ErrorText != "" {
			fmt.Println("Error status " + respErr.Status + ": " + respErr.ErrorText)
			if respErr.RequestId != "" {
				// the concurrent requests error returns the in-progress ID
				fmt.Println("Request ID: " + respErr.RequestId)
			}
		} else {
			fmt.Println("Request completed (Request ID: " + response.RequestId + ")")
		}

	},
}

var verifySearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Find information about a request",
	Long:  `Search for a Request ID and find out its status, timings and progress`,

	Run: func(cmd *cobra.Command, args []string) {
		auth := nexmo.CreateAuthFromKeySecret(Key, Secret)
		verifyClient := nexmo.NewVerifyClient(auth)

		response, respErr, err := verifyClient.Search(RequestId)

		if err != nil {
			panic(err)
		}

		if respErr.ErrorText != "" {
			fmt.Println("Error status " + respErr.Status + ": " + respErr.ErrorText)
			if respErr.RequestId != "" {
				// the concurrent requests error returns the in-progress ID
				fmt.Println("Request ID: " + respErr.RequestId)
			}
		} else {
			fmt.Println("Request ID: " + response.RequestId)
			fmt.Println("Account ID: " + response.AccountId)
			fmt.Println("Status: " + response.Status)
			fmt.Println("Number: " + response.Number)
			fmt.Println("Date Submitted: " + response.DateSubmitted)
			fmt.Println("Date Finalzed: " + response.DateFinalized)
		}

	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)
	verifyCmd.AddCommand(verifyRequestCmd)
	verifyCmd.AddCommand(verifyCheckCmd)
	verifyCmd.AddCommand(verifySearchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// smsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// smsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	verifyRequestCmd.Flags().StringVarP(&Number, "number", "n", "", "Number to verify (E164 format)")
	verifyRequestCmd.MarkFlagRequired("number")
	verifyRequestCmd.Flags().StringVarP(&Brand, "brand", "", "Verify", "Brand identity doing the verifying (default Verify)")
	verifyCheckCmd.Flags().StringVarP(&RequestId, "request-id", "r", "", "Request ID to check the code for")
	verifyCheckCmd.MarkFlagRequired("request-id")
	verifyCheckCmd.Flags().StringVarP(&Code, "code", "", "Verify", "Code to check")
	verifyCheckCmd.MarkFlagRequired("code")
	verifySearchCmd.Flags().StringVarP(&RequestId, "request-id", "r", "", "Request ID to check the code for")
	verifySearchCmd.MarkFlagRequired("request-id")
}
