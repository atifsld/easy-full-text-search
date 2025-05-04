package index

import (
	"testing"

	"github.com/atifsld/easy-full-text-search/pkg/document"
)

func TestIndex(t *testing.T) {
	// Create a new index
	idx := New()

	// Create test documents
	docs := []document.Document{
		{
			ID:   1,
			Text: "This is a document about programming languages",
		},
		{
			ID:   2,
			Text: "Another document about software development",
		},
	}

	// Add documents to the index
	idx.AddDocuments(docs)

	// Test cases
	tests := []struct {
		name     string
		query    string
		expected int
	}{
		{
			name:     "Search for 'programming'",
			query:    "programming",
			expected: 1,
		},
		{
			name:     "Search for 'software'",
			query:    "software",
			expected: 1,
		},
		{
			name:     "Search for 'document'",
			query:    "document",
			expected: 2,
		},
		{
			name:     "Search for non-existent term",
			query:    "nonexistent",
			expected: 0,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := idx.Search(tt.query)
			if len(results) != tt.expected {
				t.Errorf("Search(%q) returned %d results, want %d", tt.query, len(results), tt.expected)
			}
		})
	}
}
