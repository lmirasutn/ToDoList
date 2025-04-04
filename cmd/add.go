/*
Copyright © 2025 NAME HERE lmiras@frba.utn.edu.ar
*/
package cmd

import (
	"ToDoList/internal"
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <desc> <diasVenc>",
	Short: "Agrega una tarea",
	Run: func(cmd *cobra.Command, args []string) {
		tarea, err := internal.CargarTareas("tasks.json")
		if err != nil {
			color.Red("error al cargar tareas")
		}
		diasVenc := args[1]
		// Convertir el segundo argumento a número
		numVenc, err := strconv.Atoi(diasVenc)
		if err != nil {
			fmt.Printf("Error: el segundo argumento debe ser un número entero (recibido: '%s')\n", diasVenc)
			return
		}
		tarea, err = internal.Add(tarea, args[0], numVenc)
		if err != nil {
			color.Red("error al añadir tarea")
		}
		color.Green("Se agrego la tarea correctamente")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

}
