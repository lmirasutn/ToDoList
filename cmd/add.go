/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"ToDoList/internal"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <desc> <diasVenc>",
	Short: "Agrega una tarea",
	Run: func(cmd *cobra.Command, args []string) {
		tarea, err := internal.CargarTareas("tasks.json")
		if err != nil {
			fmt.Printf("error al cargar tareas")
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
			fmt.Printf("error al añadir tarea")
		}
		log.Println("Se agrego la tarea correctamente")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
