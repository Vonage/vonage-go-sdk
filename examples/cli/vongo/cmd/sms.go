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

// smsCmd represents the sms command
var smsCmd = &cobra.Command{
	Use:   "sms",
	Short: "SMS features and configuration",
	Long: `SMS is one of the simplest messaging features, go ahead
use these features to get started and/or test your setup`,
}

// To sets who to send the message to
var To string

// From sets who the sender should be (restrictions vary by geography/carrier)
var From string

// Message to send
var Message string

var smsSendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send an SMS",
	Long: `SMS is one of the simplest messaging features, go ahead
use these features to get started and/or test your setup`,

	Run: func(cmd *cobra.Command, args []string) {
		auth := nexmo.CreateAuthFromKeySecret(Key, Secret)
		smsClient := nexmo.NewSMSClient(auth)

		response, err := smsClient.Send(From, To, Message, nexmo.SMSOpts{})

		if err != nil {
			panic(err)
		}

		fmt.Println("I'm sending an SMS from " + From + " to " + To)
		fmt.Println("It says: " + Message)
		fmt.Println("Status: " + response.Messages[0].Status)
	},
}

func init() {
	rootCmd.AddCommand(smsCmd)
	smsCmd.AddCommand(smsSendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// smsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// smsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	smsSendCmd.Flags().StringVarP(&To, "to", "", "", "Number to send the message to (E164 format)")
	smsSendCmd.MarkFlagRequired("to")
	smsSendCmd.Flags().StringVarP(&From, "from", "", "Vongo", "Message sender (phone in E164 format, alphanumeric allowed sometimes)")
	smsSendCmd.Flags().StringVarP(&Message, "text", "", "", "Message to send")
	smsSendCmd.MarkFlagRequired("text")
}
