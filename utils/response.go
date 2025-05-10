package utils

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Error       bool        `json:"error"`
	Code        interface{} `json:"code"`
	Message     interface{} `json:"message"`
	RequestedAt string      `json:"requestedAt"`
	RespondedAt string      `json:"respondedAt"`
	Data        interface{} `json:"data"`
}

func SendResponse(c *gin.Context, result interface{}, message interface{}, httpCode int) {
	timezone := c.GetHeader("X-APP-TIMEZONE")
	if timezone == "" {
		timezone = "Asia/Jakarta"
	}

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		loc = time.Local
	}

	requestedAt := c.Query("requestedAt")
	var parsedRequestedAt string
	if requestedAt != "" {
		if t, err := time.Parse(time.RFC3339, requestedAt); err == nil {
			parsedRequestedAt = t.In(loc).Format("2006-01-02 15:04:05")
		} else {
			parsedRequestedAt = time.Now().In(loc).Format("2006-01-02 15:04:05")
		}
	} else {
		parsedRequestedAt = time.Now().In(loc).Format("2006-01-02 15:04:05")
	}

	response := Response{
		Error:       false,
		Code:        "0",
		Message:     message,
		RequestedAt: parsedRequestedAt,
		RespondedAt: time.Now().In(loc).Format("2006-01-02 15:04:05"),
		Data:        result,
	}

	c.JSON(httpCode, response)
}

func SendError(c *gin.Context, errMsg interface{}, code interface{}, httpCode int) {
	timezone := c.GetHeader("X-APP-TIMEZONE")
	if timezone == "" {
		timezone = "Asia/Jakarta"
	}

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		loc = time.Local
	}

	requestedAt := c.Query("requestedAt")
	var parsedRequestedAt string
	if requestedAt != "" {
		if t, err := time.Parse(time.RFC3339, requestedAt); err == nil {
			parsedRequestedAt = t.In(loc).Format("2006-01-02 15:04:05")
		} else {
			parsedRequestedAt = time.Now().In(loc).Format("2006-01-02 15:04:05")
		}
	} else {
		parsedRequestedAt = time.Now().In(loc).Format("2006-01-02 15:04:05")
	}

	response := Response{
		Error:       true,
		Code:        code,
		Message:     errMsg,
		RequestedAt: parsedRequestedAt,
		RespondedAt: time.Now().In(loc).Format("2006-01-02 15:04:05"),
		Data:        nil,
	}

	c.JSON(httpCode, response)
}
