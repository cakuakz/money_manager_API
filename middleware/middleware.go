package middleware

import (
	"money-manager/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware(user structs.UserLoginRequest) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uname, pass, ok := ctx.Request.BasicAuth()
		if !ok || uname != user.Username || pass != user.Password {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized: invalid credentials"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}