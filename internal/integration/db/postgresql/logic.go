package postgresql

import dbmodels "quiz-app/internal/integration/db/models"

func (db *PostgresSQLClient) CreateQuiz(quiz *dbmodels.Quiz) error {
	return db.client.Create(quiz).Error
}

func (db *PostgresSQLClient) CreateQuestions(questions []dbmodels.Question) error {
	return db.client.Create(&questions).Error
}

func (db *PostgresSQLClient) GetQuizzesByClassroomID(classroomID uint) ([]dbmodels.Quiz, error) {
	var quizzes []dbmodels.Quiz
	err := db.client.Where("classroom_id = ?", classroomID).Preload("Questions").Find(&quizzes).Error
	return quizzes, err
}

func (db *PostgresSQLClient) GetQuestionByID(id uint) (dbmodels.Question, error) {
	var question dbmodels.Question
	err := db.client.First(&question, id).Error
	return question, err
}

func (db *PostgresSQLClient) CreateResponse(resp *dbmodels.Response) error {
	return db.client.Create(resp).Error
}

func (db *PostgresSQLClient) GetResponsesByStudentID(studentID uint) ([]dbmodels.Response, error) {
	var responses []dbmodels.Response
	err := db.client.Where("student_id = ?", studentID).Preload("Question").Find(&responses).Error
	return responses, err
}
