package dto

import (
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/model"
)

type Parada struct {
	Ciudad       string
	KmRecorridos float64
}

func NewParada(parada model.Parada) *Parada {
	return &Parada{
		Ciudad:       parada.Ciudad,
		KmRecorridos: parada.KmRecorridos,
	}
}

func (parada Parada) GetModel() model.Parada {
	return model.Parada{
		Ciudad:       parada.Ciudad,
		KmRecorridos: parada.KmRecorridos,
	}
}
