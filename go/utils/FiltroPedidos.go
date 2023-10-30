package utils

import "time"

type FiltroPedido struct {
	CodigoEnvio string
	Estado      string
	FechaMayor  time.Time
	FechaMenor  time.Time
}
