package schemas

import "github.com/go-playground/validator/v10"

// use a single instance of Validate, it caches struct info
var Validator *validator.Validate

type Task struct {
	Id          string `json:"id" validate:"uuid4"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Status      string `json:"status" validate:"required,oneof='backlog' 'progress' 'done'"`
}

type TaskSort struct {
	Id            string   `json:"id" validate:"required,uuid4"`
	Status        string   `json:"status" validate:"required"`
	Sorting_order []string `json:"sorting_order" validate:"required"`
}
