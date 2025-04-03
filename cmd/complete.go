/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"ToDoList/internal"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Completar una tarea",
	Run: func(cmd *cobra.Command, args []string) {
		tarea, err := internal.CargarTareas("tasks.json")
		if err != nil {
			fmt.Printf("ha ocurrido un error")
			return
		}
		idTarea := args[0]
		// Convertir el segundo argumento a número
		id, err := strconv.Atoi(idTarea)
		if err != nil {
			fmt.Printf("Error: el segundo argumento debe ser un número entero (recibido: '%s')\n", idTarea)
			return
		}
		internal.Complete(id, tarea)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
