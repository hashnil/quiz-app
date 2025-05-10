package models

import (
	"time"

	"gorm.io/datatypes"
)

type CreateQuizRequest struct {
	ClassroomID uint      `json:"classroom_id"`
	TeacherID   uint      `json:"teacher_id"`
	Title       string    `json:"title"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Questions   []struct {
		Content       string         `json:"content"`
		Options       datatypes.JSON `json:"options"`
		CorrectOption string         `json:"correct_option"`
	} `json:"questions"`
}

type QuizResponse struct {
	ID        uint                `json:"id"`
	Title     string              `json:"title"`
	StartTime string              `json:"start_time"`
	EndTime   string              `json:"end_time"`
	Questions []QuizQuestionBrief `json:"questions"`
}

type QuizQuestionBrief struct {
	ID      uint              `json:"id"`
	Content string            `json:"content"`
	Options map[string]string `json:"options"`
}

type SubmitAnswerRequest struct {
	StudentID      uint   `json:"student_id"`
	QuestionID     uint   `json:"question_id"`
	SelectedOption string `json:"selected_option"`
	DurationMs     int    `json:"duration_ms"`
}

type SubmitAnswerResponse struct {
	ID             uint   `json:"id"`
	StudentID      uint   `json:"student_id"`
	QuestionID     uint   `json:"question_id"`
	SelectedOption string `json:"selected_option"`
	SubmittedAt    string `json:"submitted_at"`
	DurationMs     int    `json:"duration_ms"`
	Correct        bool   `json:"correct"`
}

type StudentResponse struct {
	QuestionID     uint   `json:"question_id"`
	SelectedOption string `json:"selected_option"`
	SubmittedAt    string `json:"submitted_at"`
	DurationMs     int    `json:"duration_ms"`
	Correct        bool   `json:"correct"`
}

type StudentReport struct {
	TotalAttempted int               `json:"total_attempted"`
	CorrectAnswers int               `json:"correct_answers"`
	Accuracy       float64           `json:"accuracy"`
	Responses      []StudentResponse `json:"responses"`
}
