package gopenai

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const editsAPIEndpoint = "/edits"

// EditChoice represents a possible correction for a piece of
// text provided by OpenAI's GPT-3 API.
type EditChoice struct {
	// Text is the corrected text.
	Text string `json:"text"`
	// Index is the index of the corrected text.
	Index int `json:"index"`
}

// Edit represents a response from OpenAI's GPT-3 API
// for text correction requests.
type Edit struct {
	// Created is a timestamp of when the response was created.
	Created int `json:"created"`
	// Choices is a list of possible corrections for the input text.
	Choices []EditChoice `json:"choices"`
	// Usage contains information on the usage of OpenAI's API tokens.
	Usage TokenUsage `json:"usage"`
}

// EditParams represents the parameters for a text correction
// request to OpenAI's GPT-3 API.
type EditParams struct {
	// Model is the name of the language model to use.
	Model string `json:"model"`
	// Input is the text to be corrected.
	Input string `json:"input,omitempty"`
	// Instruction is a description of the type of correction
	// to perform on the input text.
	Instruction string `json:"instruction"`
	// N is the number of corrections to return.
	N int `json:"n,omitempty"`
	// Temperature is a value used to control the randomness of the response.
	Temperature float64 `json:"temperature,omitempty"`
	// TopP is a value used to control the diversity of the response.
	TopP float64 `json:"top_p,omitempty"`
}

// EditsAPI is an interface that provides methods for text
// correction with OpenAI's GPT-3 API.
type EditsAPI interface {
	// Create makes a text correction request to OpenAI's GPT-3 API
	// using the provided parameters.
	Create(EditParams) (Edit, error)
}

type editsAPI struct {
	c client
}

func (api editsAPI) Create(params EditParams) (Edit, error) {
	url := fmt.Sprintf("%s%s", baseURL, editsAPIEndpoint)
	r, err := api.c.getJSONRequestResponse(url, http.MethodPost, params)
	if err != nil {
		return Edit{}, err
	}

	var response Edit
	if err := json.Unmarshal(r, &response); err != nil {
		return Edit{}, err
	}

	return response, nil
}
