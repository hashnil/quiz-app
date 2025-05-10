package models

import (
	"time"

	"gorm.io/datatypes"
)

type School struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"not null"`
	Location   string `gorm:"not null"`
	Classrooms []Classroom
	Users      []User
}

type Classroom struct {
	ID       uint   `gorm:"primaryKey"`
	SchoolID uint   `gorm:"not null"`
	Name     string `gorm:"not null"`
	School   School
}

type User struct {
	ID       uint `gorm:"primaryKey"`
	SchoolID uint `gorm:"not null"`
	Role     UserRole
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
	School   School
}

type UserRole string

const (
	StudentRole UserRole = "student"
	TeacherRole UserRole = "teacher"
)

type Quiz struct {
	ID          uint      `gorm:"primaryKey"`
	ClassroomID uint      `gorm:"not null"`
	TeacherID   uint      `gorm:"not null"`
	Title       string    `gorm:"not null"`
	StartTime   time.Time `gorm:"not null"`
	EndTime     time.Time `gorm:"not null"`
	Questions   []Question
}

type Question struct {
	ID            uint           `gorm:"primaryKey"`
	QuizID        uint           `gorm:"not null"`
	Content       string         `gorm:"not null"`
	Options       datatypes.JSON `gorm:"not null"`
	CorrectOption string         `gorm:"not null"`
	Quiz          Quiz
}

type Response struct {
	ID             uint      `gorm:"primaryKey"`
	StudentID      uint      `gorm:"not null"`
	QuestionID     uint      `gorm:"not null"`
	SelectedOption string    `gorm:"not null"`
	SubmittedAt    time.Time `gorm:"not null"`
	DurationMs     int       `gorm:"not null"`
	IsCorrect      bool      `gorm:"not null"`
	Question       Question
}

type Event struct {
	ID         uint   `gorm:"primaryKey"`
	UserID     uint   `gorm:"not null"`
	QuizID     *uint  `gorm:"null"`
	QuestionID *uint  `gorm:"null"`
	EventType  string `gorm:"not null"`
	SourceApp  EventSourceApp
	Timestamp  time.Time      `gorm:"not null"`
	Payload    datatypes.JSON `gorm:"not null"`
	User       User
}

type EventSourceApp string

const (
	WhiteboardApp EventSourceApp = "whiteboard"
	NotebookApp   EventSourceApp = "notebook"
)
