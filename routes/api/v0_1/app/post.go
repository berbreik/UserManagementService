package app

import (
	"encoding/json"
	modelApi "github.com/berbreik/UserManagementService/models/api"
	models "github.com/berbreik/UserManagementService/models/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostOne(c *gin.Context) {
	a := models.App{}
	//c.Bind(&u)
	json.NewDecoder(c.Request.Body).Decode(&a)
	aa, er := modelApi.InsertOne("apps", a.DbInsertOne)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, aa)
}

func PostAll(c *gin.Context) {
	a := models.Apps{}
	json.NewDecoder(c.Request.Body).Decode(&a)
	aa, er := modelApi.InsertAll("apps", a.DbInsertAll)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, aa)
}
