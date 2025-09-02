package main

type LlmRequest struct {
	Text string `json:"text"`
}
type LlmResponse struct {
	Summary  string            `json:"summary"`
	Metadata metadataStructure `json:"metadata"`
}

type llmStructure struct {
	Title     string   `json:"title" jsonschema_description:"The title of the text if available or applicable"`
	Topics    []string `json:"topics" jsonschema_description:"A list of topics covered in the text"`
	Sentiment string   `json:"sentiment" jsonschema:"enum=positive,enum=neutral,enum=negative" jsonschema_description:"The overall sentiment of the text"`
}

type metadataStructure struct {
	Title     string   `json:"title"`
	Topics    []string `json:"topics"`
	Sentiment string   `json:"sentiment"`
	Keywords  []string `json:"keywords"`
}

type NameFrequency struct {
	Name      string
	Frequency int
}
