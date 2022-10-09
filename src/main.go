package main

import (
	"net/http"
	// "os"
	// "encoding/json"
	// "bytes"
	// "io/ioutil"
	"github.com/gin-gonic/gin"
	"main/telegram"
)


func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"text": "there is no root",
	})
}

func infoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"resource": "ress-notification-service",
	})
}

func routerEngine() *gin.Engine {
	// set server mode
	gin.SetMode(gin.DebugMode)

	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", rootHandler)
	r.GET("/info", infoHandler)
	r.GET("/telegram", telegram.TelegramNotificationHandler)
	r.POST("/telegram", telegram.TelegramNotificationHandler)

	return r
}

// func main() {
//   r := gin.Default()
//   r.GET("/ping", func(c *gin.Context) {
//     c.JSON(http.StatusOK, gin.H{
//       "message": "pong",
//     })
//   })
//   r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
// }

func main() {
	// addr := ":" + os.Getenv("PORT")
	// gateway.ListenAndServe(addr, routerEngine())
	// routerEngine()
	r := routerEngine()
	r.Run()
}