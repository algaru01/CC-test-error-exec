package arythmetic

import "example/pkg/models"

func Sum(a models.Number, b models.Number) models.Number {
	return models.Number{N: a.N + b.N}
}
