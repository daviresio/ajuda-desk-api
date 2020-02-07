package service

import (
	"fmt"
	"github.com/daviresio/ajuda-desk-api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListEmpresa(c *gin.Context) {
	var empresas []model.Empresa

	query := fmt.Sprintf("%%%s%%", c.Request.URL.Query().Get("nome"))

	fmt.Println(query)

	errors := db.Preload("Contatos").Find(&empresas).GetErrors();

	fmt.Println(len(errors))

	if len(errors) > 0 {
		resErr := NewInternalServerError(errors[0].Error())
		c.JSON(resErr.Status, resErr)
		return
	}

	for _, err := range errors {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, empresas)

}

func FindEmpresa(c *gin.Context)  {
	var empresa model.Empresa

	id := GetIdParam("id", c)

	if id == 0 {
		return
	}

	notFound := db.First(&empresa, id).RecordNotFound()


	if notFound == true {
		resErr := NewNotFoundError(fmt.Sprintf("empresa with id %d not found", id))
		c.JSON(resErr.Status, resErr)
		return
	}

	c.JSON(http.StatusOK, empresa)

}


func CreateEmpresa(c *gin.Context) {
	var empresa model.Empresa

	if err := CheckInvalidJson(c.ShouldBindJSON(&empresa), c); err != nil {
		return
	}

	errors := db.Create(&empresa).GetErrors()

	if len(errors) > 0 {
		resErr := NewInternalServerError(errors[0].Error())
		c.JSON(resErr.Status, resErr)
		return
	}

	fmt.Println(len(errors))

	c.JSON(http.StatusOK, empresa)

}

func UpdateEmpresa(c *gin.Context) {
	var empresa model.Empresa

	if err := CheckInvalidJson(c.ShouldBindJSON(&empresa), c); err != nil {
		return
	}

	var originalEmpresa model.Empresa
	db.First(&originalEmpresa, empresa.Id)

	db.Save(&empresa)

	c.JSON(http.StatusOK, empresa)
}

func DeleteEmpresa(c *gin.Context) {

	id := GetIdParam("id", c)

	if id == 0 {
		return
	}

	empresa := model.Empresa{
		BaseModel: model.BaseModel{
			Id: id,
		},
	}

	db.Delete(&empresa)

	c.Status(http.StatusOK)
}
