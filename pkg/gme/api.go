package gme

import (
	"encoding/json"
	"fmt"
	"os"
)

func (q Quote) String() string {
	icon := "💤"

	if q.CurrentPrice > q.OpeningPrice {
		icon = "🚀"
	}

	return fmt.Sprintf("%s: %s%.2f", q.Symbol, icon, q.CurrentPrice)
}

func GetStatus() (Quote, error) {
	rawQuote, err := os.ReadFile("/tmp/cached-quote")
	if err != nil {
		return Quote{}, fmt.Errorf("reading quote: %w", err)
	}

	var quote Quote

	err = json.Unmarshal(rawQuote, &quote)
	if err != nil {
		return Quote{}, fmt.Errorf("unmarshalling quote: %w", err)
	}

	return quote, nil
}
