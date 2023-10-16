package dto

import (
	"time"
)

type Camion struct {
	ID                string
	Patente           string
	PesoMaximo        int
	CostoKm           int
	FechaCreacion     time.Time
	FechaModificacion time.Time
}

func NewCamion()
