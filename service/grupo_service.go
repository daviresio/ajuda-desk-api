package service

import (
	"fmt"
	"github.com/daviresio/ajuda-desk-api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListGrupo(c *gin.Context) {
	var grupos []model.Grupo

	db.Find(&grupos);

	c.JSON(http.StatusOK, grupos)

}

func FindGrupo(c *gin.Context)  {
	var grupo model.Grupo

	id := GetIdParam("id", c)

	if id == 0 {
		return
	}

	notFound := db.First(&grupo, id).RecordNotFound()


	if notFound == true {
		resErr := NewNotFoundError(fmt.Sprintf("grupo with id %d not found", id))
		c.JSON(resErr.Status, resErr)
		return
	}

	c.JSON(http.StatusOK, grupo)

}


func CreateGrupo(c *gin.Context) {
	var grupo model.Grupo

	if err := CheckInvalidJson(c.ShouldBindJSON(&grupo), c); err != nil {
		return
	}

	db.Create(&grupo)

	c.JSON(http.StatusOK, grupo)

}

func UpdateGrupo(c *gin.Context) {
	var grupo model.Grupo

	if err := CheckInvalidJson(c.ShouldBindJSON(&grupo), c); err != nil {
		return
	}

	var originalGrupo model.Grupo
	db.First(&originalGrupo, grupo.Id)

	db.Save(&grupo)

	c.JSON(http.StatusOK, grupo)
}

func DeleteGrupo(c *gin.Context) {

	id := GetIdParam("id", c)

	if id == 0 {
		return
	}

	grupo := model.Grupo{
		BaseModel: model.BaseModel{
			Id: id,
		},
	}

	db.Delete(&grupo)

	c.Status(http.StatusOK)
}