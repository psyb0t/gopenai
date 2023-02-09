package gopenai

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	imageGenerationsAPIEndpoint = "/images/generations"
	imageEditsAPIEndpoint       = "/images/edits"
	imageVariationsAPIEndpoint  = "/images/variations"
)

// ImageResponseFormat is a type that defines the format of image response.
type ImageResponseFormat string

// Constants for different types of ImageResponseFormat
const (
	ImageResponseFormatURL     ImageResponseFormat = "url"
	ImageResponseFormatB64JSON ImageResponseFormat = "b64_json"
)

// ImageSize is a type that defines the size of the image.
type ImageSize string

// Constants for different types of ImageSize
const (
	ImageSize256x256   ImageSize = "256x256"
	ImageSize512x512   ImageSize = "512x512"
	ImageSize1024x1024 ImageSize = "1024x1024"
)

// Image represents a single image in response from the OpenAI API.
type Image struct {
	// URL is the URL of the image.
	URL string `json:"url,omitempty"`
	// B64JSON is the Base64 encoded string of the image.
	B64JSON string `json:"b64_json,omitempty"`
}

// ImageGenerationParams are the parameters for generating new images.
type ImageGenerationParams struct {
	// Prompt is the prompt text used to generate the images.
	Prompt string `json:"prompt"`
	// N is the number of images to generate.
	N uint `json:"n"`
	// Size is the size of the images.
	Size ImageSize `json:"size"`
	// ResponseFormat is the format of the image response.
	ResponseFormat ImageResponseFormat `json:"response_format"`
	// User is the name of the user.
	User string `json:"user"`
}

// ImageEditParams are the parameters for editing existing images.
type ImageEditParams struct {
	// Image is the image to edit.
	Image string `mapstructure:"image"`
	// Mask is the image to use as a mask for the edit.
	Mask string `mapstructure:"mask,omitempty"`
	// Prompt is the prompt text used to edit the images.
	Prompt string `mapstructure:"prompt"`
	// N is the number of images to generate.
	N uint `mapstructure:"n,omitempty"`
	// Size is the size of the images.
	Size ImageSize `mapstructure:"size,omitempty"`
	// ResponseFormat is the format of the image response.
	ResponseFormat ImageResponseFormat `mapstructure:"response_format,omitempty"`
	// User is the name of the user.
	User string `mapstructure:"user,omitempty"`
}

// ImageVariationParams are the parameters for generating
// variations of existing images.
type ImageVariationParams struct {
	// Image is the image to generate variations from.
	Image string `mapstructure:"image"`
	// N is the number of images to generate.
	N uint `mapstructure:"n,omitempty"`
	// Size is the size of the images.
	Size ImageSize `mapstructure:"size,omitempty"`
	// ResponseFormat is the format of the image response.
	ResponseFormat ImageResponseFormat `mapstructure:"response_format,omitempty"`
	// User is an optional field that allows to specify a
	// user identifier for the API request.
	User string `mapstructure:"user,omitempty"`
}

// ImagesAPI is an interface that defines the methods for generating
// and manipulating images with OpenAI's API.
type ImagesAPI interface {
	// Create generates new images based on the given prompt.
	Create(ImageGenerationParams) ([]Image, error)
	// Edit modifies an existing image based on the given prompt and mask.
	Edit(ImageEditParams) ([]Image, error)
	// CreateVariations generates variations of an existing image.
	CreateVariations(ImageVariationParams) ([]Image, error)
}

type imagesAPI struct {
	c client
}

func (api imagesAPI) Create(params ImageGenerationParams) ([]Image, error) {
	url := fmt.Sprintf("%s%s", baseURL, imageGenerationsAPIEndpoint)
	r, err := api.c.getJSONRequestResponse(url, http.MethodPost, params)
	if err != nil {
		return nil, err
	}

	var response struct {
		Data []Image `json:"data"`
	}

	if err := json.Unmarshal(r, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api imagesAPI) Edit(params ImageEditParams) ([]Image, error) {
	url := fmt.Sprintf("%s%s", baseURL, imageEditsAPIEndpoint)

	return api.imagesFromFormData(url, params)
}

func (api imagesAPI) CreateVariations(params ImageVariationParams) ([]Image, error) {
	url := fmt.Sprintf("%s%s", baseURL, imageVariationsAPIEndpoint)

	return api.imagesFromFormData(url, params)
}

func (api imagesAPI) imagesFromFormData(url string, params interface{}) ([]Image, error) {
	data, contentType, err := structToMultipartFormData(params)
	if err != nil {
		return nil, err
	}

	r, err := api.c.getRequestResponse(url, http.MethodPost, data, contentType)
	if err != nil {
		return nil, err
	}

	var response struct {
		Data []Image `json:"data"`
	}

	if err := json.Unmarshal(r, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}
