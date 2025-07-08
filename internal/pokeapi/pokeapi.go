package pokeapi

import (
	"encoding/json"
	"errors"
	"net/http"
)

func Fetch[T any](url string, resPointer *T) error {
	if resPointer == nil {
		return errors.New("resPointer cannot be nil")
	}

	if url == "" {
		return errors.New("url cannot be empty")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("unexpected status code")
	}

	if err := json.NewDecoder(resp.Body).Decode(resPointer); err != nil {
		return err
	}

	return nil
}
