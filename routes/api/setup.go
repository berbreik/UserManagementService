package api

import (
	"github.com/berbreik/UserManagementService/routes/api/v0_1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup(router *gin.RouterGroup) {
	r := router.Group("/v0_1")
	v0_1.Setup(r)

	router.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.User{Name: "vivek"})
	})
}
