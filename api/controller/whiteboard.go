package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"quiz-app/entity/models"
	"quiz-app/internal/integration/db"
	dbmodels "quiz-app/internal/integration/db/models"

	"github.com/gin-gonic/gin"
)

type WhiteboardController struct {
	dbClient db.Client // Database client for whiteboard operations
}

func NewWhiteboardController(dbClient db.Client) *WhiteboardController {
	return &WhiteboardController{dbClient: dbClient}
}

// CreateQuiz handles the creation of a quiz with associated questions
func (wc *WhiteboardController) CreateQuiz(c *gin.Context) {
	var req models.CreateQuizRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Invalid quiz creation request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	quiz := dbmodels.Quiz{
		ClassroomID: req.ClassroomID,
		TeacherID:   req.TeacherID,
		Title:       req.Title,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
	}

	if err := wc.dbClient.CreateQuiz(&quiz); err != nil {
		log.Printf("Failed to create quiz: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create quiz"})
		return
	}

	var questions []dbmodels.Question
	for _, q := range req.Questions {
		questions = append(questions, dbmodels.Question{
			QuizID:        quiz.ID,
			Content:       q.Content,
			Options:       q.Options,
			CorrectOption: q.CorrectOption,
		})
	}

	if err := wc.dbClient.CreateQuestions(questions); err != nil {
		log.Printf("Failed to create questions for quiz %d: %v", quiz.ID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create quiz questions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"quiz_id": quiz.ID, "message": "Quiz created successfully"})
}

// ListQuizzes fetches all quizzes for a given classroom
func (wc *WhiteboardController) ListQuizzes(c *gin.Context) {
	classroomIDStr := c.Query("classroom_id")
	classroomID, err := strconv.ParseUint(classroomIDStr, 10, 64)
	if err != nil {
		log.Printf("Invalid classroom ID: %s", classroomIDStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid classroom ID"})
		return
	}

	quizzes, err := wc.dbClient.GetQuizzesByClassroomID(uint(classroomID))
	if err != nil {
		log.Printf("Failed to fetch quizzes for classroom %d: %v", classroomID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch quizzes"})
		return
	}

	var response []models.QuizResponse
	for _, quiz := range quizzes {
		var questions []models.QuizQuestionBrief
		for _, q := range quiz.Questions {
			var opts map[string]string
			if err := json.Unmarshal(q.Options, &opts); err != nil {
				log.Printf("Failed to unmarshal options for question %d: %v", q.ID, err)
				continue
			}
			questions = append(questions, models.QuizQuestionBrief{
				ID:      q.ID,
				Content: q.Content,
				Options: opts,
			})
		}

		response = append(response, models.QuizResponse{
			ID:        quiz.ID,
			Title:     quiz.Title,
			StartTime: quiz.StartTime.Format(time.RFC3339),
			EndTime:   quiz.EndTime.Format(time.RFC3339),
			Questions: questions,
		})
	}

	c.JSON(http.StatusOK, response)
}
