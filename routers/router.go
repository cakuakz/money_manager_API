package routers

import (
	"database/sql"
	"money-manager/controller"
	"money-manager/middleware"
	"money-manager/structs"

	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	usr = structs.UserLoginRequest{}
)

func StartServer(db *sql.DB) *gin.Engine {
	router := gin.Default()

	routes := router.Group("/api/user")
	{
		routes.POST("/register", func(ctx *gin.Context) {
			controller.RegisterUser(ctx, db)
		})
		routes.POST("/login", func(ctx *gin.Context) {
			payload, err := controller.LoginUser(ctx, db)
			if err != nil {
				panic(err)
			}
			usr = payload
		})
	}

	authRoutes := router.Group("/api/")
	authRoutes.Use(middleware.BasicAuthMiddleware(usr))
	{
		authRoutes.GET("/secure-route", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "You have access to this route"})
		})
	}

	return router
}