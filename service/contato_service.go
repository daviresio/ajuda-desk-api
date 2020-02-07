package service

import (
	"fmt"
	"github.com/daviresio/ajuda-desk-api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListContato(c *gin.Context) {
	var contatos []model.Contato

	errors := db.Find(&contatos).GetErrors();

	fmt.Println(len(errors))

	if len(errors) > 0 {
		resErr := NewInternalServerError(errors[0].Error())
		c.JSON(resErr.Status, resErr)
		return
	}

	for _, err := range errors {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, contatos)

}

func FindContato(c *gin.Context)  {
	var contato model.Contato

	id := GetIdParam("id", c)

	if id == 0 {
		return
	}

	notFound := db.First(&contato, id).RecordNotFound()


	if notFound == true {
		resErr := NewNotFoundError(fmt.Sprintf("contato with id %d not found", id))
		c.JSON(resErr.Status, resErr)
		return
	}

	c.JSON(http.StatusOK, contato)

}


func CreateContato(c *gin.Context) {
	var contato model.Contato

	if err := CheckInvalidJson(c.ShouldBindJSON(&contato), c); err != nil {
		return
	}

	errors := db.Create(&contato).GetErrors()

	if len(errors) > 0 {
		resErr := NewInternalServerError(errors[0].Error())
		c.JSON(resErr.Status, resErr)
		return
	}

	fmt.Println(len(errors))

	c.JSON(http.StatusOK, contato)

}

func UpdateContato(c *gin.Context) {
	var contato model.Contato

	if err := CheckInvalidJson(c.ShouldBindJSON(&contato), c); err != nil {
		return
	}

	var originalContato model.Contato
	db.First(&originalContato, contato.Id)

	db.Save(&contato)

	c.JSON(http.StatusOK, contato)
}

func DeleteContato(c *gin.Context) {

	id := GetIdParam("id", c)

	if id == 0 {
		return
	}

	contato := model.Contato{
		BaseModel: model.BaseModel{
			Id: id,
		},
	}

	db.Delete(&contato)

	c.Status(http.StatusOK)
}
