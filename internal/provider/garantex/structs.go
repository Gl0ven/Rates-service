package garantex

type Rates struct {
	Timestamp int64 `json:"timestamp"`
	Asks      []*Parameters `json:"asks"`
	Bids      []*Parameters `json:"bids"`
}

type Parameters struct {
	Price  string `json:"price"`
	Volume string `json:"volume"`
	Amount string `json:"amount"`
	Factor string `json:"factor"`
	Type   string `json:"type"`
}