package main

import (
	"example/pkg/arythmetic"
	"example/pkg/models"
)

func main() {
	a := models.Number{
		N: 1,
	}

	b := models.Number{
		N: 2,
	}

	c := arythmetic.Sum(a, b)

	println("Result is:", c.N)
}
