package domain

type TODO struct {
	Id          int64  `json:"id"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}
