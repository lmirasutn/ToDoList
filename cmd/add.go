package cmd

import (
	"ToDoList/internal" // 👈 CORRECTO, sin "/"
)

/********************* ESTRATEGIA *******************
1. Crear una constante de tareas, con id.
2.
*/

// Función para agregar una tare

func Add(tarea []internal.Task, arg string, date string) ([]internal.Task, error) {
	var idTarea int
	if len(tarea) > 0 {
		idTarea = tarea[len(tarea)-1].ID + 1

	} else {
		idTarea = 1
	}

	t := internal.Task{
		ID:          idTarea,
		Description: arg,
		DueDate:     date,
		Completed:   false,
	}
	tarea = append(tarea, t)

	internal.GuardarTareas("tasks.json", tarea)

	return tarea, nil

}
