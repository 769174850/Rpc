package router

import (
	"github.com/gin-gonic/gin"
	"newRpc/control"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	//用户的接口
	r.POST("/register", control.Register) //注册一个新用户
	r.POST("/login", control.Login)       //用户登录

	//短链接口
	r.POST("shortUrl/generate", control.GenerateShortLink) //生成短链
	r.POST("shortUrl/delete", control.DeleteShortLink)     //删除短链
	r.POST("shortUrl/update", control.UpdateShortLink)     //更新短链
	r.POST("shortUrl/get/:id", control.GetUserShortLink)   //获取用户短链
	r.POST("shortUrl/rank", control.GetShortLinkRank)      //获取短链点击量排行榜

	return r
}
