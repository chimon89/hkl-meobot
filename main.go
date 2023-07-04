package main

import (
	"github.com/gin-gonic/gin"
	"khl-meobot/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)

	r.Run(":8190")
}
