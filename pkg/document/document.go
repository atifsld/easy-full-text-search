package document

import (
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"os"
)

// Document represents a single document in the search index
type Document struct {
	Text string `xml:"abstract"`
	ID   int    `xml:"-"`
}

// LoadDocuments loads documents from a gzipped XML file
func LoadDocuments(path string) ([]Document, error) {
	fmt.Printf("Loading documents from: %s\n", path)

	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gz.Close()

	dec := xml.NewDecoder(gz)
	var feed struct {
		Docs []Document `xml:"doc"`
	}

	if err := dec.Decode(&feed); err != nil {
		return nil, fmt.Errorf("failed to decode XML: %w", err)
	}

	fmt.Printf("Found %d documents in XML\n", len(feed.Docs))

	docs := feed.Docs
	for i := range docs {
		docs[i].ID = i
	}
	return docs, nil
}
