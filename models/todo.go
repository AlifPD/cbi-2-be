package models

type Todo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Owner   string `json:"owner"`
	Checked bool   `json:"checked"`
}
