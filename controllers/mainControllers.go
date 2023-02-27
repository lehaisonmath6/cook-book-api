package controllers

import (
	"fmt"
	"net/http"

	"github.com/cook-books-api/models"
	"github.com/gin-gonic/gin"
)

// @Title CreateLecture
// @Summary tạo mới bài giảng
// @Description tạo mới bài giảng
// @Tags lecture
// @Accept  json
// @Produce  json
// @Param	session	header	string	true	"session key"
// @Param	body	body  models.Lecture	true	"int64"
// @Success 200	{object} models.Lecture
// @Failure 403 {pubKey} not found
// @router /lecture/create [post]
func CreateLecture(c *gin.Context) {
	var lecture models.Lecture
	fmt.Println(lecture)
}

func DownloadLicense(ctx *gin.Context) {
	content := "Download file here happliy"
	fileName := "hello.txt"
	ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	ctx.Header("Content-Type", "application/text/plain")
	ctx.Header("Accept-Length", fmt.Sprintf("%d", len(content)))
	ctx.Writer.Write([]byte(content))
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Download file successfully",
	})
}
