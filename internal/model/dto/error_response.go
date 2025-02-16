package dto

type ErrorResponse struct {
	Message string `json:"error"`
	Code    int    `json:"-"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}
