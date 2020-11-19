package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

type server struct {
	router *gin.Engine
}

func NewServer() *server {
	r := gin.New()
	s := &server{
		router: r,
	}
	s.router.Use(CORSMiddleware())
	s.Setup()
	return s
}

func (s *server) Setup() {
	setupLogger(s)
	s.setupRoutes()
}

func (s *server) setupRoutes() {
	setupNumbersRoutes(s)
}

func (s *server) Run() {
	port := os.Getenv("SERVER_PORT")
	address := fmt.Sprintf(":%s", port)
	s.router.Run(address)
}
