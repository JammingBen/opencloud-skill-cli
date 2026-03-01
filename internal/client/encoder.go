package client

import (
	"encoding/json"
	"fmt"

	"github.com/alpkeskin/gotoon"
)

type Format int

const (
	TOON Format = iota
	JSON
)

// Encoder is responsible for encoding the response body and status code in the specified format (TOON or JSON)
type Encoder struct {
	format Format
}

func NewEncoder(format Format) *Encoder {
	return &Encoder{format: format}
}

// EncodeBody encodes the response body in the specified format (TOON or JSON)
func (e *Encoder) EncodeBody(body string) (string, error) {
	var data any
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		return "", fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	switch e.format {
	case JSON:
		encoded, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return "", fmt.Errorf("failed to marshal response body: %w", err)
		}
		return string(encoded), nil
	case TOON:
		encoded, err := gotoon.Encode(data)
		if err != nil {
			return "", fmt.Errorf("failed to encode response body: %w", err)
		}

		return encoded, nil
	default:
		return "", fmt.Errorf("unsupported format: %d", e.format)
	}
}

// EncodeStatusCode encodes only the status code in the specified format (TOON or JSON)
func (e *Encoder) EncodeStatusCode(statusCode int) (string, error) {
	output := map[string]int{"status": statusCode}

	switch e.format {
	case JSON:
		encoded, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			return "", fmt.Errorf("failed to marshal status code: %w", err)
		}
		return string(encoded), nil
	case TOON:
		encoded, err := gotoon.Encode(output)
		if err != nil {
			return "", fmt.Errorf("failed to encode status code: %w", err)
		}

		return encoded, nil
	default:
		return "", fmt.Errorf("unsupported format: %d", e.format)
	}
}
