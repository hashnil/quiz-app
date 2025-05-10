package app

import (
	"fmt"
	"quiz-app/api/controller"
	"quiz-app/infrastructure/factory"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Service struct {
	Engine *gin.Engine // Gin Engine for handling HTTP requests
}

var (
	engineLock sync.RWMutex // Mutex for ensuring thread-safe engine initialization
	engine     *gin.Engine  // Singleton instance of the Gin engine
)

// NewService initializes the Service struct, sets up the Gin engine, and registers routes.
func NewService() (*Service, error) {
	// Initialize PostgreSQL client for database operations.
	dbClient, err := factory.InitDBClient("postgresql")
	if err != nil {
		return &Service{}, fmt.Errorf("failed to initialize postgres client: %w", err)
	}

	healthController := controller.NewHealthController()
	whiteboardController := controller.NewWhiteboardController(dbClient)
	notebookController := controller.NewNotebookController(dbClient)

	// Ensure the engine is initialized only once
	if engine == nil {
		initializeEngine()
	}

	// Register application routes
	registerRoutes(healthController, whiteboardController, notebookController)

	return &Service{
		Engine: engine,
	}, nil
}

// initializeEngine initializes the Gin engine in a thread-safe manner.
func initializeEngine() {
	engineLock.Lock()
	defer engineLock.Unlock()

	if engine == nil {
		engine = gin.New()
		engine.Use(gin.Recovery(), gin.Logger())
		gin.SetMode(gin.ReleaseMode)         // Use ReleaseMode for production
		engine.HandleMethodNotAllowed = true // Enable 405 Method Not Allowed responses
	}
}

// Run starts the HTTP server on port 8080.
func (s *Service) Run() error {
	return s.Engine.Run(":" + viper.GetString("service.port"))
}
