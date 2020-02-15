package service

import (
	"fmt"
	"github.com/daviresio/ajuda-desk-api/database"
	"github.com/daviresio/ajuda-desk-api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListTipoPlano(c *gin.Context) {
	var tipoPlanos []model.TipoPlano

	errors := database.DB.Find(&tipoPlanos).GetErrors();

	fmt.Println(len(errors))

	if len(errors) > 0 {
		resErr := NewInternalServerError(errors[0].Error())
		c.JSON(resErr.Status, resErr)
		return
	}

	for _, err := range errors {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, tipoPlanos)

}

func FindTipoPlano(c *gin.Context)  {
	var tipoPlano model.TipoPlano

	id := GetIdParam("id", c)

	if id == 0 {
		return
	}

	notFound := database.DB.First(&tipoPlano, id).RecordNotFound()


	if notFound == true {
		resErr := NewNotFoundError(fmt.Sprintf("tipoPlano with id %d not found", id))
		c.JSON(resErr.Status, resErr)
		return
	}

	c.JSON(http.StatusOK, tipoPlano)

}


func CreateTipoPlano(c *gin.Context) {
	var tipoPlano model.TipoPlano

	if err := CheckInvalidJson(c.ShouldBindJSON(&tipoPlano), c); err != nil {
		return
	}

	errors := database.DB.Create(&tipoPlano).GetErrors()

	if len(errors) > 0 {
		resErr := NewInternalServerError(errors[0].Error())
		c.JSON(resErr.Status, resErr)
		return
	}

	fmt.Println(len(errors))

	c.JSON(http.StatusOK, tipoPlano)

}

func UpdateTipoPlano(c *gin.Context) {
	var tipoPlano model.TipoPlano

	if err := CheckInvalidJson(c.ShouldBindJSON(&tipoPlano), c); err != nil {
		return
	}

	var originalTipoPlano model.TipoPlano
	database.DB.First(&originalTipoPlano, tipoPlano.Id)

	database.DB.Save(&tipoPlano)

	c.JSON(http.StatusOK, tipoPlano)
}

func DeleteTipoPlano(c *gin.Context) {

	id := GetIdParam("id", c)

	if id == 0 {
		return
	}

	tipoPlano := model.TipoPlano{
		BaseModel: model.BaseModel{
			Id: id,
		},
	}

	database.DB.Delete(&tipoPlano)

	c.Status(http.StatusOK)
}

