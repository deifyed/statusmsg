package tickers

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"time"
)

const (
	defaultStoreFilepath        = "/home/julius/.local/state/statusbar/tickers"
	defaultStoreFilePermissions = 0o755
	lifespan                    = 1 * time.Hour
)

var errCacheMiss = errors.New("cache miss")

type cacheHit struct {
	Percentage string
	Timestamp  time.Time
}

type cacheElement struct {
	Symbol     string    `json:"symbol"`
	Timestamp  time.Time `json:"timestamp"`
	Percentage string    `json:"percentage"`
}

func cacheLookup(symbol string) (cacheHit, error) {
	err := ensureDirectory(defaultStoreFilepath)
	if err != nil {
		return cacheHit{}, fmt.Errorf("ensuring directory: %w", err)
	}

	targetFile := path.Join(defaultStoreFilepath, symbol)

	f, err := os.Open(targetFile)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return cacheHit{}, fmt.Errorf("opening file: %w", err)
		}

		return cacheHit{}, errCacheMiss
	}

	var element cacheElement

	err = json.NewDecoder(f).Decode(&element)
	if err != nil {
		return cacheHit{}, fmt.Errorf("decoding: %w", err)
	}

	return cacheHit{
		Percentage: element.Percentage,
		Timestamp:  element.Timestamp,
	}, nil
}

func cacheStore(symbol string, timestamp time.Time, percentage string) error {
	err := ensureDirectory(defaultStoreFilepath)
	if err != nil {
		return fmt.Errorf("ensuring directory: %w", err)
	}

	targetFile := path.Join(defaultStoreFilepath, symbol)

	f, err := os.OpenFile(targetFile, os.O_WRONLY|os.O_CREATE, defaultStoreFilePermissions)
	if err != nil {
		return fmt.Errorf("opening file: %w", err)
	}

	element := cacheElement{Symbol: symbol, Timestamp: timestamp, Percentage: percentage}

	err = json.NewEncoder(f).Encode(&element)
	if err != nil {
		return fmt.Errorf("encoding: %w", err)
	}

	return nil
}

func ensureDirectory(dir string) error {
	err := os.MkdirAll(dir, 0o755)
	if err != nil {
		return fmt.Errorf("creating directory: %w", err)
	}

	return nil
}
