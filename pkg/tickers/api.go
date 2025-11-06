package tickers

import (
	"errors"
	"fmt"
	"time"
)

func GetCurrentPercentage(symbol string) (string, error) {
	hit, err := cacheLookup(symbol)
	if err == nil {
		expiry := hit.Timestamp.Add(lifespan)

		if time.Now().Before(expiry) {
			return hit.Percentage, nil
		}
	} else if !errors.Is(err, errCacheMiss) {
		return "", fmt.Errorf("checking cache: %w", err)
	}

	q, err := getCurrentPercentage(symbol)
	if err != nil {
		return "", fmt.Errorf("getting current percentage: %w", err)
	}

	err = cacheStore(symbol, q.LastUpdated, q.Percentage)
	if err != nil {
		return "", fmt.Errorf("storing quote: %w", err)
	}

	return q.Percentage, nil
}
