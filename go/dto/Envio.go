package dto

import (
	"time"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/model"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/utils"
)

type Envio struct {
	ID                string    `json:"id,omitempty"`
	IDcamion          string    `json:"id_camion"`
	Pedidos           []string  `json:"pedidos"`
	Paradas           []Parada  `json:"paradas"`
	Estado            string    `json:"estado"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}

func NewEnvio(envio model.Envio) *Envio {
	return &Envio{
		ID:                utils.GetStringIDFromObjectID(envio.ID),
		IDcamion:          envio.IDcamion,
		Pedidos:           []string{},
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
		Pedidos:           envio.Pedidos,
		Paradas:           envio.getParadas(),
		Estado:            envio.Estado,
		FechaCreacion:     envio.FechaCreacion,
		FechaModificacion: envio.FechaModificacion,
	}
}
func (envio Envio) getParadas() []model.Parada {
	var paradasEnvio []model.Parada
	for _, parada := range envio.Paradas {
		paradasEnvio = append(paradasEnvio, parada.GetModel())
	}
	return paradasEnvio
}
