package handlers

import (
	"fama/core/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	defaultLang = "en"
)

type numberToWordsRequest struct {
	Number int
	Lang   string
}

func newNumberToWordsRequest(c *gin.Context) (*numberToWordsRequest, *handlers.Error) {
	lang := c.Query("lang")
	if lang == "" {
		lang = defaultLang
	}

	n := c.Param("number")
	number, err := strconv.Atoi(n)

	if err != nil {

		e := handlers.NewErrorWithStatus(InvalidRequestTitle, http.StatusBadRequest)
		param := handlers.InvalidParam{
			Name: "number",
			Code: InvalidNumberError,
		}

		e.InvalidParams = []handlers.InvalidParam{param}
		return nil, e
	}

	return &numberToWordsRequest{
		Number: number,
		Lang:   lang,
	}, nil
}
