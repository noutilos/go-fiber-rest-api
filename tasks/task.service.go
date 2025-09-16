package tasks

import (
	"errors"
	"strconv"

	config "github.com/dev-khalid/go-fiber-rest-api/config"
)

type TaskService struct{}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (s *TaskService) Create(payload *Task) (*Task, error) {
	result := config.DB.Create(&payload)
	if result.Error != nil {
		return nil, result.Error
	}

	// TODO: Push this event to Kafka.

	return payload, nil
}

func (s *TaskService) Update(payload *Task) (*Task, error) {
	// Implement the logic to update a task
	return nil, nil
}

func (s *TaskService) Delete(id int) error {
	result := config.DB.Delete(&Task{}, id)

	if result.RowsAffected == 0 {
		return errors.New("Task not found for ID: " + strconv.Itoa(id))
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *TaskService) Get(id int) (Task, error) {
	var task Task

	result := config.DB.First(&task, id)
	if result.Error != nil {
		return Task{}, result.Error
	}

	return task, nil
}

func (s *TaskService) GetAll(queryParams *TaskQueryParams) ([]*Task, error) {
	var tasks []*Task
	
	query := config.DB.Model(&Task{})
	
	whereClause, values := queryParams.GetSQLConditions()
	if whereClause != "" {
		query = query.Where(whereClause, values...)
	}
	
	sortField := queryParams.GetSortField()
	sortDirection := queryParams.GetSortDirection()
	query = query.Order(sortField + " " + sortDirection)
	
	query = query.Offset(queryParams.GetOffset()).Limit(queryParams.Limit)
	
	result := query.Find(&tasks)
	
	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}