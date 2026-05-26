package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CMDFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCMDFlags() *CMDFlags {
	cf := CMDFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo, enter title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by specifying index and new title. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Enter id of Todo to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Enter id of Todo to toggle")
	flag.BoolVar(&cf.List, "list", false, "Show all todos")

	flag.Parse()

	return &cf
}

func (cf *CMDFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit. Please use id:new_title")
			os.Exit(1)
			return
		}
		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: Invalid index for edit")
			os.Exit(1)
			return
		}

		todos.edit(index, parts[1])
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	case cf.Del != -1:
		todos.delete(cf.Del)
	default:
		fmt.Println("Invalid Command")

	}
}
