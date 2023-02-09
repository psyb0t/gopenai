package gopenai

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const moderationsAPIEndpoint = "/moderations"

// Moderation represents the response of moderation API.
// It contains the moderation ID, the used model and the moderation results.
type Moderation struct {
	// ID is the identifier of the moderation.
	ID string `json:"id"`
	// Model is the name of the model that was used for the moderation.
	Model string `json:"model"`
	// Results contains the moderation result for each individual prompt.
	Results []ModerationResult `json:"results"`
}

// ModerationResult represents the moderation result for an individual prompt.
// It contains the moderation categories, category scores
// and whether the prompt is flagged or not.
type ModerationResult struct {
	// Categories represents the different moderation categories.
	Categories ModerationResultCategories `json:"categories"`
	// CategoryScores represents the scores of the different
	// moderation categories.
	CategoryScores ModerationResultCategoryScores `json:"category_scores"`
	// Flagged indicates whether the prompt is flagged by the model or not.
	Flagged bool `json:"flagged"`
}

// ModerationResultCategories represents the different
// moderation categories for a prompt.
type ModerationResultCategories struct {
	// Hate indicates whether the prompt contains hate content.
	Hate bool `json:"hate"`
	// HateThreatening indicates whether the prompt contains
	// hate threatening content.
	HateThreatening bool `json:"hate/threatening"`
	// SelfHarm indicates whether the prompt contains self-harm content.
	SelfHarm bool `json:"self-harm"`
	// Sexual indicates whether the prompt contains sexual content.
	Sexual bool `json:"sexual"`
	// SexualMinors indicates whether the prompt contains
	// sexual content with minors.
	SexualMinors bool `json:"sexual/minors"`
	// Violence indicates whether the prompt contains violent content.
	Violence bool `json:"violence"`
	// ViolenceGraphic indicates whether the prompt contains
	// graphic violent content.
	ViolenceGraphic bool `json:"violence/graphic"`
}

// ModerationResultCategoryScores represents the scores of the
// different moderation categories for a prompt.
type ModerationResultCategoryScores struct {
	// Hate is the score for the hate category.
	Hate float64 `json:"hate"`
	// HateThreatening is the score for the hate threatening category.
	HateThreatening float64 `json:"hate/threatening"`
	// SelfHarm is the score for the self-harm category.
	SelfHarm float64 `json:"self-harm"`
	// Sexual is the score for the sexual category.
	Sexual float64 `json:"sexual"`
	// SexualMinors is the score for the sexual with minors category.
	SexualMinors float64 `json:"sexual/minors"`
	// Violence is the score for the violent category.
	Violence float64 `json:"violence"`
	// ViolenceGraphic is the score for the graphic violent category.
	ViolenceGraphic float64 `json:"violence/graphic"`
}

// ModerationParams represents the parameters to use for moderation
type ModerationParams struct {
	// Input is the text input to be moderated
	Input string `json:"input"`
	// Model is the name of the model to be used for moderation
	Model string `json:"model,omitempty"`
}

// ModerationsAPI is the interface for the OpenAI moderations API.
type ModerationsAPI interface {
	Create(ModerationParams) (Moderation, error)
}

type moderationsAPI struct {
	c client
}

func (api moderationsAPI) Create(params ModerationParams) (Moderation, error) {
	url := fmt.Sprintf("%s%s", baseURL, moderationsAPIEndpoint)
	r, err := api.c.getJSONRequestResponse(url, http.MethodPost, params)
	if err != nil {
		return Moderation{}, err
	}

	var response Moderation
	if err := json.Unmarshal(r, &response); err != nil {
		return Moderation{}, err
	}

	return response, nil
}
