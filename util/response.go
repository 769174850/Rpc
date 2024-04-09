package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type respTemplate struct {
	Status int         `json:"status"`
	Info   string      `json:"info"`
	Data   interface{} `json:"data"`
} //声明返回值的结构体

var InternalError = respTemplate{
	Status: 500,
	Info:   "internal error",
	Data:   "",
} //定义内部错误的返回值

var Unauthorized = respTemplate{
	Status: 401,
	Info:   "invalid username or password",
	Data:   "",
} //定义账号不存在错误的返回值

var ParamError = respTemplate{
	Status: 300,
	Info:   "params error",
	Data:   "",
} //定义获取信息失败错误的返回值

var OK = respTemplate{
	Status: 200,
	Info:   "success",
	Data:   "",
} //定义成功的返回值

var AlreadyReported = respTemplate{
	Status: 400,
	Info:   "Username or password already exists",
	Data:   "",
} //定义账号已存在的返回值

func RespInternalErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, InternalError)
} //返回内部错误500

func RespErr(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Unauthorized)
} //返回账号不存在错误401

func RespOKWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   "success",
		"data":   data,
	})
} //带有数据的返回200

func RespParamErr(c *gin.Context) {
	c.JSON(http.StatusBadRequest, ParamError)
} //返回获取信息失败400

func RespOK(c *gin.Context) {
	c.JSON(http.StatusOK, OK)
} //返回成功的返回值200

func RespAlreadyReported(c *gin.Context) {
	c.JSON(http.StatusBadRequest, AlreadyReported)
} //返回账号已存在错误400
