package domain

type Todo struct {
	ID          string   `json:"id"`
	CreatedAt   string   `json:"createdAt"`
	FinishedAt  string   `json:"finishedAt"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags,omitempty"`
	Done        bool     `json:"done"`
}

func (r *Todo) CheckDone() {
	if r.Done {
		r.FinishedAt = r.CreatedAt
	}
}
