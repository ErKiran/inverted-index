package pkg

import (
	"log"
	"os"
	"time"
)

func Indexer() ([]Document, Index, error) {
	start := time.Now()
	path := os.Getenv("DOCUMENT_PATH")
	docs, err := loadDocuments(path)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	idx := make(Index)
	idx.Add(docs)

	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

	return docs, idx, nil
}
