package gopenai

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"

	"github.com/mitchellh/mapstructure"
)

func structToMultipartFormData(v interface{}) (*bytes.Buffer, string, error) {
	data := &bytes.Buffer{}
	writer := multipart.NewWriter(data)

	structMap := map[string]interface{}{}
	if err := mapstructure.Decode(v, &structMap); err != nil {
		return nil, "", err
	}

	for k, v := range structMap {
		strVal := fmt.Sprintf("%v", v)
		if (k == "image" || k == "mask" || k == "file") && strVal != "" {
			file, err := os.Open(strVal)
			if err != nil {
				return nil, "", err
			}

			part, err := writer.CreateFormFile(k, path.Base(strVal))
			if err != nil {
				return nil, "", err
			}

			if _, err := io.Copy(part, file); err != nil {
				return nil, "", err
			}

			continue
		}

		if err := writer.WriteField(k, strVal); err != nil {
			return nil, "", err
		}
	}

	writer.Close()

	return data, writer.FormDataContentType(), nil
}
