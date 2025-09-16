package tasks

import (
	"strings"

	"github.com/dev-khalid/go-fiber-rest-api/common/types"
)

// TaskQueryParams represents query parameters specific to Task entity
type TaskQueryParams struct {
	*types.QueryParams[Task]
}

// NewTaskQueryParams creates a new TaskQueryParams instance with default values
func NewTaskQueryParams() *TaskQueryParams {
	return &TaskQueryParams{
		QueryParams: types.NewQueryParams[Task](),
	}
}

func (tq *TaskQueryParams) GetAllowedFields() []string {
	return []string{
		"title",
		"description", 
		"status",
		"task_type",
		"user_id",
	}
}

// ParseFromMap parses query parameters from a map (typically from fiber.Ctx.Queries())
func (tq *TaskQueryParams) ParseFromMap(queryMap map[string]string) {
	tq.ParseQueryString(queryMap, tq.GetAllowedFields())
	tq.Validate()
}

// GetSQLConditions returns SQL WHERE conditions and values for GORM
func (tq *TaskQueryParams) GetSQLConditions() (string, []interface{}) {
	var conditions []string
	var values []interface{}

	for field, value := range tq.Fields {
		switch field {
		case "title":
			conditions = append(conditions, "title LIKE ?")
			values = append(values, "%"+value.(string)+"%")
		case "description":
			conditions = append(conditions, "description LIKE ?")
			values = append(values, "%"+value.(string)+"%")
		case "status":
			conditions = append(conditions, "status = ?")
			values = append(values, value)
		case "task_type":
			conditions = append(conditions, "task_type = ?")
			values = append(values, value)
		case "user_id":
			conditions = append(conditions, "user_id = ?")
			values = append(values, value)
		}
	}

	if len(conditions) == 0 {
		return "", nil
	}

	return strings.Join(conditions, " AND "), values
}