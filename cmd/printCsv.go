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
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// printCsvCmd represents the printCsv command
var printCsvCmd = &cobra.Command{
	Use:   "print-csv [CSV_FILES...]",
	Short: "Print CSV File line-by-line",
	Long: `Print CSV File line-by-line in a key:value format
Used for debug purposes and double-check the data`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires at least one file argument")
		}
		for i := 0; i < len(args); i++ {
			if _, err := os.Stat(args[i]); err == nil {
				continue
			} else if os.IsNotExist(err) {
				return fmt.Errorf("%s doesn't exist or can't be read", args[i])
			} else {
				return fmt.Errorf("unexpected error while opening %s", args[i])
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < len(args); i++ {
			// Open the file
			csvfile, err := os.Open(args[i])
			if err != nil {
				log.Fatalln("Couldn't open the csv file", err)
			}

			// Parse the file
			r := csv.NewReader(csvfile)
			//r := csv.NewReader(bufio.NewReader(csvfile))

			// Iterate through the records
			for {
				// Read each record from csv
				record, err := r.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(record)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(printCsvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printCsvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printCsvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
