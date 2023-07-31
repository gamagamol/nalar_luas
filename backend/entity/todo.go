package entity

type Todo struct {
	TodoId int    `json:"id,omitempty"`
	UserId int    `json:"user_id,omitempty"`
	Text   string `json:"todo" validate:"required"`
}