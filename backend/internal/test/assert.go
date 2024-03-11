package test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertProtoEq(t *testing.T, expected, actual json.Marshaler, msgAndArgs ...any) {
	expPretty, _ := beatifyJSON(expected)
	actPretty, _ := beatifyJSON(actual)

	assert.Equal(t, string(expPretty), string(actPretty), msgAndArgs...)
}

func beatifyJSON(v any) ([]byte, error) {
	var buf bytes.Buffer

	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
