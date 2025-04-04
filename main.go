package main

import (
	"ToDoList/cmd"
	"fmt"
	"os"
)

func main() {

	if err := cmd.Execute(); err != nil {
		fmt.Println("Error ejecutando el programa:", err)
		os.Exit(1)
	}
}
