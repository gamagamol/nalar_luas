package entity

type User struct {
	UserId   int    `json:"id,omitempty"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SessionMap struct {
	SessionId string
	UserId    int
}