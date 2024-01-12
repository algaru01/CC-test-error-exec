package main

import (
	"example/pkg/arythmetic"
	"example/pkg/models"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

func main() {
	a := models.Number{
		N: 1,
	}

	b := models.Number{
		N: 2,
	}

	c := arythmetic.Sum(a, b)

	id := uuid.New()

	fmt.Println("Result is:", c.N)
	println("ID is:", id.String())

	nombreArchivo := "ejemplo.txt"

	// Crear o truncar el archivo
	archivo, err := os.Create(nombreArchivo)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer archivo.Close()

	fmt.Println("Archivo creado exitosamente:", nombreArchivo)

	time.Sleep(10 * time.Second)

	fmt.Println("me voy")
}
