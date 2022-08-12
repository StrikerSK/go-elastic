package domain

type Todo struct {
	ID          string `json:"id"`
	Time        string `json:"time"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
