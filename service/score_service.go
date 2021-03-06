package service

import (
	"fmt"
	"github.com/daviresio/ajuda-desk-api/database"
	"github.com/daviresio/ajuda-desk-api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListScore(c *gin.Context) {
	var scores []model.Score

	errors := database.DB.Find(&scores).GetErrors();

	fmt.Println(len(errors))

	if len(errors) > 0 {
		resErr := NewInternalServerError(errors[0].Error())
		c.JSON(resErr.Status, resErr)
		return
	}

	for _, err := range errors {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, scores)

}

func FindScore(c *gin.Context)  {
	var score model.Score

	id := GetIdParam("id", c)

	if id == 0 {
		return
	}

	notFound := database.DB.First(&score, id).RecordNotFound()


	if notFound == true {
		resErr := NewNotFoundError(fmt.Sprintf("score with id %d not found", id))
		c.JSON(resErr.Status, resErr)
		return
	}

	c.JSON(http.StatusOK, score)

}


func CreateScore(c *gin.Context) {
	var score model.Score

	if err := CheckInvalidJson(c.ShouldBindJSON(&score), c); err != nil {
		return
	}

	errors := database.DB.Create(&score).GetErrors()

	if len(errors) > 0 {
		resErr := NewInternalServerError(errors[0].Error())
		c.JSON(resErr.Status, resErr)
		return
	}

	fmt.Println(len(errors))

	c.JSON(http.StatusOK, score)

}

func UpdateScore(c *gin.Context) {
	var score model.Score

	if err := CheckInvalidJson(c.ShouldBindJSON(&score), c); err != nil {
		return
	}

	var originalScore model.Score
	database.DB.First(&originalScore, score.Id)

	database.DB.Save(&score)

	c.JSON(http.StatusOK, score)
}

func DeleteScore(c *gin.Context) {

	id := GetIdParam("id", c)

	if id == 0 {
		return
	}

	score := model.Score{
		BaseModel: model.BaseModel{
			Id: id,
		},
	}

	database.DB.Delete(&score)

	c.Status(http.StatusOK)
}

