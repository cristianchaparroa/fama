package rest

func setupNumbersRoutes(s *server) {
	handler, err := loadNumbersHandler()
	checkError(err)

	s.router.GET("/numbers/:number/words", handler.ToWords)
}
