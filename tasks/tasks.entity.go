package tasks

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
	ID          int        `json:"id"`
	Title       string     `json:"title" validate:"required"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status" validate:"omitempty,oneof=pending in_progress completed failed canceled"`
	TaskType    TaskType   `json:"task_type" validate:"omitempty,oneof=generate_payment_receipt send_email"`
	UserID      int        `json:"user_id" validate:"required"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"updated_at"`
}