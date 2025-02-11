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
func (auth *AuthMiddleware) ValidateToken(c *gin.Context) {
	authToken := c.GetHeader("Authorization")

	if authToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token no encontrado"})
		return
	}
	user, err := auth.authClient.GetUserInfo(authToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	if (user.Rol != utils.RolAdministrador) && (user.Rol != utils.RolUsuario) && (user.Rol != utils.RolConductor) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "El rol del usuario no es válido"})
		return
	}
	utils.SetUserInContext(c, user)

	c.Next()
}
