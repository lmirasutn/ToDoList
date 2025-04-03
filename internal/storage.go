package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

func CargarTareas(filePath string) ([]Task, error) {
	var tarea []Task

	// Verificar si el archivo existe
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("Archivo no encontrado, creando uno nuevo...")
		return []Task{}, nil // Retorna un slice vacío pero sin error
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error abriendo el archivo %s: %w", filePath, err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&tarea)
	if err != nil {
		return nil, fmt.Errorf("error decodificando JSON en %v: %w", filePath, err)
	}

	return tarea, nil
}
func GuardarTareas(filePath string, tarea []Task) error {

	data, err := json.MarshalIndent(tarea, "", " ") // marshalIndent, hara que el struct se transforme en json, y lo almacenara en data.
	if err != nil {
		return fmt.Errorf("error de codificacion en %v", filePath)
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("error con el archivo %s", filePath)
	}

	return nil
}
