package rest

import (
	"fama/app/numbers/adapters/handlers"
	"fama/core"
	_ "fama/numbers/managers"
)

func loadNumbersHandler() (*handlers.NumbersHandler, error) {
	var handler *handlers.NumbersHandler
	invokeFun := func(h *handlers.NumbersHandler) {
		handler = h
	}
	err := core.Injector.Invoke(invokeFun)
	return handler, err
}
