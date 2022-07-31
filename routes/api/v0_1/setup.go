package v0_1

import (
	"github.com/berbreik/UserManagementService/routes/api/v0_1/app"
	"github.com/berbreik/UserManagementService/routes/api/v0_1/authRequest"
	"github.com/berbreik/UserManagementService/routes/api/v0_1/user"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup) {
	// Users
	router.GET("/users", user.GetList)
	router.GET("/user", user.Get)
	router.GET("/user/:id", user.GetById)

	router.POST("/user", user.PostOne)
	router.POST("/users", user.PostAll)

	router.POST("/user/login", user.Login)

	router.PUT("/user/update/:id", user.UpdateOneById)
	router.PUT("/user/update", user.UpdateOne)
	router.PUT("/users/update", user.UpdateAll)

	router.PUT("/user", user.ReplaceOne)
	router.PUT("/users", user.ReplaceAll)

	router.DELETE("/user/:id", user.DeleteOne)
	router.DELETE("/users", user.DeleteAll)

	// Apps
	router.GET("/apps", app.GetList)
	router.GET("/app", app.Get)
	router.GET("/app/:id", app.GetById)

	router.POST("/app", app.PostOne)
	router.POST("/apps", app.PostAll)

	router.PUT("/app/update/:id", app.UpdateOneById)
	router.PUT("/app/update", app.UpdateOne)
	router.PUT("/apps/update", app.UpdateAll)

	router.PUT("/app", app.ReplaceOne)
	router.PUT("/apps", app.ReplaceAll)

	router.DELETE("/app/:id", user.DeleteOne)
	router.DELETE("/apps", user.DeleteAll)

	// IsAuth
	router.POST("/isAuth", authRequest.IsAuth)
}
