package routes

import (
	"github.com/berbreik/UserManagementService/routes/api"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup) {

	r := router.Group("/api")
	api.Setup(r)
}
