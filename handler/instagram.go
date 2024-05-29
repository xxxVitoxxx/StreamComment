package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Instagram struct{}

func (ig Instagram) Router(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.GET("/instagram/webhook", ig.subscription)
}

// subscription create an endpoint that accepts and processes webhooks.
func (ig Instagram) subscription(c *gin.Context) {
	mode, _ := c.GetQuery("hub.mode")
	token, _ := c.GetQuery("hub.verify_token")
	if mode != "subscribe" || token != viper.GetString("verify_token") {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	challenge, _ := c.GetQuery("hub.challenge")
	c.String(http.StatusOK, challenge)
}
