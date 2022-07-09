package models

type Record struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Expiration int64    `json:"expiration"`
}
