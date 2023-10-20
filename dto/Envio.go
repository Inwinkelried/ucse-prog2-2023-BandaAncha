package dto

import (
	"time"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/model"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/utils"
)

type Envio struct {
	ID                string
	IDcamion          string
	IDpedido          string
	Paradas           []Parada
	Estado            string
	FechaCreacion     time.Time
	FechaModificacion time.Time
}

func NewEnvio(envio model.Envio) *Envio {
	return &Envio{
		ID:                utils.GetStringIDFromObjectID(envio.ID),
		IDcamion:          envio.IDcamion,
		IDpedido:          envio.IDpedido,
		Paradas:           []Parada{},
		Estado:            "A despachar",
		FechaCreacion:     time.Now(),
		FechaModificacion: time.Now(),
	}
}
func (envio Envio) GetModel() model.Envio {
	return model.Envio{
		ID:                utils.GetObjectIDFromStringID(envio.ID),
		IDcamion:          envio.IDcamion,
		IDpedido:          envio.IDpedido,
		Paradas:           []model.Parada{},
		Estado:            envio.Estado,
		FechaCreacion:     envio.FechaCreacion,
		FechaModificacion: envio.FechaModificacion,
	}
}
