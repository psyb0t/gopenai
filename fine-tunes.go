package gopenai

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const fineTunesAPIEndpoint = "/fine-tunes"

// FineTune represents the information of a fine-tuning task.
type FineTune struct {
	// ID is the unique identifier for the fine-tuning task.
	ID string `json:"id"`
	// Model is the name of the OpenAI GPT-3 model that was fine-tuned.
	Model string `json:"model"`
	// CreatedAt is the timestamp when the fine-tuning task was created.
	CreatedAt int `json:"created_at"`
	// Events are the events that have taken place during the fine-tuning task.
	Events []FineTuneEvent `json:"events"`
	// FineTunedModel is the name of the fine-tuned model that was
	// created as a result of the fine-tuning task.
	FineTunedModel string `json:"fine_tuned_model"`
	// Hyperparams are the hyperparameters used during the fine-tuning task.
	Hyperparams FineTuneHyperparams `json:"hyperparams"`
	// OrganizationID is the ID of the organization that
	// created the fine-tuning task.
	OrganizationID string `json:"organization_id"`
	// ResultFiles are the files associated with the
	// results of the fine-tuning task.
	ResultFiles []File `json:"result_files"`
	// Status is the current status of the fine-tuning task.
	Status string `json:"status"`
	// ValidationFiles are the files used to validate the fine-tuning task.
	ValidationFiles []interface{} `json:"validation_files"`
	// TrainingFiles are the files used to train the fine-tuning task.
	TrainingFiles []File `json:"training_files"`
	// UpdatedAt is the timestamp when the fine-tuning task was last updated.
	UpdatedAt int `json:"updated_at"`
}

// FineTuneEvent represents a single event that has taken
// place during a fine-tuning task.
type FineTuneEvent struct {
	// CreatedAt is the timestamp when the event took place.
	CreatedAt int `json:"created_at"`
	// Level is the severity level of the event.
	Level string `json:"level"`
	// Message is the description of the event.
	Message string `json:"message"`
}

// FineTuneHyperparams represents the hyperparameters
// used during a fine-tuning task.
type FineTuneHyperparams struct {
	// BatchSize is the size of the batch used during the fine-tuning task.
	BatchSize int `json:"batch_size"`
	// LearningRateMultiplier is the learning rate multiplier used
	// during the fine-tuning task.
	LearningRateMultiplier float64 `json:"learning_rate_multiplier"`
	// NEpochs is the number of epochs used during the fine-tuning task.
	NEpochs int `json:"n_epochs"`
	// PromptLossWeight is the weight assigned to the
	// prompt loss during the fine-tuning task.
	PromptLossWeight float64 `json:"prompt_loss_weight"`
}

// FineTuneParams defines the parameters required to create a fine-tune task.
type FineTuneParams struct {
	// TrainingFile is the path to the training file.
	TrainingFile string `json:"training_file"`
	// ValidationFile is the path to the validation file.
	ValidationFile string `json:"validation_file,omitempty"`
	// Model is the identifier of the base model to be fine-tuned.
	Model string `json:"model,omitempty"`
	// NEpochs is the number of training epochs to be performed.
	NEpochs int `json:"n_epochs,omitempty"`
	// BatchSize is the number of samples per training iteration.
	BatchSize int `json:"batch_size,omitempty"`
	// LearningRateMultiplier is the multiplier applied to the learning rate.
	LearningRateMultiplier float64 `json:"learning_rate_multiplier,omitempty"`
	// PromptLossWeight is the weight applied to the prompt loss.
	PromptLossWeight float64 `json:"prompt_loss_weight,omitempty"`
	// ComputeClassificationMetrics indicates whether to compute
	// classification metrics.
	ComputeClassificationMetrics bool `json:"compute_classification_metrics,omitempty"`
	// ClassificationNClasses is the number of classes in the
	// classification task.
	ClassificationNClasses int `json:"classification_n_classes,omitempty"`
	// ClassificationPositiveClass is the positive class label in
	// the classification task.
	ClassificationPositiveClass string `json:"classification_positive_class,omitempty"`
	// ClassificationBetas are the betas for Fbeta metrics computation.
	ClassificationBetas []interface{} `json:"classification_betas,omitempty"`
	// Suffix is an optional string to be appended to the
	// fine-tuned model identifier.
	Suffix string `json:"suffix,omitempty"`
}

// FineTunesAPI represents the interface for managing
// fine-tuning models in OpenAI GPT.
type FineTunesAPI interface {
	// GetAll retrieves all the fine-tuning models available.
	GetAll() ([]FineTune, error)
	// GetByID retrieves a specific fine-tuning model based on its ID.
	GetByID(id string) (FineTune, error)
	// Create creates a new fine-tuning model.
	Create(FineTuneParams) (FineTune, error)
	// Cancel cancels a specific fine-tuning model based on its ID.
	Cancel(id string) (FineTune, error)
	// GetEvents retrieves all the events of a specific
	// fine-tuning model based on its ID.
	GetEvents(fineTuneID string) ([]FineTuneEvent, error)
}

type fineTunesAPI struct {
	c client
}

func (api fineTunesAPI) GetAll() ([]FineTune, error) {
	url := fmt.Sprintf("%s%s", baseURL, fineTunesAPIEndpoint)
	r, err := api.c.getJSONRequestResponse(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	var response struct {
		Data []FineTune `json:"data"`
	}

	if err := json.Unmarshal(r, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api fineTunesAPI) GetByID(id string) (FineTune, error) {
	url := fmt.Sprintf("%s%s/%s", baseURL, fineTunesAPIEndpoint, id)
	r, err := api.c.getJSONRequestResponse(url, http.MethodGet, nil)
	if err != nil {
		return FineTune{}, err
	}

	var response FineTune
	if err := json.Unmarshal(r, &response); err != nil {
		return FineTune{}, err
	}

	return response, nil
}

func (api fineTunesAPI) Create(params FineTuneParams) (FineTune, error) {
	url := fmt.Sprintf("%s%s", baseURL, fineTunesAPIEndpoint)
	r, err := api.c.getJSONRequestResponse(url, http.MethodPost, params)
	if err != nil {
		return FineTune{}, err
	}

	var response FineTune
	if err := json.Unmarshal(r, &response); err != nil {
		return FineTune{}, err
	}

	return response, nil
}

func (api fineTunesAPI) Cancel(id string) (FineTune, error) {
	url := fmt.Sprintf("%s%s/%s/cancel", baseURL, fineTunesAPIEndpoint, id)
	r, err := api.c.getJSONRequestResponse(url, http.MethodPost, nil)
	if err != nil {
		return FineTune{}, err
	}

	var response FineTune
	if err := json.Unmarshal(r, &response); err != nil {
		return FineTune{}, err
	}

	return response, nil
}

// TODO: support stream
func (api fineTunesAPI) GetEvents(fineTuneID string) ([]FineTuneEvent, error) {
	url := fmt.Sprintf("%s%s/%s/events", baseURL, fineTunesAPIEndpoint, fineTuneID)
	r, err := api.c.getJSONRequestResponse(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	var response struct {
		Data []FineTuneEvent `json:"data"`
	}

	if err := json.Unmarshal(r, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}
