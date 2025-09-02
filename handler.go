package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/gofiber/fiber/v3"
	"github.com/openai/openai-go/v2"
)

func analyze(c fiber.Ctx) error {

	req := new(LlmRequest)

	if err := c.Bind().Body(req); err != nil {
		return err
	}

	client := openai.NewClient()

	var summary string
	var mdataStructure metadataStructure
	var errors []error

	var wg sync.WaitGroup

	wg.Add(2)

	go func(smr *string, errs []error) {
		defer wg.Done()

		summary, err := generateSummary(req.Text, client)
		if err != nil {
			errs = append(errs, err)
			return
		}

		*smr = summary
	}(&summary, errors)

	go func(ms *metadataStructure, errs []error) {
		defer wg.Done()

		mdata, err := generateMetadata(req.Text, client)
		if err != nil {
			errs = append(errs, err)
			return
		}

		*ms = mdata
	}(&mdataStructure, errors)

	wg.Wait()

	if len(errors) > 0 {
		log.Println(errors)
		return fmt.Errorf("Failed to process request")
	}

	response := LlmResponse{
		Summary:  summary,
		Metadata: mdataStructure,
	}

	v, err := json.Marshal(response)
	if err != nil {
		return err
	}

	return c.Send(v)
}

func generateSummary(text string, client openai.Client) (string, error) {
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(fmt.Sprintf("Analyze the following text and generate a summary of 1-2 sentences: %s", text)),
		},
		Model: openai.ChatModelGPT4oMini,
	})

	if err != nil {
		return "", err
	}

	return chatCompletion.Choices[0].Message.Content, nil
}

func generateMetadata(text string, client openai.Client) (metadataStructure, error) {

	var ms metadataStructure

	schema := generateJsonSchema(llmStructure{})

	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        "metadata",
		Description: openai.String("Metadata extracted from the text"),
		Schema:      schema,
		Strict:      openai.Bool(true),
	}

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(fmt.Sprintf("Analyze the following text and extract structured metadata as JSON: %s", text)),
		},
		ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
			OfJSONSchema: &openai.ResponseFormatJSONSchemaParam{
				JSONSchema: schemaParam,
			},
		},
		Model: openai.ChatModelGPT4oMini,
	})
	if err != nil {
		return metadataStructure{}, err
	}

	err = json.Unmarshal([]byte(chatCompletion.Choices[0].Message.Content), &ms)
	if err != nil {
		return metadataStructure{}, err
	}

	words := extractThreeNouns(text)

	ms.Keywords = words

	return ms, nil
}
