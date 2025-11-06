// Package tickers handles retrieving information regarding stock tickers
package tickers

import (
	"errors"
	"fmt"
	"time"

	"github.com/deifyed/statusmsg/pkg/tickers/finnhub"
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
		if !errors.Is(err, finnhub.ErrUnrecoverable) {
			return "", fmt.Errorf("getting current percentage: %w", err)
		}

		// Hack to prevent spamming API upon unrecoverable error
		q.LastUpdated = time.Now().Add(12 * time.Hour)
		q.Percentage = "err"
	}

	err = cacheStore(symbol, q.LastUpdated, q.Percentage)
	if err != nil {
		return "", fmt.Errorf("storing quote: %w", err)
	}

	return q.Percentage, nil
}
