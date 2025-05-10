package controller

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"quiz-app/entity/models"
	"quiz-app/internal/integration/db"
	dbmodels "quiz-app/internal/integration/db/models"

	"github.com/gin-gonic/gin"
)

type NotebookController struct {
	dbClient db.Client // Database client interface
}

func NewNotebookController(dbClient db.Client) *NotebookController {
	return &NotebookController{dbClient: dbClient}
}

// SubmitAnswer handles submission of a student's quiz answer.
func (nc *NotebookController) SubmitAnswer(c *gin.Context) {
	var req models.SubmitAnswerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Failed to bind request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	question, err := nc.dbClient.GetQuestionByID(req.QuestionID)
	if err != nil {
		log.Printf("Question ID %d not found: %v", req.QuestionID, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	isCorrect := req.SelectedOption == question.CorrectOption
	response := dbmodels.Response{
		StudentID:      req.StudentID,
		QuestionID:     req.QuestionID,
		SelectedOption: req.SelectedOption,
		SubmittedAt:    time.Now(),
		DurationMs:     req.DurationMs,
		IsCorrect:      isCorrect,
	}

	if err := nc.dbClient.CreateResponse(&response); err != nil {
		log.Printf("Failed to store response: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save response"})
		return
	}

	c.JSON(http.StatusOK, models.SubmitAnswerResponse{
		ID:             response.ID,
		StudentID:      response.StudentID,
		QuestionID:     response.QuestionID,
		SelectedOption: response.SelectedOption,
		SubmittedAt:    response.SubmittedAt.Format(time.RFC3339),
		DurationMs:     response.DurationMs,
		Correct:        response.IsCorrect,
	})
}

// StudentReport generates a performance report for a given student.
func (nc *NotebookController) StudentReport(c *gin.Context) {
	studentIDStr := c.Param("studentID")
	studentIDUint64, err := strconv.ParseUint(studentIDStr, 10, 64)
	if err != nil {
		log.Printf("Invalid student ID: %s", studentIDStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}
	studentID := uint(studentIDUint64)

	responses, err := nc.dbClient.GetResponsesByStudentID(studentID)
	if err != nil {
		log.Printf("Failed to fetch responses for student %d: %v", studentID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate report"})
		return
	}

	total := len(responses)
	if total == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No responses found for student"})
		return
	}

	var correct int
	var responseDTOs []models.StudentResponse

	for _, r := range responses {
		if r.IsCorrect {
			correct++
		}
		responseDTOs = append(responseDTOs, models.StudentResponse{
			QuestionID:     r.QuestionID,
			SelectedOption: r.SelectedOption,
			SubmittedAt:    r.SubmittedAt.Format(time.RFC3339),
			DurationMs:     r.DurationMs,
			Correct:        r.IsCorrect,
		})
	}

	report := models.StudentReport{
		TotalAttempted: total,
		CorrectAnswers: correct,
		Accuracy:       float64(correct) / float64(total),
		Responses:      responseDTOs,
	}

	c.JSON(http.StatusOK, report)
}
