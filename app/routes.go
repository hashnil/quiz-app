package app

import (
	"quiz-app/api/controller"
)

// registerRoutes sets up the routes for the application.
func registerRoutes(healthController *controller.HealthController,
	whiteboardController *controller.WhiteboardController, notebookController *controller.NotebookController) {

	// Health check route
	engine.GET("/health", healthController.HealthCheck)

	// API v1 routes grouped under `/api/v1/user`
	apiV1 := engine.Group("/api/v1")

	// TODO: Middleware for authentication and authorization can be added here
	apiV1.POST("/quizzes", whiteboardController.CreateQuiz)
	apiV1.GET("/quizzes", whiteboardController.ListQuizzes)
	apiV1.POST("/quizzes/:quizID/submit", notebookController.SubmitAnswer)
	apiV1.GET("/reports/student/:studentID", notebookController.StudentReport)
}
