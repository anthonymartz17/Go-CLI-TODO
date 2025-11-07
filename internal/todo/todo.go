package todo

type Todo struct {
ID	 int `json:"id"`
Task string `json:"task"`
Done bool `json:"done"`
}

func NewTodo(task string) *Todo {
return &Todo{
	Task: task,
	Done: false,	
}
}
