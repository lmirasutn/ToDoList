package cmd

import (
	"ToDoList/internal"
	"fmt"
)

func List(lista []internal.Task) {

	for i := 0; i < len(lista); i++ {
		fmt.Println("Tarea: ", lista[i].ID)
		fmt.Println("Descripcion: ", lista[i].Description)
		fmt.Println("Fecha de vencimiento: ", lista[i].DueDate)

		if lista[i].Completed {
			fmt.Println("Estado: Finalizada")
		} else {
			fmt.Println("Estado: Pendiente")
		}
		fmt.Println("-----------------------")
	}

}
