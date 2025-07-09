package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/dbedggood/pokedexcli/internal/pokecache"
)

var cache *pokecache.Cache

func Fetch[T any](url string, resPointer *T) error {
	if url == "" || resPointer == nil {
		return errors.New("args cannot be empty")
	}

	if cache == nil {
		cache = pokecache.NewCache(time.Second * 5)
	}
	if cachedData, exists := cache.Get(url); exists {
		if err := json.Unmarshal(cachedData, resPointer); err != nil {
			return err
		}
		return nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status code")
	}

	fetchedData, err := io.ReadAll(res.Body)
	if err := json.Unmarshal(fetchedData, resPointer); err != nil {
		return err
	}

	cache.Add(url, fetchedData)
	return nil
}
