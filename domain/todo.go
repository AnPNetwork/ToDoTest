package domain

type TODO struct {
	Id          int64  `json:"id"`
	Description string `json:"desc"`
	Do          bool   `json:"do"`
}
