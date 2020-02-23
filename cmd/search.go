// Copyright © 2020 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Searches duckduckgo for your tings",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		htmlSearch(combineStrings(args))
	},
}

func combineStrings(args []string) string {
	var combinedString string = ""
	// combine all the arguements into one string
	for _, str := range args {
		// append the arguements into one string
		combinedString += str
	}
	return combinedString
}
func htmlSearch(query string) {
	var baseURL string = "https://duckduckgo.com/html/?q=" + query
	soup.Header("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
	soup.Header("Accept-Language", "en-CA,en-US;q=0.7,en;q=0.3")
	resp, err := soup.Get(baseURL)
	if err != nil {
		panic(err.Error())
	}
	doc := soup.HTMLParse(resp)
	// resultsTitle := doc.FindAll("div", "class", "result")[0].Find("div", "class", "result__body").Find("a", "class", "result__a").FullText()
	// resultsURL := strings.TrimSpace(doc.FindAll("div", "class", "result")[0].Find("div", "class", "result__body").Find("a", "class", "result__url").FullText())
	// resultsBody := doc.FindAll("div", "class", "result")[0].Find("div", "class", "result__body").Find("a", "class", "result__snippet").FullText()

	for i, result := range doc.FindAll("div", "class", "result") {
		fmt.Println(i+1, ".")
		if i >= 9 {
			break
		}
		resultTitle := result.Find("div", "class", "result__body").Find("a", "class", "result__a").FullText()
		resultURL := strings.TrimSpace(result.Find("div", "class", "result__body").Find("a", "class", "result__url").FullText())
		resultBody := result.Find("div", "class", "result__body").Find("a", "class", "result__snippet").FullText()

		fmt.Println(resultTitle)
		fmt.Println(resultURL)
		fmt.Println(resultBody)
	}
}
func init() {
	RootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
