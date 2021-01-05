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
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// sendDataCmd represents the sendData command
var sendDataCmd = &cobra.Command{
	Use:   "send-data",
	Short: "Send parsed CSV file to MQTT broker",
	Long:  `Parse the CSV file to a MQTT-understandable format and then send it to the configured broker`,
	Run: func(cmd *cobra.Command, args []string) {
		opts := MQTT.NewClientOptions()
		opts.AddBroker(broker)
		opts.SetClientID(id)
		opts.SetUsername(user)
		opts.SetPassword(password)
		opts.SetCleanSession(cleansess)
		opts.SetStore(MQTT.NewFileStore(":memory:"))
		client := MQTT.NewClient(opts)
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
		num := 1
		fmt.Println("Sample Publisher Started")
		for i := 0; i < num; i++ {
			fmt.Println("---- doing publish ----")
			token := client.Publish(topic, byte(qos), false, "test "+time.Now().String())
			token.Wait()
		}
		client.Disconnect(250)
		fmt.Println("Sample Publisher Disconnected")
	},
}

func init() {
	rootCmd.AddCommand(sendDataCmd)

	sendDataCmd.Flags().StringVarP(&broker, "broker", "b", "", "Broker URI (required)")
	sendDataCmd.MarkFlagRequired("broker")
	sendDataCmd.Flags().StringVarP(&user, "user", "u", "", "User (optional)")
	sendDataCmd.Flags().StringVarP(&password, "password", "p", "", "Password (optional)")
	sendDataCmd.Flags().StringVarP(&topic, "topic", "t", "", "Topic (required)")
	sendDataCmd.MarkFlagRequired("topic")
	sendDataCmd.Flags().StringVarP(&id, "id", "", "", "ClientID (optional)")
	sendDataCmd.Flags().BoolVarP(&cleansess, "cleansess", "", false, "Set Clean Session, default to false (optional)")
	sendDataCmd.Flags().IntVarP(&qos, "qos", "", 0, "Set QOS to either 0 1 or 2, default to 0 (optional)")

	cobra.OnInitialize(readConfig)
}

func readConfig() {
	if err := viper.ReadInConfig(); err == nil {
		if verbose {
			fmt.Println("Reading config file")
		}
		viper.BindPFlags(rootCmd.Flags())
		sendDataCmd.Flags().VisitAll(func(f *pflag.Flag) {
			fmt.Println(f.Name)
			if viper.IsSet(f.Name) && viper.GetString(f.Name) != "" {
				sendDataCmd.Flags().Set(f.Name, viper.GetString(f.Name))
			}
		})
	}
}
