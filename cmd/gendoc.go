// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"github.com/apaz037/go-metadata-api/api"
	"github.com/go-chi/docgen"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

var (
	routes bool
)

// gendocCmd represents the gendoc command
var gendocCmd = &cobra.Command{
	Use:   "gendoc",
	Short: "Generates route documentation",
	Long:  `Generates route documentation`,
	Run: func(cmd *cobra.Command, args []string) {
		generateRouteDocs()
	},
}

func init() {
	rootCmd.AddCommand(gendocCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gendocCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gendocCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func generateRouteDocs() {
	api := api.New()
	fmt.Println("generating routes.md file: ")

	md := docgen.MarkdownRoutesDoc(api, docgen.MarkdownOpts{
		ProjectPath: "github.com/apaz037/go-metadata-api",
		Intro:       "A RESTful Golang API for persisting application metadata",
	})

	if err := ioutil.WriteFile("routes.md", []byte(md), 0644); err != nil {
		log.Println(err)
		return
	}
	fmt.Println("OK")
}
