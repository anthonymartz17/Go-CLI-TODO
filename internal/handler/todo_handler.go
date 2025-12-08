package handler

import (
	"errors"
	"strings"

	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/controller"
)

//TodoController defines the behavior for managing todo items.
type TodoController interface{
	PrintList() error
	AddTask(task string) error
	UpdateTask(taskId,task string) error
	ToggleDone(taskId string)error
	HandleDelete(taskId string) error

}

//Ensures *controller.TodoController implements TodoController interface
var _ TodoController = (*controller.TodoController)(nil)


type TodoHandler struct{
	controller TodoController
}



func New(ctrl TodoController) *TodoHandler{
  return &TodoHandler{
		controller: ctrl,
	}
}



func(h *TodoHandler)HandleList()error{
	return h.controller.PrintList()
}

func(h *TodoHandler)HandleAdd(args []string)error{
	if len(args) == 0{
		return errors.New("missing task | Usage: add <task>")
	}
	taskReq:= strings.Join(args," ")

	return h.controller.AddTask(taskReq)

}
func(h *TodoHandler)HandleUpdate(args []string)error{
	if len(args) < 2{
		return errors.New("missing ID or task | Usage: update <id> <task>")
	}

 id:= args[0]
 task:= strings.Join(args[1:]," ")

 return h.controller.UpdateTask(id,task)

}

func(h *TodoHandler)HandleDone(args []string) error{
	if len(args) == 0{
		return errors.New("missing ID | Usage: done <id>")
	}

	id:= args[0]

	return h.controller.ToggleDone(id)

}


func(h *TodoHandler)HandleDelete(args []string)error{
	if len(args) == 0 {
		return errors.New("missing ID | Usage: delete <id>")
	}

	id:= args[0]

	return h.controller.HandleDelete(id)
}