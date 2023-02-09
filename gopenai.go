// Package gopenai provides Go bindings for the OpenAI API.
package gopenai

import (
	"net/http"
	"time"
)

const (
	baseURL               = "https://api.openai.com/v1"
	defaultRequestTimeout = time.Second * 30
)

// Config holds the configuration values for the OpenAI client.
type Config struct {
	APIKey         string
	OrganizationID string
	RequestTimeout time.Duration
}

// Client is the interface for interacting with the OpenAI API.
type Client interface {
	// Models returns the ModelsAPI for interacting with the models.
	Models() ModelsAPI
	// Completions returns the CompletionsAPI for interacting with completions.
	Completions() CompletionsAPI
	// Edits returns the EditsAPI for interacting with edits.
	Edits() EditsAPI
	// Images returns the ImagesAPI for interacting with images.
	Images() ImagesAPI
	// Embeddings returns the EmbeddingsAPI for interacting with embeddings.
	Embeddings() EmbeddingsAPI
	// Files returns the FilesAPI for interacting with files.
	Files() FilesAPI
	// FineTunes returns the FineTunesAPI for interacting with fine-tunes.
	FineTunes() FineTunesAPI
	// Moderations returns the ModerationsAPI for interacting with moderations.
	Moderations() ModerationsAPI
}

// New returns a new OpenAI client with the given configuration.
// If RequestTimeout is not set in the config, it will be set to defaultRequestTimeout.
func New(cfg Config) Client {
	if cfg.RequestTimeout == 0 {
		cfg.RequestTimeout = defaultRequestTimeout
	}

	return client{
		cfg: cfg,
		httpClient: &http.Client{
			Timeout: cfg.RequestTimeout,
		},
	}
}
