package dto

import (
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/clients/responses"
)

type Usuario struct {
	Codigo   string `json:"codigo"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Rol      string `json:"rol"`
}

func NewUser(userInfo *responses.UserInfo) Usuario {
	usuario := Usuario{}
	if userInfo != nil {
		usuario.Codigo = userInfo.Codigo
		usuario.Username = userInfo.Username
		usuario.Email = userInfo.Email
		usuario.Rol = userInfo.Rol
	}
	return usuario
}
