package handler

import (
	"api_gateway/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthInterface interface {
	Login(*gin.Context)
}

type authImplement struct{}

func NewAuth() AuthInterface {
	return &authImplement{}
}

type BodyPayloadAuth struct {
	Username string
	Password string
}

func (a *authImplement) Login(g *gin.Context) {

	bodyPayloadAuth := BodyPayloadAuth{}
	err := g.BindJSON(&bodyPayloadAuth)

	loginMasuk := usecase.NewLogin().Autentikasi(bodyPayloadAuth.Username, bodyPayloadAuth.Password)

	if loginMasuk {
		g.JSON(http.StatusOK, gin.H{
			"message": "Anda berhasil login",
			"data":    bodyPayloadAuth,
		})
	} else {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "Anda gagal login",
			"data":    err,
		})
	}
}
