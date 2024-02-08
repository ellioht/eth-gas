package api

import (
	"context"
	"github.com/ellioht/eth-gas/config"
	database "github.com/ellioht/eth-gas/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	db     *database.Database
	router *gin.Engine
}

func NewServer(cfg config.API, db *database.Database) *Server {
	router := gin.Default()

	server := &Server{
		db:     db,
		router: router,
	}

	server.setupRoutes()

	return server
}

func (s *Server) Start(port string) error {
	return s.router.Run(":" + port)
}

func (s *Server) setupRoutes() {
	s.router.GET("/gasprices", s.handleGetGasPrices)
}

func (s *Server) handleGetGasPrices(c *gin.Context) {
	start := time.Now().Add(-24 * time.Hour) // last 24 hours
	end := time.Now()

	records, err := s.db.RetrieveGasPrices(context.Background(), start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve gas prices"})
		return
	}

	c.JSON(http.StatusOK, records)
}
