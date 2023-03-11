package gopenai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
)

const (
	errMsgTpl                    = "Message: %s | Type: %s | Code: %s | Param: %v"
	contentTypeJSON              = "application/json"
	headerNameContentType        = "Content-Type"
	headerNameAuthorization      = "Authorization"
	headerNameOpenAIOrganization = "OpenAI-Organization"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type client struct {
	cfg        Config
	httpClient httpClient
}

func (c client) Models() ModelsAPI {
	return modelsAPI{c: c}
}

func (c client) ChatCompletions() ChatCompletionsAPI {
	return chatCompletionsAPI{c: c}
}

func (c client) Completions() CompletionsAPI {
	return completionsAPI{c: c}
}

func (c client) Edits() EditsAPI {
	return editsAPI{c: c}
}

func (c client) Images() ImagesAPI {
	return imagesAPI{c: c}
}

func (c client) Embeddings() EmbeddingsAPI {
	return embeddingsAPI{c: c}
}

func (c client) Files() FilesAPI {
	return filesAPI{c: c}
}

func (c client) FineTunes() FineTunesAPI {
	return fineTunesAPI{c: c}
}

func (c client) Moderations() ModerationsAPI {
	return moderationsAPI{c: c}
}

func (c *client) getHTTPResponse(url, method string, data io.Reader, contentType string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	if contentType != "" {
		req.Header.Add(headerNameContentType, contentType)
	}

	req.Header.Add(headerNameAuthorization, fmt.Sprintf("Bearer %s", c.cfg.APIKey))
	if c.cfg.OrganizationID != "" {
		req.Header.Add(headerNameOpenAIOrganization, c.cfg.OrganizationID)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
			return nil, ErrRequestTimeout
		}

		return nil, err
	}

	return resp, err
}

func (c *client) streamRequestResponse(url, method string, data io.Reader, contentType string, dst io.Writer) error {
	resp, err := c.getHTTPResponse(url, method, data, contentType)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(dst, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) getRequestResponse(url, method string, data io.Reader, contentType string) ([]byte, error) {
	resp, err := c.getHTTPResponse(url, method, data, contentType)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// check if response is an error
	if resp.StatusCode >= 400 {
		var errResp struct {
			Error struct {
				Message string      `json:"message"`
				Type    string      `json:"type"`
				Param   interface{} `json:"param"`
				Code    string      `json:"code"`
			} `json:"error"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return nil, err
		}

		errMsg := fmt.Sprintf(errMsgTpl, errResp.Error.Message,
			errResp.Error.Type, errResp.Error.Code, errResp.Error.Param)

		return nil, errors.New(errMsg)
	}

	return io.ReadAll(resp.Body)
}

func (c client) getJSONRequestResponse(url, method string, data interface{}) ([]byte, error) {
	var reqBody io.Reader
	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		reqBody = bytes.NewReader(jsonData)
	}

	return c.getRequestResponse(url, method, reqBody, contentTypeJSON)
}
