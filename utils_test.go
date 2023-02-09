package gopenai

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	File  string `mapstructure:"file,omitempty"`
	Field string `mapstructure:"field,omitempty"`
}

func TestStructToMultipartFormData(t *testing.T) {
	testCases := []struct {
		name                 string
		input                interface{}
		expectedDataContains []string
		expectedContentType  string
		expectErr            bool
	}{
		{
			name: "inputs properly included",
			input: testStruct{
				File:  "./.fixture/test-file.jsonl",
				Field: "test",
			},
			expectedDataContains: []string{
				`Content-Disposition: form-data; name="file"; filename="test-file.jsonl"`,
				`Content-Type: application/octet-stream`,
				`{"prompt": "this is a prompt", "completion": "this is a completion"}`,
				`Content-Disposition: form-data; name="field"`,
				`test`,
			},
			expectedContentType: "multipart/form-data",
			expectErr:           false,
		},
		{
			name: "invalid file path",
			input: testStruct{
				File:  "/invalid/path/to/file.txt",
				Field: "Test",
			},
			expectedDataContains: nil,
			expectedContentType:  "",
			expectErr:            true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, contentType, err := structToMultipartFormData(tc.input)
			dataStr := data.String()

			dataContainsExpectedStrings := true
			for _, expectedStr := range tc.expectedDataContains {
				if !strings.Contains(dataStr, expectedStr) {
					dataContainsExpectedStrings = false

					break
				}
			}

			assert.True(t, dataContainsExpectedStrings)
			assert.True(t, strings.HasPrefix(contentType, tc.expectedContentType))

			if tc.expectErr {
				assert.NotNil(t, err)
			}
		})
	}
}
