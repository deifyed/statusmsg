package finnhub

type quoteResponse struct {
	OpeningPrice float64 `json:"o"`
	CurrentPrice float64 `json:"c"`
}

type quote struct {
	OpeningPrice float64
	CurrentPrice float64
}
