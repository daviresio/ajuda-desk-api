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

	router.GET("/contatos/:id", service.FindContato)
	router.GET("/contatos", service.ListContato)
	router.POST("/contatos", service.CreateContato)
	router.PUT("/contatos", service.UpdateContato)
	router.DELETE("/contatos/:id", service.DeleteContato)

	router.GET("/empresas/:id", service.FindEmpresa)
	router.GET("/empresas", service.ListEmpresa)
	router.POST("/empresas", service.CreateEmpresa)
	router.PUT("/empresas", service.UpdateEmpresa)
	router.DELETE("/empresas/:id", service.DeleteEmpresa)

	router.GET("/scores/:id", service.FindScore)
	router.GET("/scores", service.ListScore)
	router.POST("/scores", service.CreateScore)
	router.PUT("/scores", service.UpdateScore)
	router.DELETE("/scores/:id", service.DeleteScore)

}

func GetRouter() *gin.Engine {
	return router
}