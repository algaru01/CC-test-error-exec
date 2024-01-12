package main

import (
	"example/pkg/arythmetic"
	"example/pkg/models"
	"fmt"

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
}
