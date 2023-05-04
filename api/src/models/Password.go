package models

type Password struct {
	Actual string `json:"actual"`
	NewOne string `json:"newOne"`
}
