package handler

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type TodoHandler struct{
	Reader *bufio.Reader
}

func NewTodoHandler() *TodoHandler{
  return &TodoHandler{
		Reader: bufio.NewReader(os.Stdin),
	}
}

func (h *TodoHandler)PromptInput()([]string,error){
	input,err:= h.Reader.ReadString('\n')

	if err != nil{
		return nil,err
	}

	fields:= strings.Fields(strings.TrimSpace(input))
	
	return fields,nil
}


func(h *TodoHandler)HandleCommand(fields []string)error{

	if len(fields) == 0 {
		return errors.New("no command provided")
}

	command:= strings.ToLower(fields[0])
	
	
	switch command{
	case "list":
		return h.handleList()
	case "add":

		if len(fields) < 2{
			return errors.New("missing task | Usage: add <task>")
		}

		task:= strings.Join(fields[1:]," ")
		return h.handleAdd(task)
		
	case "update":
		if len(fields) < 3{
			return errors.New("missing ID or task | Usage: update <id> <task>")
		}

	 id:= fields[1]
	 task:= fields[2]

	 return h.handleUpdate(id,task)

	case "delete":
		if len(fields) < 2{
			return errors.New("missing ID | Usage: delete <id> <task>")
		}
		id:= fields[1]

		return h.handleDelete(id)

	case "end":
		fmt.Println("Program ended")
		os.Exit(0)
		return nil
	default:
		return errors.New("invalid command. Try again")
	}

}


func(h *TodoHandler)handleList()error{
  return nil
}

func(h *TodoHandler)handleAdd(task string)error{
	return nil
}
func(h *TodoHandler)handleUpdate(taskId,task string)error{
	return nil
}
func(h *TodoHandler)handleDelete(taskId string)error{
	return nil
}