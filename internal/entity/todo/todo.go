package todo

import "github.com/anthonymartz17/Go-CLI-TODO.git/internal/util"

type Todo struct {
Id	 string `json:"id"`
Task string `json:"task"`
Done bool `json:"done"`
}

func NewTodo(task string) *Todo {
return &Todo{
	Id: util.GenerateID(),
	Task: task,
	Done: false,	
}
}
