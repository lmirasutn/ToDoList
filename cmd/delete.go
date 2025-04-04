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

var (
	deleteAll       bool
	deleteCompleted bool
	deletePending   bool
)
var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Eliminar una tarea",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tarea, err := internal.CargarTareas("tasks.json")
		if err != nil {
			color.Red("ha ocurrido un error")
			return
		}
		if deleteAll {
			internal.DeleteAllTasks(tarea)
			color.Green("Todas las tareas han sido eliminadas correctamente.")
			return
		}

		if deleteCompleted {
			internal.DeleteCompletedTasks(tarea)
			color.Green("Las tareas completadas han sido eliminadas correctamente.")
			return
		}

		// Si se pasó un ID como argumento
		if len(args) == 1 {
			idTarea := args[0]

			id, err := strconv.Atoi(idTarea)
			if err != nil {
				fmt.Printf("Error: el segundo argumento debe ser un número entero (recibido: '%s')\n", idTarea)
				return
			}
			tarea, err = internal.Delete(id, tarea)
			if err != nil {
				color.Red("No existe esa tarea")
			} else {
				color.Yellow("Tarea %v eliminada", id)
			}
		} else {
			fmt.Println("Debe proporcionar un ID o una bandera (--all, --completed)")
		}
	},
}

func init() {
	deleteCmd.Flags().BoolVar(&deleteAll, "all", false, "Eliminar todas las tareas")
	deleteCmd.Flags().BoolVar(&deleteCompleted, "completed", false, "Eliminar tareas completadas")
	rootCmd.AddCommand(deleteCmd)
}
