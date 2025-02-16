package dto

type SendCoinRequestDto struct {
	FromUserId uint   `json:"from_user_id"`
	ToUser     string `json:"toUser"`
	Amount     int    `json:"amount"`
}
