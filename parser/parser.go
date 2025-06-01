package parser

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func Parse[T any](url string) (T, error) {
	var res T

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return res, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Smth/1.0)")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return res, errors.New("unexpected status: " + resp.Status)
	}

	var r io.Reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gr, err := gzip.NewReader(resp.Body)
		if err != nil {
			return res, err
		}
		defer gr.Close()
		r = gr
	}

	if err := json.NewDecoder(r).Decode(&res); err != nil {
		return res, err
	}
	return res, nil
}
