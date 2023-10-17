package dto

import (
	"time"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/model"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/utils"
)

type Camion struct {
	ID                string
	Patente           string
	PesoMaximo        int
	CostoKm           int
	FechaCreacion     time.Time
	FechaModificacion time.Time
}

func NewCamion(camion model.Camion) *Camion {
	return &Camion{
		ID:                utils.GetStringIDFromObjectID(camion.ID),
		Patente:           camion.Patente,
		PesoMaximo:        camion.PesoMaximo,
		CostoKm:           camion.CostoKm,
		FechaCreacion:     time.Now(),
		FechaModificacion: time.Now(),
	}
}
func (camion Camion) GetModel() model.Camion {
	return model.Camion{
		ID:                utils.GetObjectIDFromStringID(camion.ID),
		Patente:           camion.Patente,
		PesoMaximo:        camion.PesoMaximo,
		CostoKm:           camion.CostoKm,
		FechaCreacion:     camion.FechaCreacion,
		FechaModificacion: camion.FechaModificacion,
	}
}
