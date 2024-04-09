package control

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"newRpc/dao"
	"newRpc/model"
	"newRpc/util"
	"time"
)

var jwtSecret = []byte("Password-Secret-key")

func generateToken(userID int64) (string, string, error) {
	tokenExpire := time.Now().Add(2 * time.Hour).Unix()         //token过期时间
	refreshTokenExpire := time.Now().Add(12 * time.Hour).Unix() //refreshToken过期时间

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ //创建jwt令牌
		"user_id": userID,
		"exp":     tokenExpire,
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ //刷新jwt令牌
		"user_id": userID,
		"exp":     refreshTokenExpire,
	})

	tokenStr, err := token.SignedString(jwtSecret) //签署令牌
	if err != nil {
		return "", "", err
	}

	refreshTokenStr, err := refreshToken.SignedString(jwtSecret) //签署刷新令牌
	if err != nil {
		return "", "", err
	}

	return tokenStr, refreshTokenStr, nil
}

func Login(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := c.ShouldBindJSON(&loginRequest) //获取用户名和密码
	if err != nil {
		util.RespInternalErr(c) //获取失败返回错误
		return
	}

	userID, err := dao.VerifyUserAndGetID(loginRequest.Username, loginRequest.Password) //锁定唯一用户排除用户同名
	if err != nil {
		util.RespErr(c)
		return
	}
	if userID == 0 {
		util.RespErr(c)
		return
	}

	token, refreshToken, err := generateToken(userID) //生成jwt令牌
	if err != nil {
		util.RespInternalErr(c)
		return
	}

	util.RespOKWithData(c, gin.H{
		"refresh_token": refreshToken,
		"token":         token,
	}) //返回生成的令牌
}

func Register(c *gin.Context) {
	var u model.User
	err := c.ShouldBindJSON(&u) //获取注册的用户信息
	if err != nil {
		util.RespParamErr(c) //获取失败返回参数错误
		return
	}

	users, err := dao.GetUser() //检查用户是否被注册
	if err != nil {
		log.Println(err)
		util.RespInternalErr(c) //检索失败返回服务错误
		return
	}

	for _, user := range users {
		if user.Username == u.Username {
			util.RespAlreadyReported(c) //返回用户名已被注册返回用户名已被注册的错误
			return
		}
	}

	err = dao.AddUser(u) //添加用户
	if err != nil {
		log.Println(err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}
