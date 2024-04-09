package main

import (
	"log"
	"newRpc/cache"
	"newRpc/control"
	"newRpc/dao"
	"newRpc/router"
)

func main() {
	_, err := dao.InitDB() //初始化数据库
	if err != nil {
		log.Println(err)
		return
	}

	err = cache.InitRedis() //初始化redis
	if err != nil {
		log.Println(err)
		return
	}

	control.InitClient() //初始化kitex

	r := router.InitRouter()
	r.Run(":8080")
}
