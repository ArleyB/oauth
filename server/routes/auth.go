package routes

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/hlandau/passlib.v1"
	"gopkg.in/hlandau/passlib.v1/abstract"
)

func Compare(pltx string, cptx string) bool{
	passlib.UseDefaults(passlib.DefaultsLatest)
	passlib.DefaultSchemes = []abstract.Scheme{passlib.DefaultSchemes[3]}
	err := passlib.VerifyNoUpgrade(pltx, cptx)
	if err != nil {
		return false
	}
	return true
}

type pw struct {
	Pltx string `json:"pltx"`
	Cptx string `json:"cptx"`
}

func oauth(ctx *gin.Context){
	var pass pw
	ctx.ShouldBindJSON(&pass)
	ctx.JSON(200, gin.H{"data":	Compare(pass.Pltx,pass.Cptx)})
}

func RouteAuth(route *gin.RouterGroup){
	router := route.Group("/auth")
	{
		router.POST("/", oauth)
	}
}