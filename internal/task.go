package internal

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
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

func Delete(idTarea int, tasks []Task) ([]Task, error) {
	indexFound := -1
	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == idTarea {
			indexFound = i
			break
		}
	}
	if indexFound >= 0 {
		tasks = append(tasks[:indexFound], tasks[indexFound+1:]...)
	} else {
		return nil, fmt.Errorf("no existe esa tarea")
	}
	// los 3 puntos se usan porque no estamos agregando un elemento, sino que estamos agregando una sublista

	GuardarTareas("tasks.json", tasks)

	return tasks, nil

}

func DeleteAllTasks(tasks []Task) error {
	tasks = []Task{}
	GuardarTareas("tasks.json", tasks)

	return nil

}

func DeleteCompletedTasks(tasks []Task) ([]Task, error) {
	// Crear un nuevo slice solo con tareas NO completadas
	var newTasks []Task
	for _, task := range tasks {
		if !task.Completed {
			newTasks = append(newTasks, task)
		}
	}

	// Guardar los cambios en el archivo
	err := GuardarTareas("tasks.json", newTasks)
	if err != nil {
		return nil, err
	}

	return newTasks, nil
}
func List(lista []Task) {
	if len(lista) == 0 {
		color.Yellow("No hay tareas para mostrar.")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Descripción", "Estado", "Fecha Límite"})

	for _, task := range lista {
		estado := color.YellowString("Pendiente")
		if task.Completed {
			estado = color.GreenString("Finalizada")
		}

		table.Append([]string{
			fmt.Sprintf("%d", task.ID),
			task.Description,
			estado,
			task.DueDate,
		})
	}

	table.SetBorder(true)
	table.SetRowLine(true)
	table.Render()
}

func configurarFecha(num int) string {

	now := time.Now()
	nuevaFecha := now.AddDate(0, 0, num)

	return nuevaFecha.Format("02/01/2006")
}
