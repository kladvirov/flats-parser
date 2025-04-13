package parser

import (
	"encoding/json"
	"io"
	"net/http"
)

func Parse[T any](url string) (T, error) {
	response, err := http.Get(url)
	if err != nil {
		return *new(T), err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return *new(T), err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return *new(T), err
	}

	var result T
	if err := json.Unmarshal(body, &result); err != nil {
		return *new(T), err
	}

	return result, nil
}
