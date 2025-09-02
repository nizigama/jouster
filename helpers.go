package main

import (
	"log"
	"sort"
	"strings"

	"github.com/invopop/jsonschema"
	"github.com/jdkato/prose/v2"
)

func generateJsonSchema(t llmStructure) interface{} {
	// Structured Outputs uses a subset of JSON schema
	// These flags are necessary to comply with the subset
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}

	schema := reflector.Reflect(t)
	return schema
}

func extractThreeNouns(text string) []string {
	// Create a document with prose
	doc, err := prose.NewDocument(text)
	if err != nil {
		log.Println(err)
		return []string{}
	}

	// Count noun frequencies
	nouns := []string{}
	for _, tok := range doc.Tokens() {

		switch tok.Tag {
		case "NN", "NNS", "NNP", "NNPS":
			nouns = append(nouns, strings.ToLower(tok.Text))
		}
	}

	// Count frequencies
	frequencies := make(map[string]int)
	for _, noun := range nouns {
		frequencies[noun]++
	}

	var items []NameFrequency
	for name, freq := range frequencies {
		items = append(items, NameFrequency{name, freq})
	}

	// Sort by frequency (descending)
	sort.Slice(items, func(i, j int) bool {
		return items[i].Frequency > items[j].Frequency
	})

	// Extract top 3 names
	var top3 []string
	for i := 0; i < len(items) && i < 3; i++ {
		top3 = append(top3, items[i].Name)
	}

	return top3
}
