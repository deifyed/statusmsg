package tickers

import (
	"fmt"
	"os"
	"time"

	"github.com/deifyed/statusmsg/pkg/tickers/finnhub"
)

const defaultTickerAPIEnvironmentKey = "TICKER_API_KEY"

type quote struct {
	Percentage  string
	LastUpdated time.Time
}

func getCurrentPercentage(symbol string) (quote, error) {
	apiClient := finnhub.Client{Key: getKey()}

	apiQuote, err := apiClient.GetQuote(symbol)
	if err != nil {
		return quote{}, fmt.Errorf("fetching quote: %w", err)
	}

	q := quote{
		Percentage:  apiQuote.Percentage(),
		LastUpdated: time.Now(),
	}

	return q, nil
}

func getKey() string {
	key, ok := os.LookupEnv(defaultTickerAPIEnvironmentKey)
	if !ok {
		return "n/a"
	}

	return key
}
