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
	"encoding/binary"
	"fmt"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/cobra"
)

var systemKey string
var systemToken string
var clientID string

// authCbCmd represents the authCb command
var authCbCmd = &cobra.Command{
	Use:   "auth-cb",
	Short: "Authenticate on ClearBlade platform",
	Long:  `Subscribe to auth topic on ClearBlade platform to obtain a valid auth token`,
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			fmt.Println("sendData called")
		}
		opts := MQTT.NewClientOptions()
		opts.AddBroker(broker)
		opts.SetClientID(clientID)
		opts.SetUsername(systemKey)
		opts.SetPassword(systemToken)
		opts.SetCleanSession(cleansess)
		opts.SetStore(MQTT.NewFileStore(":memory:"))
		choke := make(chan [][]byte)
		opts.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
			choke <- [][]byte{[]byte(msg.Topic()), msg.Payload()}
		})

		client := MQTT.NewClient(opts)
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

		if token := client.Subscribe(topic, byte(qos), nil); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}
		incoming := <-choke
		tokLen := binary.BigEndian.Uint16(incoming[1][:2])
		tok := string(incoming[1][2 : tokLen+2])
		if verbose {
			fmt.Printf("RECEIVED TOPIC: %s MESSAGE: %s\n", incoming[0], tok)
		} else {
			fmt.Println(tok)
		}
		client.Disconnect(250)
		if verbose {
			fmt.Println("Sample Subscriber Disconnected")
		}
	},
}

func init() {
	rootCmd.AddCommand(authCbCmd)

	authCbCmd.Flags().StringVarP(&broker, "broker", "b", "", "Broker URI (required)")
	authCbCmd.MarkFlagRequired("broker")
	authCbCmd.Flags().StringVarP(&systemKey, "system-key", "", "", "System Key (required)")
	authCbCmd.MarkFlagRequired("system-key")
	authCbCmd.Flags().StringVarP(&systemToken, "system-token", "", "", "System Token (required)")
	authCbCmd.MarkFlagRequired("system-token")
	authCbCmd.Flags().StringVarP(&clientID, "client-id", "", "", "ClientID composed of User email and password, e.g.: email:password (required)")
	authCbCmd.MarkFlagRequired("client-id")
	authCbCmd.Flags().BoolVarP(&cleansess, "cleansess", "", false, "Set Clean Session, default to false (optional)")
	authCbCmd.Flags().IntVarP(&qos, "qos", "", 0, "Set QOS to either 0 1 or 2, default to 0 (optional)")
}
