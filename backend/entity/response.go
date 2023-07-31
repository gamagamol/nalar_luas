package entity

type ErrorResponse struct {
	Field string
	Tag string
	Value string
}

type ResponseError struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Errors  []ErrorResponse `json:"errors"`
}

type ResponseSuccess struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	SessionId string `json:"sessionId"`
}


