package internal

import (
	"fmt"
	"log"
	"time"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	DueDate     string `json:"due_date,omitempty"`
	Completed   bool   `json:"completed"`
}

func Add(tasks []Task, arg string, date int) ([]Task, error) {
	var idtasks int
	if len(tasks) > 0 {
		idtasks = tasks[len(tasks)-1].ID + 1

	} else {
		idtasks = 1
	}

	t := Task{
		ID:          idtasks,
		Description: arg,
		DueDate:     configurarFecha(date),
		Completed:   false,
	}
	tasks = append(tasks, t)
	GuardarTareas("tasks.json", tasks)

	return tasks, nil
}

func Complete(index int, tasks []Task) {
	if index > len(tasks) && index > 0 {
		log.Fatal(" ERROR COMPLETE, No existe esa tasks")
	}
	tasks[index-1].Completed = true

	GuardarTareas("tasks.json", tasks)
}

func Delete(index int, tasks []Task) ([]Task, error) {
	if len(tasks) < index {

		return nil, fmt.Errorf("no existe tal indice %v", index)

	}
	tasks = append(tasks[:index], tasks[index+1:]...)
	// los 3 puntos se usan porque no estamos agregando un elemento, sino que estamos agregando una sublista

	GuardarTareas("tasks.json", tasks)

	return tasks, nil

}

func List(lista []Task) {

	for i := 0; i < len(lista); i++ {
		fmt.Println("Tarea:", lista[i].ID)
		fmt.Println("Descripcion:", lista[i].Description)
		fmt.Println("Fecha de vencimiento:", lista[i].DueDate)

		if lista[i].Completed {
			fmt.Println("Estado: Finalizada")
		} else {
			fmt.Println("Estado: Pendiente")
		}
		fmt.Println("-----------------------")
	}

}

func configurarFecha(num int) string {

	now := time.Now()
	nuevaFecha := now.AddDate(0, 0, num)

	return nuevaFecha.Format("02/01/2006")
}
