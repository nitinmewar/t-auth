package models

import "github.com/gin-gonic/gin"

type SlackRequest struct {
	Context    *gin.Context
	StatusCode int
	ErrorType  string
	Message    string
}
