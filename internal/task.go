package internal

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	DueDate     string `json:"due_date,omitempty"`
	Completed   bool   `json:"completed"`
}
