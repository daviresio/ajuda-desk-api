package routes

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

	router.GET("/tipo-empresas/:id", service.FindTipoEmpresa)
	router.GET("/tipo-empresas", service.ListTipoEmpresa)
	router.POST("/tipo-empresas", service.CreateTipoEmpresa)
	router.PUT("/tipo-empresas", service.UpdateTipoEmpresa)
	router.DELETE("/tipo-empresas/:id", service.DeleteTipoEmpresa)

	router.GET("/tipo-planos/:id", service.FindTipoPlano)
	router.GET("/tipo-planos", service.ListTipoPlano)
	router.POST("/tipo-planos", service.CreateTipoPlano)
	router.PUT("/tipo-planos", service.UpdateTipoPlano)
	router.DELETE("/tipo-planos/:id", service.DeleteTipoPlano)

	router.GET("/anotacoes/:id", service.FindAnotacao)
	router.GET("/anotacoes", service.ListAnotacao)
	router.POST("/anotacoes", service.CreateAnotacao)
	router.PUT("/anotacoes", service.UpdateAnotacao)
	router.DELETE("/anotacoes/:id", service.DeleteAnotacao)

}

func GetRouter() *gin.Engine {
	return router
}

