package models

import "time"

type TaskStatus string

const (
	StatusProcessing TaskStatus = "processing"
	StatusSuccess    TaskStatus = "success"
	StatusFailed     TaskStatus = "failed"
	StatusCancelled  TaskStatus = "cancelled"
)

type TaskPriority string

const (
	PriorityHigh TaskPriority = "high"
	PriorityMed  TaskPriority = "med"
	PriorityLow  TaskPriority = "low"
)

type Task struct {
	ID        int          `json:"id" db:"id"`
	UserID    string       `json:"user_id" db:"user_id"`
	Title     string       `json:"title" db:"title"`
	Body      string       `json:"body" db:"body"`
	Status    TaskStatus   `json:"status" db:"status"`
	Priority  TaskPriority `json:"priority" db:"priority"`
	StartDate time.Time    `json:"start_date" db:"start_date"`
	EndDate   time.Time    `json:"end_date" db:"end_date"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
}

// NewTask creates a new task with default values
func NewTask(userID, title, body string, priority TaskPriority, startDate, endDate time.Time) *Task {
	return &Task{
		UserID:    userID,
		Title:     title,
		Body:      body,
		Status:    StatusProcessing, // Default status
		Priority:  priority,
		StartDate: startDate,
		EndDate:   endDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// UpdateStatus changes the task status and updates the timestamp
func (t *Task) UpdateStatus(status TaskStatus) {
	t.Status = status
	t.UpdatedAt = time.Now()
}

// UpdatePriority changes the task priority and updates the timestamp
func (t *Task) UpdatePriority(priority TaskPriority) {
	t.Priority = priority
	t.UpdatedAt = time.Now()
}

// IsCompleted checks if the task is completed (success status)
func (t *Task) IsCompleted() bool {
	return t.Status == StatusSuccess
}

// IsFailed checks if the task has failed
func (t *Task) IsFailed() bool {
	return t.Status == StatusFailed
}

// IsActive checks if the task is still being processed
func (t *Task) IsActive() bool {
	return t.Status == StatusProcessing
}
