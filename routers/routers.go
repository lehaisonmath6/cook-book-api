package routers

import (
	"github.com/cook-books-api/controllers"
	"github.com/gin-gonic/gin"
)

// @title Swagger Cook-Books
// @version 1.0
// @description This is a Cook books
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1/cook-books-api
func Register(r *gin.Engine) {
	v1 := r.Group("/v1/cook-books-api")
	{
		v1.GET("/lecture/", controllers.CreateLecture)
	}

}
