package model

type PersonLocation struct {
	Person   Person   `json:"person"`
	Location Location `json:"location"`
}
