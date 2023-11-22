package controllers

import (
	"github.com/gin-gonic/gin"
	"server/internal/core/helpers"
)

func (u *HTTPHandler) PingHandler(c *gin.Context) {
	// healthcheck
	helpers.JSON(c, "pong", 200, nil, nil)
}
