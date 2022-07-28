package router

import "github.com/gin-gonic/gin"

func Run(host string) {
	router := gin.Default()
	router.Run(host)
}
