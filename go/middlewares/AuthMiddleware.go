package middlewares

import (
	"net/http"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/clients"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/utils"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	authClient clients.AuthClientInterface
}

func NewAuthMiddleware(authClient clients.AuthClientInterface) *AuthMiddleware {
	return &AuthMiddleware{
		authClient: authClient,
	}
}

// Este middleware se ejecuta en el grupo de rutas privadas.
func (auth *AuthMiddleware) ValidateToken(c *gin.Context) {
	//Se obtiene el header necesario con nombre "Authorization"
	authToken := c.GetHeader("Authorization")

	if authToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token no encontrado"})
		return
	}
	//Obtener la informacion del usuario a partir del token desde el servicio externo
	user, err := auth.authClient.GetUserInfo(authToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	//Validar rol del usuario
	if (user.Rol != utils.RolAdministrador) && (user.Rol != utils.RolUsuario) && (user.Rol != utils.RolConductor) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "El rol del usuario no es válido"})
		return
	}
	utils.SetUserInContext(c, user)

	c.Next()
}
