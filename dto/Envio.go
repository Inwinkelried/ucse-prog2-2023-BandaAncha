package dto

import (
	"time"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/model"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/utils"
)

type Envio struct {
	ID                string
	IDcamion          string
	Pedidos           []Pedido
	Paradas           []Parada
	Estado            string
	FechaCreacion     time.Time
	FechaModificacion time.Time
}

func NewEnvio(envio model.Envio) *Envio {
	return &Envio{
		ID:                utils.GetStringIDFromObjectID(envio.ID),
		IDcamion:          envio.IDcamion,
		Pedidos:           []Pedido{},
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
		Pedidos:           envio.getPedidos(),
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
func (envio Envio) getPedidos() []model.Pedido {
	var pedidoEnvio []model.Pedido
	for _, pedido := range envio.Pedidos {
		pedidoEnvio = append(pedidoEnvio, pedido.GetModel())
	}
	return pedidoEnvio

}
