package controller

import (
	"github.com/daviresio/ajuda-desk-api/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func init() {

	router.Use(cors.Default())

	//router.GET("/empresas/:id")
	//router.GET("/empresas")
	//router.POST("/empresas")
	//router.PUT("/empresas")
	//router.DELETE("/empresas/:id")
	//
	//router.GET("/usuarios/:id")
	//router.GET("/usuarios")
	//router.POST("/usuarios")
	//router.PUT("/usuarios")
	//router.DELETE("/usuarios/:id")

	router.GET("/grupos/:id", service.FindGrupo)
	router.GET("/grupos", service.ListGrupo)
	router.POST("/grupos", service.CreateGrupo)
	router.PUT("/grupos", service.UpdateGrupo)
	router.DELETE("/grupos/:id", service.DeleteGrupo)

}

func GetRouter() *gin.Engine {
	return router
}