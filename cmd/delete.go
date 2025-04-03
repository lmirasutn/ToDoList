package cmd

import (
	"ToDoList/internal"
	"encoding/json"
	"fmt"
	"os"
)

func Delete(index int, tasks []internal.Task) []internal.Task {
	if len(tasks) < index {

		fmt.Errorf("No existe tal indice")
	}
	tasks = append(tasks[:index], tasks[index+1:]...)
	// los 3 puntos se usan porque no estamos agregando un elemento, sino que estamos agregando una sublista

	data, err := json.MarshalIndent(tasks, "", " ") // basicamente, vuelvo a escribir el archivo, ahora sin el elemento eliminado.

	if err != nil {
		fmt.Errorf("Error en la decodificacion de json")

	}
	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Errorf("Error en la decodificacion de json")

	}
	return tasks

}
