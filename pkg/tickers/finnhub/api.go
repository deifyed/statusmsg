// Package finnhub handles retrieving stock information using the finnhub API
package finnhub

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

var ErrUnrecoverable = errors.New("unrecoverable")

const defaultBaseURL = "https://finnhub.io/api/v1"

type Client struct {
	Key string
}

func (client Client) GetQuote(symbol string) (quote, error) {
	targetURL, err := url.ParseRequestURI(fmt.Sprintf("%s/quote", defaultBaseURL))
	if err != nil {
		return quote{}, fmt.Errorf("parsing API URI: %w", err)
	}

	params := url.Values{}
	params.Add("symbol", symbol)
	targetURL.RawQuery = params.Encode()

	request, err := http.NewRequest(http.MethodGet, targetURL.String(), nil)
	if err != nil {
		return quote{}, fmt.Errorf("creating request: %w", err)
	}

	request.Header.Add("X-Finnhub-Token", client.Key)

	httpClient := http.Client{}
	rawResponse, err := httpClient.Do(request)
	if err != nil {
		return quote{}, fmt.Errorf("executing request: %w", err)
	}

	if rawResponse.StatusCode != http.StatusOK {
		if rawResponse.StatusCode == http.StatusUnauthorized {
			return quote{}, fmt.Errorf("unauthorized: %w", ErrUnrecoverable)
		}

		return quote{}, fmt.Errorf("response not OK, got %d", rawResponse.StatusCode)
	}

	var response quoteResponse

	err = json.NewDecoder(rawResponse.Body).Decode(&response)
	if err != nil {
		return quote{}, fmt.Errorf("decoding response: %w", err)
	}

	return quote(response), nil
}

func (q quote) Percentage() string {
	percentage := (q.CurrentPrice - q.OpeningPrice) / q.OpeningPrice * 100

	return fmt.Sprintf("%.2f", percentage)
}
