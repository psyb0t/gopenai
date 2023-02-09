package gopenai

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const filesAPIEndpoint = "/files"

// File represents a file in the OpenAI API
type File struct {
	// ID is the identifier of the file
	ID string `json:"id"`
	// Bytes is the number of bytes of the file
	Bytes int `json:"bytes"`
	// CreatedAt is a timestamp of when the file was created
	CreatedAt int `json:"created_at"`
	// Filename is the name of the file
	Filename string `json:"filename"`
	// Purpose is the purpose of the file
	Purpose string `json:"purpose"`
}

// FileParams contains the parameters to create a new file in the OpenAI API
type FileParams struct {
	// File is the file content
	File string `mapstructure:"file"`
	// Purpose is the purpose of the file
	Purpose string `mapstructure:"purpose"`
}

// DeletedFile represents a file that has been deleted from the OpenAI API
type DeletedFile struct {
	// ID is the identifier of the file
	ID string `json:"id"`
	// Deleted is whether or not the file was deleted successfully
	Deleted bool `json:"deleted"`
}

// FilesAPI represents the OpenAI API for managing files
type FilesAPI interface {
	// GetAll returns all the files in the OpenAI API
	GetAll() ([]File, error)
	// GetByID returns a file by its identifier
	GetByID(id string) (File, error)
	// Create creates a new file in the OpenAI API
	Create(FileParams) (File, error)
	// DeleteByID deletes a file by its identifier
	DeleteByID(id string) (DeletedFile, error)
	// DownloadByID downloads a file by its identifier
	DownloadByID(id string, dst io.Writer) error
}

type filesAPI struct {
	c client
}

func (api filesAPI) GetAll() ([]File, error) {
	url := fmt.Sprintf("%s%s", baseURL, filesAPIEndpoint)
	r, err := api.c.getJSONRequestResponse(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	var response struct {
		Data []File `json:"data"`
	}

	if err := json.Unmarshal(r, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api filesAPI) GetByID(id string) (File, error) {
	url := fmt.Sprintf("%s%s/%s", baseURL, filesAPIEndpoint, id)
	r, err := api.c.getJSONRequestResponse(url, http.MethodGet, nil)
	if err != nil {
		return File{}, err
	}

	var response File
	if err := json.Unmarshal(r, &response); err != nil {
		return File{}, err
	}

	return response, nil
}

func (api filesAPI) DeleteByID(id string) (DeletedFile, error) {
	url := fmt.Sprintf("%s%s/%s", baseURL, filesAPIEndpoint, id)
	r, err := api.c.getJSONRequestResponse(url, http.MethodDelete, nil)
	if err != nil {
		return DeletedFile{}, err
	}

	var response DeletedFile
	if err := json.Unmarshal(r, &response); err != nil {
		return DeletedFile{}, err
	}

	return response, nil
}

func (api filesAPI) Create(params FileParams) (File, error) {
	url := fmt.Sprintf("%s%s", baseURL, filesAPIEndpoint)
	data, contentType, err := structToMultipartFormData(params)
	if err != nil {
		return File{}, err
	}

	r, err := api.c.getRequestResponse(url, http.MethodPost, data, contentType)
	if err != nil {
		return File{}, err
	}

	var response File
	if err := json.Unmarshal(r, &response); err != nil {
		return File{}, err
	}

	return response, nil
}

func (api filesAPI) DownloadByID(id string, dst io.Writer) error {
	url := fmt.Sprintf("%s%s/%s/content", baseURL, filesAPIEndpoint, id)
	err := api.c.streamRequestResponse(url, http.MethodGet, nil, "", dst)
	if err != nil {
		return err
	}

	return nil
}
