package main

import (
	"example/pkg/arythmetic"
	"example/pkg/models"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func main() {

	// err := os.RemoveAll("/")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	///////

	if len(os.Args) != 3 {
		fmt.Println("Usage: program <int> <int>")
		os.Exit(1)
	}

	num1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: Number 1 is not an int")
		os.Exit(1)
	}
	num2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error: Number 2 is not an int")
		os.Exit(1)
	}

	a := models.Number{
		N: num1,
	}

	b := models.Number{
		N: num2,
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
