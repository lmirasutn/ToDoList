/*
Copyright © 2025 NAME HERE lmiras@frba.utn.edu.ar
*/
package cmd

import (
	"ToDoList/internal"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Completar una tarea",
	Run: func(cmd *cobra.Command, args []string) {
		tarea, err := internal.CargarTareas("tasks.json")
		if err != nil {
			color.Red("ha ocurrido un error")
			return
		}
		idTarea := args[0]

		id, err := strconv.Atoi(idTarea)
		if err != nil {
			color.Red("Error: el segundo argumento debe ser un número entero (recibido: '%s')\n", idTarea)
			return
		}
		internal.Complete(id, tarea)
		color.Green("Tarea completada!")

	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

}
