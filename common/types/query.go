package types

import (
	"strconv"
	"strings"
)

type QueryParams[T any] struct {
	Fields map[string]interface{} `json:"fields"` // Generic field filters
	Page   int                    `json:"page"`   // Page number for pagination
	Limit  int                    `json:"limit"`  // Number of items per page
	Sort   string                 `json:"sort"`   // Sort field and direction (e.g., "created_at:desc")
}

// NewQueryParams creates a new QueryParams instance with default values
func NewQueryParams[T any]() *QueryParams[T] {
	return &QueryParams[T]{
		Fields: make(map[string]interface{}),
		Page:   1,
		Limit:  10,
		Sort:   "id:asc",
	}
}

func (q *QueryParams[T]) SetField(key string, value interface{}) {
	q.Fields[key] = value
}

func (q *QueryParams[T]) GetOffset() int {
	if q.Page <= 0 {
		q.Page = 1
	}
	return (q.Page - 1) * q.Limit
}

func (q *QueryParams[T]) GetSortField() string {
	if q.Sort == "" {
		return "id"
	}
	
	parts := strings.Split(q.Sort, ":")
	return parts[0]
}

func (q *QueryParams[T]) GetSortDirection() string {
	if q.Sort == "" {
		return "ASC"
	}
	
	parts := strings.Split(q.Sort, ":")
	if len(parts) > 1 && strings.ToUpper(parts[1]) == "DESC" {
		return "DESC"
	}
	return "ASC"
}

// ParseQueryString parses query string values into QueryParams
func (q *QueryParams[T]) ParseQueryString(queryMap map[string]string, allowedFields []string) {
	// Parse pagination
	if page, err := strconv.Atoi(queryMap["page"]); err == nil && page > 0 {
		q.Page = page
	}
	
	if limit, err := strconv.Atoi(queryMap["limit"]); err == nil && limit > 0 && limit <= 100 {
		q.Limit = limit
	}
	
	// Parse sort
	if sort := queryMap["sort"]; sort != "" {
		q.Sort = sort
	}
	
	// Parse field filters
	allowedFieldsMap := make(map[string]bool)
	for _, field := range allowedFields {
		allowedFieldsMap[field] = true
	}
	
	for key, value := range queryMap {
		// Skip pagination and sort parameters
		if key == "page" || key == "limit" || key == "sort" {
			continue
		}
		
		// Only allow specified fields
		if allowedFieldsMap[key] && value != "" {
			q.Fields[key] = value
		}
	}
}

// Validate ensures the query parameters are within acceptable ranges
func (q *QueryParams[T]) Validate() {
	if q.Page <= 0 {
		q.Page = 1
	}
	
	if q.Limit <= 0 {
		q.Limit = 10
	} else if q.Limit > 100 {
		q.Limit = 100
	}
	
	if q.Fields == nil {
		q.Fields = make(map[string]interface{})
	}
}