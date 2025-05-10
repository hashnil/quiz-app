package db

import dbmodels "quiz-app/internal/integration/db/models"

type Client interface {
	CreateQuiz(quiz *dbmodels.Quiz) error
	CreateQuestions(questions []dbmodels.Question) error
	GetQuizzesByClassroomID(classroomID uint) ([]dbmodels.Quiz, error)
	GetQuestionByID(id uint) (dbmodels.Question, error)
	CreateResponse(resp *dbmodels.Response) error
	GetResponsesByStudentID(studentID uint) ([]dbmodels.Response, error)
}
