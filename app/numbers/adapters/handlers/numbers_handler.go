package handlers

import (
	"fama/core"
	"fama/numbers/ports"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	err := core.Injector.Provide(newNumbersHandler)
	core.CheckInjection(err, "newNumbersHandler")
}

type NumbersHandler struct {
	manager ports.NumbersManager
}

func newNumbersHandler(manager ports.NumbersManager) *NumbersHandler {
	return &NumbersHandler{
		manager: manager,
	}
}

func (h *NumbersHandler) ToWords(c *gin.Context) {
	req, err := newNumberToWordsRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	words, e := h.manager.ToWords(req.Number, req.Lang)

	if e != nil {
		generateError(c, http.StatusUnprocessableEntity, e)
		return
	}

	response := newNumberToWordsResponse(StatusOK, req.Lang, words)
	c.JSON(http.StatusOK, response)
}
