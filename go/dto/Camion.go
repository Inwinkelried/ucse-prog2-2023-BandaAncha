package dto

import (
	"time"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/model"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/utils"
)

type Camion struct {
	ID                string    `json:"id,omitempty"`
	Patente           string    `json:"patente"`
	PesoMaximo        int       `json:"peso_maximo"`
	CostoKm           int       `json:"costo_km"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
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
