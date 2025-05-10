# Quiz Reporting APP Documentation

This document outlines the API contract for the Quiz Reporting backend service that powers the Whiteboard (teacher) and Notebook (student) applications.

**Base URL**: `http://localhost:8090/api/v1`
**GitHub Repository**: [https://github.com/hashnil/quiz-app](https://github.com/hashnil/quiz-app)

## Table of Contents

- [Authentication](#authentication)
- [Quiz Management](#quiz-management)
- [Student Responses](#student-responses)
- [Event Ingestion](#event-ingestion)
- [Reports](#reports)
- [Data Models](#data-models)

## Authentication

Authentication details will be added in a future update. Currently, the API endpoints do not require authentication for testing purposes.

## Quiz Management

### Create Quiz

Creates a new quiz with questions.

**Endpoint**: `POST /quizzes`  
**Tags**: `Quiz`

**Request Body**:
```json
{
  "classroom_id": 123,
  "teacher_id": 456,
  "title": "Math Quiz - Fractions",
  "start_time": "2025-05-15T09:00:00Z",
  "end_time": "2025-05-15T09:30:00Z",
  "questions": [
    {
      "content": "What is 1/2 + 1/4?",
      "options": {
        "A": "1/6",
        "B": "3/4",
        "C": "2/4",
        "D": "3/6"
      },
      "correct_option": "B"
    },
    {
      "content": "What is 2/3 of 18?",
      "options": {
        "A": "6",
        "B": "9",
        "C": "12",
        "D": "15"
      },
      "correct_option": "C"
    }
  ]
}
```

**Response** (200 OK):
```json
{
  "quiz_id": 789,
  "message": "Quiz created successfully"
}
```

### List Quizzes

Retrieves all quizzes for a specific classroom.

**Endpoint**: `GET /quizzes?classroom_id={classroom_id}`  
**Tags**: `Quiz`

**Query Parameters**:
- `classroom_id` (required): ID of the classroom

**Response** (200 OK):
```json
[
  {
    "id": 789,
    "classroom_id": 123,
    "teacher_id": 456,
    "title": "Math Quiz - Fractions",
    "start_time": "2025-05-15T09:00:00Z",
    "end_time": "2025-05-15T09:30:00Z"
  },
  {
    "id": 790,
    "classroom_id": 123,
    "teacher_id": 456,
    "title": "Science Quiz - Planets",
    "start_time": "2025-05-16T10:00:00Z",
    "end_time": "2025-05-16T10:30:00Z"
  }
]
```

## Student Responses

### Submit Answer

Submits a student's answer to a quiz question.

**Endpoint**: `POST /quizzes/{quizID}/submit`  
**Tags**: `Notebook`

**Path Parameters**:
- `quizID` (required): ID of the quiz

**Request Body**:
```json
{
  "student_id": 101,
  "question_id": 501,
  "selected_option": "B",
  "duration_ms": 15000
}
```

**Response** (200 OK):
```json
{
  "id": 1001,
  "student_id": 101,
  "question_id": 501,
  "selected_option": "B",
  "submitted_at": "2025-05-15T09:10:15Z",
  "duration_ms": 15000,
  "is_correct": true
}
```

### Submit Structured Response

Alternative endpoint for submitting structured responses.

**Endpoint**: `POST /responses`  
**Tags**: `Ingestion`

**Request Body**:
```json
{
  "student_id": 102,
  "question_id": 502,
  "selected_option": "A",
  "submitted_at": "2025-05-15T09:15:20Z",
  "duration_ms": 12500,
  "is_correct": false
}
```

**Response** (200 OK): Empty response with status 200.

## Event Ingestion

### Ingest Events

Ingests client-side events from Whiteboard and Notebook apps.

**Endpoint**: `POST /events`  
**Tags**: `Ingestion`

**Request Body**:
```json
[
  {
    "event_name": "student.join.quiz",
    "timestamp": "2025-05-15T09:00:05Z",
    "user_id": "101",
    "quiz_id": "789",
    "source_app": "notebook",
    "payload": {
      "device_type": "tablet",
      "app_version": "2.1.0"
    }
  },
  {
    "event_name": "student.submit.answer",
    "timestamp": "2025-05-15T09:05:10Z",
    "user_id": "101",
    "quiz_id": "789",
    "question_id": "501",
    "source_app": "notebook",
    "payload": {
      "selected_option": "B",
      "duration_ms": 15000
    }
  }
]
```

**Response** (200 OK):
```json
{
  "status": "success",
  "processed": 2
}
```

## Reports

### Student Performance Report

Retrieves performance metrics for a specific student.

**Endpoint**: `GET /reports/student/{studentID}`  
**Tags**: `Reports`

**Path Parameters**:
- `studentID` (required): ID of the student

**Response** (200 OK):
```json
{
  "total_attempted": 15,
  "correct_answers": 12,
  "accuracy": 0.8,
  "responses": [
    {
      "id": 1001,
      "student_id": 101,
      "question_id": 501,
      "selected_option": "B",
      "submitted_at": "2025-05-15T09:10:15Z",
      "duration_ms": 15000,
      "is_correct": true
    },
    {
      "id": 1002,
      "student_id": 101,
      "question_id": 502,
      "selected_option": "A",
      "submitted_at": "2025-05-15T09:12:30Z",
      "duration_ms": 12000,
      "is_correct": false
    }
    // Additional responses...
  ]
}
```

### Classroom Engagement Report

Retrieves engagement metrics for a specific classroom.

**Endpoint**: `GET /reports/classroom/{classroom_id}/engagement`  
**Tags**: `Reports`

**Path Parameters**:
- `classroom_id` (required): ID of the classroom

**Response** (200 OK): JSON object with engagement metrics (schema details to be specified).

### Content Effectiveness Report

Analyzes question-level effectiveness for a specific quiz.

**Endpoint**: `GET /reports/quiz/{quiz_id}/content-effectiveness`  
**Tags**: `Reports`

**Path Parameters**:
- `quiz_id` (required): ID of the quiz

**Response** (200 OK): JSON object with effectiveness metrics (schema details to be specified).

## Data Models

### Quiz

```json
{
  "id": 789,
  "classroom_id": 123,
  "teacher_id": 456,
  "title": "Math Quiz - Fractions",
  "start_time": "2025-05-15T09:00:00Z",
  "end_time": "2025-05-15T09:30:00Z"
}
```

### Response

```json
{
  "id": 1001,
  "student_id": 101,
  "question_id": 501,
  "selected_option": "B",
  "submitted_at": "2025-05-15T09:10:15Z",
  "duration_ms": 15000,
  "is_correct": true
}
```

### Event

```json
{
  "event_name": "student.submit.answer",
  "timestamp": "2025-05-15T09:05:10Z",
  "user_id": "101",
  "quiz_id": "789",
  "question_id": "501",
  "source_app": "notebook",
  "payload": {
    "selected_option": "B",
    "duration_ms": 15000
  }
}
```

### Create Quiz Request

```json
{
  "classroom_id": 123,
  "teacher_id": 456,
  "title": "Math Quiz - Fractions",
  "start_time": "2025-05-15T09:00:00Z",
  "end_time": "2025-05-15T09:30:00Z",
  "questions": [
    {
      "content": "What is 1/2 + 1/4?",
      "options": {
        "A": "1/6",
        "B": "3/4",
        "C": "2/4",
        "D": "3/6"
      },
      "correct_option": "B"
    }
  ]
}
```

### Submit Answer Request

```json
{
  "student_id": 101,
  "question_id": 501,
  "selected_option": "B",
  "duration_ms": 15000
}
```

## Error Handling

All endpoints return standard HTTP status codes:

- `200 OK`: Request succeeded
- `400 Bad Request`: Invalid request parameters
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server-side error

Error responses include a JSON body with error details:

```json
{
  "error": "Invalid request",
  "message": "Question ID is required",
  "code": "INVALID_PARAM"
}
```

## Event Taxonomy

Events follow the naming convention: `<actor>.<action>.<object>`

| Event Name | Source App | Description |
|------------|------------|-------------|
| student.join.quiz | Notebook | Student enters quiz session |
| student.submit.answer | Notebook | Answer submitted for a question |
| student.focus.lost | Notebook | App goes out of focus |
| teacher.start.quiz | Whiteboard | Quiz session begins |
| teacher.push.question | Whiteboard | New question shown to students |
| teacher.end.quiz | Whiteboard | Quiz session ends |
| system.sync.latency | Backend | Latency between push and receive |
| system.quiz.timeout | Notebook | Student missed deadline |

## API Versioning

The API is versioned in the URL path. The current version is `v1`.

## Rate Limiting

Rate limiting information will be provided in HTTP headers:
- `X-Rate-Limit-Limit`: Number of allowed requests in the current period
- `X-Rate-Limit-Remaining`: Number of remaining requests in the current period
- `X-Rate-Limit-Reset`: Time at which the rate limit resets, in UTC epoch seconds

## Contact

For API support or questions, please contact the development team via the [GitHub repository](https://github.com/hashnil/quiz-app) by creating an issue.
