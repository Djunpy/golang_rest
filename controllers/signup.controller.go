package controllers

import (
	"github.com/gin-gonic/gin"
	db "golang_rest_app/db/sqlc"
	"golang_rest_app/utils"
	"net/http"
)

type AuthController struct {
	db *db.Queries
}

func NewAuthController(db *db.Queries) *AuthController {
	return &AuthController{db}
}

func (ac *AuthController) SignUpUser(ctx *gin.Context) {
	var credentials *db.User

	if err := ctx.BindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	hashedPassword := utils.HashPassword(credentials.Password)

	args := &db.CreateUserParams{
		Name:     credentials.Name,
		Email:    credentials.Email,
		Password: hashedPassword,
		Verified: true,
		Role:     "is_active",
	}
	user, err := ac.db.CreateUser(ctx, *args)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": user}})
}
