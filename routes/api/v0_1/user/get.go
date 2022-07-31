package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetList(c *gin.Context) {
	u := models.Users{}
	r, er := modelApi.FetchAll("users", c, u.DbFetchAll)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, r)
}

func Get(c *gin.Context) {
	u := models.User{}
	r, er := modelApi.FetchOne("users", c, u.DbFetchOne)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, r)
}

func GetById(c *gin.Context) {
	u := models.User{}
	r, er := modelApi.FetchById("users", c, u.DbFetchOne)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, r)
}
