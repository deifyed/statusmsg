package finnhub

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const defaultBaseURL = "https://finnhub.io/api/v1"

type Client struct {
	Key string
}

type quote struct {
	OpeningPrice float64
	CurrentPrice float64
}

type quoteResponse struct {
	OpeningPrice float64 `json:"o"`
	CurrentPrice float64 `json:"c"`
}

func (q quote) Percentage() string {
	percentage := (q.CurrentPrice - q.OpeningPrice) / q.OpeningPrice * 100

	return fmt.Sprintf("%.2f", percentage)
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

	var response quoteResponse

	err = json.NewDecoder(rawResponse.Body).Decode(&response)
	if err != nil {
		return quote{}, fmt.Errorf("decoding response: %w", err)
	}

	return quote(response), nil
}
