package service

import (
	"fmt"
	"github.com/daviresio/ajuda-desk-api/database"
	"github.com/daviresio/ajuda-desk-api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListAnotacao(c *gin.Context) {
	var anotacoes []model.Anotacao

	c.JSON(http.StatusOK, anotacoes)

}

func FindAnotacao(c *gin.Context)  {
	var anotacao model.Anotacao

	id := GetIdParam("id", c)

	if id == 0 {
		return
	}

	notFound := database.DB.First(&anotacao, id).RecordNotFound()


	if notFound == true {
		resErr := NewNotFoundError(fmt.Sprintf("anotacao with id %d not found", id))
		c.JSON(resErr.Status, resErr)
		return
	}

	c.JSON(http.StatusOK, anotacao)

}


func CreateAnotacao(c *gin.Context) {
	var anotacao model.Anotacao

	if err := CheckInvalidJson(c.ShouldBindJSON(&anotacao), c); err != nil {
		return
	}

	database.DB.Create(&anotacao)

	c.JSON(http.StatusOK, anotacao)

}

func UpdateAnotacao(c *gin.Context) {
	var anotacao model.Anotacao

	if err := CheckInvalidJson(c.ShouldBindJSON(&anotacao), c); err != nil {
		return
	}

	var originalAnotacao model.Anotacao
	database.DB.First(&originalAnotacao, anotacao.Id)

	database.DB.Save(&anotacao)

	c.JSON(http.StatusOK, anotacao)
}

func DeleteAnotacao(c *gin.Context) {

	id := GetIdParam("id", c)

	if id == 0 {
		return
	}

	anotacao := model.Anotacao{
		BaseModel: model.BaseModel{
			Id: id,
		},
	}

	database.DB.Delete(&anotacao)

	c.Status(http.StatusOK)
}