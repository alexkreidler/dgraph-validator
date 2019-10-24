/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"os"

	"github.com/spf13/cobra"

	"github.com/dgraph-io/dgraph/schema"
	"github.com/dgraph-io/dgraph/x"

	"compress/gzip"
	"io"
	"io/ioutil"
	"path/filepath"
)

func readSchema(filename string) *schema.ParsedSchema {
	f, err := os.Open(filename)
	x.Check(err)
	defer f.Close()
	var r io.Reader = f
	if filepath.Ext(filename) == ".gz" {
		r, err = gzip.NewReader(f)
		x.Check(err)
	}

	buf, err := ioutil.ReadAll(r)
	x.Check(err)

	result, err := schema.Parse(string(buf))
	x.Check(err)
	return result
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dgraph-validator",
	Short: "Validate dgraph schema files",
	Long:  `Links to the dgraph schema package to validate schema files`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		readSchema(path)
	},
	Args: cobra.MinimumNArgs(1),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dgraph-validator.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
