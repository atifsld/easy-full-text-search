package index

import (
	"fmt"

	"github.com/atifsld/easy-full-text-search/pkg/analyzer"
	"github.com/atifsld/easy-full-text-search/pkg/document"
)

// Index represents a full-text search index
type Index struct {
	documents map[int]document.Document
	tokens    map[string][]int
}

// New creates a new empty search index
func New() *Index {
	return &Index{
		documents: make(map[int]document.Document),
		tokens:    make(map[string][]int),
	}
}

// AddDocument adds a single document to the index
func (idx *Index) AddDocument(doc document.Document) {
	idx.documents[doc.ID] = doc
	for _, token := range analyzer.Analyze(doc.Text) {
		ids := idx.tokens[token]
		if ids != nil && ids[len(ids)-1] == doc.ID {
			// Don't add same ID twice
			continue
		}
		idx.tokens[token] = append(ids, doc.ID)
	}
}

// AddDocuments adds multiple documents to the index
func (idx *Index) AddDocuments(docs []document.Document) {
	for _, doc := range docs {
		idx.AddDocument(doc)
	}
}

// LoadFromFile loads documents from a gzipped XML file and adds them to the index
func (idx *Index) LoadFromFile(path string) error {
	docs, err := document.LoadDocuments(path)
	if err != nil {
		return fmt.Errorf("failed to load documents: %w", err)
	}
	idx.AddDocuments(docs)
	return nil
}

// Intersection finds the common document IDs between two sorted slices
func Intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	r := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

// Search finds documents matching the query
func (idx *Index) Search(text string) []document.Document {
	var docIDs []int
	for _, token := range analyzer.Analyze(text) {
		if ids, ok := idx.tokens[token]; ok {
			if docIDs == nil {
				docIDs = ids
			} else {
				docIDs = Intersection(docIDs, ids)
			}
		} else {
			// Token doesn't exist
			return nil
		}
	}

	// Convert IDs to documents
	results := make([]document.Document, len(docIDs))
	for i, id := range docIDs {
		results[i] = idx.documents[id]
	}
	return results
}
