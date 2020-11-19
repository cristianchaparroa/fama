package handlers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	contentType            = "Definition-Type"
	ProblemJSONContentType = "application/problem+json"
)

func GenerateError(c *gin.Context, e *Error) {
	log.WithFields(map[string]interface{}{"module": "errors"}).Error(e.String())

	if e.Status == http.StatusInternalServerError {
		AbortWithError(c, e.Status, e)
		return
	}

	c.Header(contentType, ProblemJSONContentType)
	c.JSON(e.Status, e)
}

func AbortWithError(c *gin.Context, status int, e *Error) {
	c.AbortWithStatusJSON(status, e)
}
