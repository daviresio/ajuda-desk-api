package service

import (
	"fmt"
	"github.com/daviresio/ajuda-desk-api/database"
	"github.com/daviresio/ajuda-desk-api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListTipoEmpresa(c *gin.Context) {
	var tipoEmpresas []model.TipoEmpresa

	errors := database.DB.Find(&tipoEmpresas).GetErrors();

	fmt.Println(len(errors))

	if len(errors) > 0 {
		resErr := NewInternalServerError(errors[0].Error())
		c.JSON(resErr.Status, resErr)
		return
	}

	for _, err := range errors {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, tipoEmpresas)

}

func FindTipoEmpresa(c *gin.Context)  {
	var tipoEmpresa model.TipoEmpresa

	id := GetIdParam("id", c)

	if id == 0 {
		return
	}

	notFound := database.DB.First(&tipoEmpresa, id).RecordNotFound()


	if notFound == true {
		resErr := NewNotFoundError(fmt.Sprintf("tipoEmpresa with id %d not found", id))
		c.JSON(resErr.Status, resErr)
		return
	}

	c.JSON(http.StatusOK, tipoEmpresa)

}


func CreateTipoEmpresa(c *gin.Context) {
	var tipoEmpresa model.TipoEmpresa

	if err := CheckInvalidJson(c.ShouldBindJSON(&tipoEmpresa), c); err != nil {
		return
	}

	errors := database.DB.Create(&tipoEmpresa).GetErrors()

	if len(errors) > 0 {
		resErr := NewInternalServerError(errors[0].Error())
		c.JSON(resErr.Status, resErr)
		return
	}

	fmt.Println(len(errors))

	c.JSON(http.StatusOK, tipoEmpresa)

}

func UpdateTipoEmpresa(c *gin.Context) {
	var tipoEmpresa model.TipoEmpresa

	if err := CheckInvalidJson(c.ShouldBindJSON(&tipoEmpresa), c); err != nil {
		return
	}

	var originalTipoEmpresa model.TipoEmpresa
	database.DB.First(&originalTipoEmpresa, tipoEmpresa.Id)

	database.DB.Save(&tipoEmpresa)

	c.JSON(http.StatusOK, tipoEmpresa)
}

func DeleteTipoEmpresa(c *gin.Context) {

	id := GetIdParam("id", c)

	if id == 0 {
		return
	}

	tipoEmpresa := model.TipoEmpresa{
		BaseModel: model.BaseModel{
			Id: id,
		},
	}

	database.DB.Delete(&tipoEmpresa)

	c.Status(http.StatusOK)
}

