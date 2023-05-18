package routes

import (
	"github.com/gin-gonic/gin"
	"golang_rest_app/controllers"
)

type AuthRoutes struct {
	authController controllers.AuthController
}

func NewAuthRoutes(authController controllers.AuthController) AuthRoutes {
	return AuthRoutes{authController}
}

func (rc *AuthRoutes) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")
	router.POST("/register", rc.authController.SignUpUser)
}
