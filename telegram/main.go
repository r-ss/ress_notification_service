package main

import (
	"net/http"
	"os"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
)

// Binding from JSON
type InputMessage struct {
	Text string `form:"message" json:"message"  binding:"required"`
	Silent string `form:"silent" json:"silent",default=False`
	ParseMode string `form:"parse_mode" json:"parse_mode"`
}

func telegramNotificationHandler(c *gin.Context) {

	var inputMessage InputMessage

	// querystring_message := c.DefaultQuery("message", "No_message_querystring")

	// try bind InputMessage from querystring
	if err := c.ShouldBind(&inputMessage); err != nil {

		// Then try bing InputMessage from input json
		if err := c.ShouldBindJSON(&inputMessage); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}



	requestBody, err := json.Marshal(map[string]string{
		"chat_id": os.Getenv("TG_CHAT_ID"),
		"text": inputMessage.Text,
		"disable_notification": inputMessage.Silent,
		"parse_mode": inputMessage.ParseMode,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Unable to Marshal json.": err.Error()})
		return
    }

	resp, err := http.Post(os.Getenv("TG_API_URL"), "application/json", bytes.NewBuffer(requestBody))

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Unable to read body": err.Error()})
		return
    }

	c.JSON(http.StatusOK, gin.H{"status": "ok", "body": string(body)})
}

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
	r.GET("/telegram", telegramNotificationHandler)
	r.POST("/telegram", telegramNotificationHandler)

	return r
}

func main() {
	addr := ":" + os.Getenv("PORT")
	gateway.ListenAndServe(addr, routerEngine())
}