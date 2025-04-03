package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

func CargarTareas(filePath string) ([]Task, error) {

	var tarea []Task
	file, err := os.Open(filePath) // abrimos el archivo en una variable file
	if err != nil {
		return nil, fmt.Errorf("error cargando json")
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&tarea) // a newdecoder le pasamos el file y a decode, donde vamos a guardar el json
	if err != nil {

		return nil, fmt.Errorf("error decodificando JSON en  %v", filePath)
	}

	return tarea, nil
}

func GuardarTareas(filePath string, tarea []Task) error {

	data, err := json.MarshalIndent(tarea, "", " ") // marshalIndent, hara que el struct se transforme en json, y lo almacenara en data.
	if err != nil {
		return fmt.Errorf("error de codificacion en %v", filePath)
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		return fmt.Errorf("error con el archivo %s", filePath)
	}

	return nil
}
