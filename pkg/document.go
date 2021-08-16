package pkg

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Document struct {
	Text string
	ID   int
}

func loadDocuments(path string) ([]Document, error) {
	var docs []Document
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			b, err := ioutil.ReadFile(path) // just pass the file name
			if err != nil {
				fmt.Print(err)
			}
			str := string(b) // convert content to a 'string'
			docs = append(docs, Document{Text: str})

			return nil
		})
	if err != nil {
		log.Println(err)
	}

	for i := range docs {
		docs[i].ID = i
	}
	return docs, nil
}
