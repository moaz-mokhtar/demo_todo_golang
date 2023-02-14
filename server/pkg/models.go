package pkg

type TodoItem struct {
	Id          int `json:"id"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

type TodoDTO struct {
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}
