package main

import (
	"log"
	api "newRpc/kitex_gen/api/shortlinkservice"
)

func main() {
	svr := api.NewServer(new(ShortLinkServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
