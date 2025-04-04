/*
Copyright © 2025 NAME HERE lmiras@frba.utn.edu.ar
*/
package cmd

import (
	"ToDoList/internal"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista todas las tareas",
	Run: func(cmd *cobra.Command, args []string) {
		tarea, err := internal.CargarTareas("tasks.json")
		if err != nil {
			color.Red("ha ocurrido un error")
			return
		}
		internal.List(tarea)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
