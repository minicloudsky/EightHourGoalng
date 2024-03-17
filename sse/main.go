package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestBody struct {
	Messages         []Message `json:"messages"`
	Stream           bool      `json:"stream"`
	Model            string    `json:"model"`
	Temperature      float64   `json:"temperature"`
	PresencePenalty  float64   `json:"presence_penalty"`
	FrequencyPenalty float64   `json:"frequency_penalty"`
	TopP             float64   `json:"top_p"`
}

func main() {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			// 如果是预检请求（OPTIONS方法），则直接返回200
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	// SSE路由
	r.POST("/v1/chat/completions", func(c *gin.Context) {
		c.Header("Content-Type", "text/plain")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")
		c.Header("Access-Control-Allow-Origin", "*")
		var requestBody RequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 打印绑定后的结构体
		log.Printf("Received JSON request: %+v\n", requestBody)
		var content string
		if len(requestBody.Messages) > 0 {
			content = strings.Replace(requestBody.Messages[len(requestBody.Messages)-1].Content, "?", "!", -1)
			content = strings.Replace(content, "？", "!", -1)
			content = strings.Replace(content, "吗", "", -1)
		}
		c.SSEvent("", content)
	})
	// 启动HTTP服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
