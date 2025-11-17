package handler

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/controller"
)

type todoHandler struct{
	Reader *bufio.Reader
	Controller controller.TodoController
}

type TodoHandler interface{

	HandleList() error
	HandleAdd(fields []string) error
	HandleUpdate(fields []string) error
	HandleDone(fields []string) error
	HandleDelete(fields []string) error
}

func NewTodoHandler(ctrl controller.TodoController) TodoHandler{
  return &todoHandler{
		Reader: bufio.NewReader(os.Stdin),
		Controller: ctrl,
	}
}



func(h *todoHandler)HandleList()error{
	return h.Controller.PrintList()
}

func(h *todoHandler)HandleAdd(fields []string)error{
	if len(fields) == 0{
		return errors.New("missing task | Usage: add <task>")
	}
	taskReq:= strings.Join(fields," ")

	return h.Controller.AddTask(taskReq)
	
}
func(h *todoHandler)HandleUpdate(fields []string)error{
	if len(fields) < 2{
		return errors.New("missing ID or task | Usage: update <id> <task>")
	}

 id:= fields[0]
 task:= strings.Join(fields[1:]," ")

 return h.Controller.HandleUpdate(id,task)

}

func(h *todoHandler)HandleDone(fields []string) error{
	if len(fields) == 0{
		return errors.New("missing ID | Usage: done <id>")
	}

	id:= fields[0]

	return h.Controller.ToggleDone(id)

}


func(h *todoHandler)HandleDelete(fields []string)error{
	if len(fields) < 1{
		return errors.New("missing ID | Usage: delete <id>")
	}

	id:= fields[0]

	return h.Controller.HandleDelete(id)
}