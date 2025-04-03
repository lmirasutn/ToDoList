package cmd

import (
	"ToDoList/internal"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Complete(index int, tasks []internal.Task) {
	if index > len(tasks) && index > 0 {
		log.Fatal(" ERROR COMPLETE, No existe esa tarea")
	}
	tasks[index-1].Completed = true

	data, err := json.MarshalIndent(tasks, "", " ") // basicamente, vuelvo a escribir el archivo, ahora sin el elemento eliminado.

	if err != nil {
		fmt.Errorf("Error en la decodificacion de json")

	}
	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Errorf("Error en la decodificacion de json")

	}
}
