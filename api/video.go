package api

import (
	"singo/serializer"
	"github.com/gin-gonic/gin"
)

//	CreateVideo 视频创建接口
func CreateVideo(c *gin.Context) {
	c.JSON (200, serializer.Response {
		Code: 0,
		Msg: "创建成功",
	})
}