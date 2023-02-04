package pkg

type TodoItem struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

// todo slice to seed record todo data.
var todos = []TodoItem{
	{
		ID:          "1",
		Description: "Learn Go language",
		Priority:    1,
	},
	{
		ID:          "2",
		Description: "Participate in Open Source with Go",
		Priority:    3,
	},
	{
		ID:          "3",
		Description: "Create sample project with Go",
		Priority:    1,
	},
}