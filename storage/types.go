package storage

type TodoNewRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Todo struct {
	ID          int    `json:"ID"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewTodo(title string, description string) *Todo {
	return &Todo{
		Title:       title,
		Description: description,
	}
}
