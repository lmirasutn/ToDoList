package cmd

import (
	"github.com/spf13/cobra"
)

// Definir el comando raíz
var rootCmd = &cobra.Command{

	Use:   "ToDoList",
	Short: "Una aplicación de lista de tareas en Go",
	Long:  "ToDoList es una aplicación CLI para gestionar tareas de manera sencilla.",
}

// Ejecutar el comando raíz y devolver error
func Execute() error {

	return rootCmd.Execute()
}
