package tasks

import (
	"time"
)

type TaskType string

const (
	GeneratePaymentReceipt TaskType = "generate_payment_receipt"
	SendEmail              TaskType = "send_email"
)

type TaskStatus string

const (
	Pending   TaskStatus = "pending"
	InProgress TaskStatus = "in_progress"
	Completed  TaskStatus = "completed"
	Failed     TaskStatus = "failed"
	Canceled   TaskStatus = "canceled"
)

type Task struct {
	ID          uint        `json:"id" gorm:"primaryKey"`
	Title       string     	`json:"title" validate:"required,min=3,max=255"`
	Description string     	`json:"description" validate:"max=1000"`
	Status      TaskStatus 	`json:"status" validate:"omitempty,oneof=pending in_progress completed failed canceled" gorm:"default:pending"`
	TaskType    TaskType   	`json:"task_type" validate:"omitempty,oneof=generate_payment_receipt send_email" gorm:"default:generate_payment_receipt"`
	UserID      int        	`json:"user_id" validate:"required,gt=0"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}