package gopenai

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const modelsAPIEndpoint = "/models"

// ModelPermission represents a permission a user has on an OpenAI model.
type ModelPermission struct {
	// ID is the unique identifier for the model permission.
	ID string `json:"id"`
	// Created is the timestamp indicating when the permission was created.
	Created uint `json:"created"`
	// AllowCreateEngine represents the permission to
	// create engines for the model.
	AllowCreateEngine bool `json:"allow_create_engine"`
	// AllowSampling represents the permission to sample the model.
	AllowSampling bool `json:"allow_sampling"`
	// AllowLogprobs represents the permission to retrieve
	// log probabilities from the model.
	AllowLogprobs bool `json:"allow_logprobs"`
	// AllowSearchIndices represents the permission to
	// search the indices of the model.
	AllowSearchIndices bool `json:"allow_search_indices"`
	// AllowView represents the permission to view the model's metadata.
	AllowView bool `json:"allow_view"`
	// AllowFineTuning represents the permission to fine-tune the model.
	AllowFineTuning bool `json:"allow_fine_tuning"`
	// Organization is the name of the organization that owns the model.
	Organization string `json:"organization"`
	// Group is the group that the model belongs to.
	Group interface{} `json:"group"`
	// IsBlocking indicates whether this permission blocks other permissions.
	IsBlocking bool `json:"is_blocking"`
}

// Model represents information about an OpenAI model.
type Model struct {
	// ID is the unique identifier for the model.
	ID string `json:"id"`
	// OwnedBy is the name of the user or organization that owns the model.
	OwnedBy string `json:"owned_by"`
	// Created is the timestamp indicating when the model was created.
	Created uint `json:"created"`
	// Root is the root model in the lineage of the model.
	Root string `json:"root"`
	// Parent is the parent model in the lineage of the model.
	Parent interface{} `json:"parent"`
	// Permission is a slice of ModelPermission objects
	// representing the permissions on the model.
	Permission []ModelPermission `json:"permission"`
}

// DeletedModel represents a model that has been deleted from the OpenAI API
type DeletedModel struct {
	// ID is the identifier of the model
	ID string `json:"id"`
	// Deleted is whether or not the model was deleted successfully
	Deleted bool `json:"deleted"`
}

// ModelsAPI is the interface for the OpenAI models API.
type ModelsAPI interface {
	// GetAll returns all available models.
	GetAll() ([]Model, error)
	// GetByID returns the model with the specified ID.
	GetByID(id string) (Model, error)
	// DeleteByID deletes a model by its identifier
	DeleteByID(id string) (DeletedModel, error)
}

type modelsAPI struct {
	c client
}

func (api modelsAPI) GetAll() ([]Model, error) {
	url := fmt.Sprintf("%s%s", baseURL, modelsAPIEndpoint)
	r, err := api.c.getJSONRequestResponse(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	var response struct {
		Data []Model `json:"data"`
	}

	if err := json.Unmarshal(r, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api modelsAPI) GetByID(id string) (Model, error) {
	url := fmt.Sprintf("%s%s/%s", baseURL, modelsAPIEndpoint, id)
	r, err := api.c.getJSONRequestResponse(url, http.MethodGet, nil)
	if err != nil {
		return Model{}, err
	}

	var response Model
	if err := json.Unmarshal(r, &response); err != nil {
		return Model{}, err
	}

	return response, nil
}

func (api modelsAPI) DeleteByID(id string) (DeletedModel, error) {
	url := fmt.Sprintf("%s%s/%s", baseURL, modelsAPIEndpoint, id)
	r, err := api.c.getJSONRequestResponse(url, http.MethodDelete, nil)
	if err != nil {
		return DeletedModel{}, err
	}

	var response DeletedModel
	if err := json.Unmarshal(r, &response); err != nil {
		return DeletedModel{}, err
	}

	return response, nil
}
