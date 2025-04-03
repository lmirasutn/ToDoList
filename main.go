package main

import (
	"ToDoList/cmd" // Asegúrate de que el import tenga el nombre correcto del módulo
	"fmt"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error ejecutando el programa:", err)
		os.Exit(1)
	}
}
