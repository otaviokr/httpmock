// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/otaviokr/httpmock/examples"
)

// serve_multipleCmd represents the serve_multiple command
var serve_multipleCmd = &cobra.Command{
	Use:   "serve_multiple",
	Short: "An example on how to use httpmock to serve multiple different pages",
	Long: `Check the code for more information.`,
	Run: func(cmd *cobra.Command, args []string) {
		examples.ServeMultiple()
	},
}

func init() {
	RootCmd.AddCommand(serve_multipleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serve_multipleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serve_multipleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
