package router

import (
	"github.com/gin-gonic/gin"
	"github.com/palp1tate/MultiLevelCommentDemo/api"
)

func InitRouter() *gin.Engine {
	Router := gin.Default()

	Router.POST("/add_moment", api.AddMoment)
	Router.POST("/add_comment", api.AddComment)
	Router.GET("/get_comments", api.GetComments)

	return Router
}
