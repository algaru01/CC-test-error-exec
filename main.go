package main

import (
	"example/pkg/arythmetic"
	"example/pkg/models"

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

	println("Result is:", c.N)
	println("ID is:", id.String())
}
