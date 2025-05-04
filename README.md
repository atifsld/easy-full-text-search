# Easy Full Text Search

A simple full-text search library for Go for prototyping and small applications. This project was created as my first foray into Go programming. It's a learning exercise where I wanted to explore the language while building something practical. The implementation is intentionally kept simple and straightforward. One of the main motivations for this project was to understand Go modules and learn how to publish Go packages.

## Acknowledgments

This project is based on [go-full-text-search](https://github.com/AkhilSharma90/go-full-text-search) by Akhil Sharma. I've made some modifications to make it more accessible and added some improvements while learning Go. All credit for the original implementation goes to the original author.

## Features

- Simple API for creating and querying search indices
- Support for loading documents from gzipped XML files
- Built-in text analysis pipeline including:
  - Tokenization
  - Lowercase conversion
  - Stopword removal
  - Stemming

## Installation

```bash
go get github.com/atifsld/easy-full-text-search
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/atifsld/easy-full-text-search/pkg/index"
)

func main() {
    // Create a new index
    idx := index.New()
    
    // Load documents from a file
    err := idx.LoadFromFile("path/to/your/documents.xml.gz")
    if err != nil {
        log.Fatal(err)
    }
    
    // Search for documents
    results := idx.Search("your search query")
    
    // Print results
    for _, doc := range results {
        fmt.Printf("Document %d: %s\n", doc.ID, doc.Text)
    }
}
```

## Document Format

The library expects documents in a gzipped XML format with the following structure:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<feed>
    <doc>
        <abstract>Document content goes here...</abstract>
    </doc>
    <!-- More documents... -->
</feed>
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details. 