package cmd

import (
	"encoding/json"
	"fmt"
	"fts/pkg"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var documentFileName = "./document.json"
var indexFileName = "./index.json"

var docs []pkg.Document
var idx pkg.Index

var indexerCmd = &cobra.Command{
	Use:   "indexer",
	Short: "Parse the Document and Build the Inverted Index",
	Long:  `Parse the Document and Build the Inverted Index`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(documentFileName); os.IsNotExist(err) {
			docs, idx, err := pkg.Indexer()
			if err != nil {
				panic(err)
			}

			document, _ := json.MarshalIndent(docs, "", " ")
			index, _ := json.MarshalIndent(idx, "", " ")

			_ = ioutil.WriteFile("./document.json", document, 0644)
			_ = ioutil.WriteFile("./index.json", index, 0644)
			return
		}
		fmt.Println("Sucessfully Indexed Documents....")
	},
}

var searchCmd = &cobra.Command{
	Use:   "retriver",
	Short: "Search the Build Inverted Index",
	Long:  `Parse the Document and Build the Inverted Index`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(indexFileName); os.IsNotExist(err) {
			fmt.Println("Index file does not exist....")
			fmt.Println("Please run with indexer flag")
			return
		}

		if _, err := os.Stat(documentFileName); os.IsNotExist(err) {
			fmt.Println("Document file does not exist....")
			fmt.Println("Please run with indexer flag")
			return
		}

		documentFile, _ := ioutil.ReadFile(documentFileName)
		indexFile, _ := ioutil.ReadFile(indexFileName)

		_ = json.Unmarshal([]byte(documentFile), &docs)
		_ = json.Unmarshal([]byte(indexFile), &idx)
		pkg.Query(args[0], docs, &idx)
	},
}

func init() {
	rootCmd.AddCommand(indexerCmd)
	rootCmd.AddCommand(searchCmd)
}
