package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type commandFlag struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *commandFlag {
	cf := commandFlag{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo specifying the title")
	flag.StringVar(&cf.Edit, "Edit", "", "Edit a todo by index and specifying the new title")
	flag.IntVar(&cf.Del, "del", -1, " Delete a todo by index and specifying the new title")
	flag.IntVar(&cf.Toggle, "toggle", -1, " Specify a todo by index to toggle")
	flag.BoolVar(&cf.List, "list", false, " List all the todos")
	flag.Parse()
	return &cf
}
func (cf *commandFlag) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid edit command ! Please use id:NewTitle")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error : Invalid Index ")
			os.Exit(1)
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
