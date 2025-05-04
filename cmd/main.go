package main

import (
	"fmt"

	"github.com/atifsld/easy-full-text-search/pkg/document"
	"github.com/atifsld/easy-full-text-search/pkg/index"
)

func main() {
	// Create a new index
	idx := index.New()

	// Load documents from the XML file
	docs, err := document.LoadDocuments("test.xml.gz")
	if err != nil {
		fmt.Printf("Error loading documents: %v\n", err)
		return
	}

	fmt.Printf("Loaded %d documents:\n", len(docs))
	for _, doc := range docs {
		fmt.Printf("Document %d: %s\n", doc.ID, doc.Text)
	}

	// Add documents to the index
	idx.AddDocuments(docs)

	// Test some searches
	queries := []string{
		"programmer",
		"golang",
		"javascript",
		"bugs",
		"promise",
	}

	fmt.Println("\nSearch Results:")
	fmt.Println("==============")

	for _, query := range queries {
		fmt.Printf("\nSearching for: %q\n", query)
		results := idx.Search(query)

		if len(results) == 0 {
			fmt.Println("No results found")
			continue
		}

		fmt.Printf("Found %d results:\n", len(results))
		for _, doc := range results {
			fmt.Printf("- Document %d: %s\n", doc.ID, doc.Text)
		}
	}
}
