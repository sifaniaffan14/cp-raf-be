package utils

import (
	"time"
	"math"
	"reflect"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Error       bool        `json:"error"`
	Code        interface{}      `json:"code"`
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

const (
	DefaultPerPage = 15
	DefaultPage    = 1
)

type Meta struct {
	CurrentPage int    `json:"current_page"`
	PerPage     int    `json:"per_page"`
	Total       int    `json:"total"`
	LastPage    int    `json:"last_page"`
}

type PaginationPayload struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}

func GetPaginationParams(c *gin.Context) (int, int) {
	var payload PaginationPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		return DefaultPage, DefaultPerPage
	}

	page := payload.Page
	perPage := payload.PerPage

	if page <= 0 {
		page = DefaultPage
	}
	if perPage <= 0 {
		perPage = DefaultPerPage
	}

	return page, perPage
}

func Paginate(c *gin.Context, items interface{}, page, perPage int) (interface{}, Meta) {
	switch reflect.TypeOf(items).Kind() {
	case reflect.Slice:
		return paginateArray(items, page, perPage)
	default:
		return nil, Meta{}
	}
}

func paginateArray(items interface{}, page, perPage int) (interface{}, Meta) {
	val := reflect.ValueOf(items)
	if val.Kind() != reflect.Slice {
		return nil, Meta{}
	}

	total := val.Len()
	lastPage := int(math.Ceil(float64(total) / float64(perPage)))
	offset := (page - 1) * perPage

	if offset > total {
		return []interface{}{}, Meta{
			CurrentPage: page,
			PerPage:     perPage,
			Total:       total,
			LastPage:    lastPage,
		}
	}

	end := offset + perPage
	if end > total {
		end = total
	}

	data := val.Slice(offset, end).Interface()
	meta := Meta{
		CurrentPage: page,
		PerPage:     perPage,
		Total:       total,
		LastPage:    lastPage,
	}
	return data, meta
}
