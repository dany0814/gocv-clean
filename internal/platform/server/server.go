package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
}

func New(host string, port uint) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),
	}

	srv.Routes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on...", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) Routes() {
	s.engine.GET("/health", CheckHandler())
}

func CheckHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "everything is ok!")
	}
}
