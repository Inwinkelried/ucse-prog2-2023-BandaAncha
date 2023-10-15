package dto

import (
	"time"
)

type Camion struct {
	Id                 string
	PesoMaximo         int
	CostoKm            int
	FechaAlta          time.Time
	FechaActualizacion time.Time
}
