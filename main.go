package main

import (
	"log"

	_ "github.com/cook-books-api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization,Content-Length, Access-Control-Allow-Origin, content-type, Content-Type, Accept, accept, Accept-Encoding,Accept-Language, sessionkey, token, session")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func main() {
	r := gin.New()
	r.Use(CORSMiddleware())

	// fs := http.FileServer(http.Dir("./static"))
	r.Static("/v1/cook-books-api/media", "./media")
	r.Static("/v1/cook-books-api/documents", "./documents")
	r.Static("/v1/cook-books-api/images", "./images")
	r.GET("/v1/cook-books-api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("v1")))
	log.Println("[INFO] link swagger", "http://localhost"+":8012"+"/v1/cook-books-api/swagger/index.html")
	r.Run(":8012")
	// r.GET("/v1/cook-books-api/")

}
