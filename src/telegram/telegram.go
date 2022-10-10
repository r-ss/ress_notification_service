package telegram

import (
	"fmt"
	"net/http"
	// "os"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"main/secrets"
)


// Binding from JSON
type InputMessage struct {
	Text string `form:"message" json:"message"  binding:"required"`
	Silent string `form:"silent" json:"silent",default=False`
	ParseMode string `form:"parse_mode" json:"parse_mode"`
}

func TelegramNotificationHandler(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	var inputMessage InputMessage

	

	secrets, err := secrets.ReadSecrets()
    if err != nil {
        panic(err)
    }

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
		"chat_id": "95622548",
		"text": inputMessage.Text,
		"disable_notification": inputMessage.Silent,
		"parse_mode": inputMessage.ParseMode,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Unable to Marshal json.": err.Error()})
		return
    }

	var URL = "https://api.telegram.org/bot" + secrets.Token + "/sendMessage"

	fmt.Println(URL)

	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(requestBody))

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Unable to read body": err.Error()})
		return
    }

	c.JSON(http.StatusOK, gin.H{"status": "ok", "body": string(body)})
}
