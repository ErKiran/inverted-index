In `.env` file pass the relative file location of the Email cropus


### RUN AS CLI 

* First Load and Index the cropus  
`go run main.go indexer`  
This will load the documents from the file and Index the document.

* Search the Indexed Document 
  
  `go run main.go retriever.go Willamson`

  This will traverse through the index document and loads the all the matching documents in the terminal
