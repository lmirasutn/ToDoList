package main

import (
	"ToDoList/cmd" // Asegúrate de que el módulo es correcto
	"ToDoList/internal"
	"log"
)

func main() {

	tarea, err := internal.CargarDatos("tasks.json")
	if err != nil {
		log.Fatalf("Error cargando tareas: %v", err)
	}
	/* Ejemplo 1 - Tarea simple
	 */
	tarea, err = cmd.Add(tarea, "Comprar leche", "2023-05-01")
	if err != nil {
		log.Print(err)
	}

	cmd.List(tarea)
	//tarea = cmd.Delete(0, tarea)
	cmd.Complete(12, tarea)
	cmd.List(tarea)
}
