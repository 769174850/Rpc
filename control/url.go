package control

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
	"log"
	"newRpc/kitex_gen/api"
	"newRpc/kitex_gen/api/shortlinkservice"
	"newRpc/util"
)

var NewClient shortlinkservice.Client

func InitClient() error {
	var err error
	NewClient, err = shortlinkservice.NewClient("hello", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func GenerateShortLink(c *gin.Context) {
	var req api.GenerateShortLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.RespParamErr(c)
		return
	}

	//调用Kitex服务
	resp, err := NewClient.GenerateShortLink(context.Background(), &req)
	if err != nil {
		util.RespInternalErr(c)
		return
	}

	util.RespOKWithData(c, resp)
}

func DeleteShortLink(c *gin.Context) {
	var req api.DeleteShortLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.RespErr(c)
		return
	}

	//调用Kitex服务
	resp, err := NewClient.DeleteShortLink(context.Background(), &req)
	if err != nil {
		util.RespInternalErr(c)
		return
	}

	util.RespOKWithData(c, resp)

}

func UpdateShortLink(c *gin.Context) {
	var req api.UpdateShortLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.RespErr(c)
		return
	}

	// 调用Kitex服务
	resp, err := NewClient.UpdateShortLink(context.Background(), &req)
	if err != nil {
		util.RespInternalErr(c)
		return
	}

	util.RespOKWithData(c, resp)
}

func GetUserShortLink(c *gin.Context) {
	var req api.GetUserShortLinksRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.RespErr(c)
		return
	}

	// 调用Kitex服务
	resp, err := NewClient.GetUserShortLinks(context.Background(), &req)
	if err != nil {
		util.RespInternalErr(c)
		return
	}

	util.RespOKWithData(c, resp)
}

func GetShortLinkRank(c *gin.Context) {
	var req api.GetShortLinkRankingsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.RespErr(c)
		return
	}

	// 调用Kitex服务
	resp, err := NewClient.GetShortLinkRankings(context.Background(), &req)
	if err != nil {
		util.RespInternalErr(c)
		return
	}

	util.RespOKWithData(c, resp)
}
