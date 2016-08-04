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

// serve_genericCmd represents the serve_generic command
var serve_genericCmd = &cobra.Command{
	Use:   "serve_generic",
	Short: "Return always the same response, regardless the address requested.",
	Long: `It will send requests to 3 different URLS, and it will receive alwyas the same response:
	URLs: "/", "/anythingGoes" and "/another/example/to/test.html"

	The response is always: "they all answer the same!"`,
	Run: func(cmd *cobra.Command, args []string) {
		examples.ServeGeneric()
	},
}

func init() {
	RootCmd.AddCommand(serve_genericCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serve_genericCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serve_genericCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
