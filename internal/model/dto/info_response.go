package dto

type Inventory struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

type CoinTransaction struct {
	FromUser string `json:"fromUser,omitempty"`
	ToUser   string `json:"toUser,omitempty"`
	Amount   int    `json:"amount"`
}

type CoinHistory struct {
	Received []CoinTransaction `json:"received,omitempty"`
	Sent     []CoinTransaction `json:"sent,omitempty"`
}

type InfoResponse struct {
	Coins       int         `json:"coins"`
	Inventory   []Inventory `json:"inventory"`
	CoinHistory CoinHistory `json:"CoinHistory,omitempty"`
}
