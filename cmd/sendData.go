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
)

var broker string
var port string
var user string
var topic string
var password string
var id int
var cleansess bool
var qos int

// sendDataCmd represents the sendData command
var sendDataCmd = &cobra.Command{
	Use:   "send-data",
	Short: "Send parsed CSV file to MQTT broker",
	Long:  `Parse the CSV file to a MQTT-understandable format and then send it to the configured broker`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sendData called")
	},
}

func init() {
	rootCmd.AddCommand(sendDataCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendDataCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendDataCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
