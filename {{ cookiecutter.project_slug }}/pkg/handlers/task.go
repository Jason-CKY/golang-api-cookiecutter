package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/core"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/schemas"
)

// @Summary Show all tasks
// @Description get all tasks
// @ID get-all-tasks
// @Accept  json
// @Produce  json
// @Success 200 {object} []schemas.Task
// @Failure 500 {object} echo.HTTPError
// @Router /tasks [get]
func GetAllTasks(c echo.Context) error {
	tasks, err := core.GetTasks()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return c.JSON(http.StatusOK, &tasks)
}

// @Summary Show task by ID
// @Description get task by ID
// @ID get-task-by-id
// @Accept  json
// @Produce  json
// @Param id  path string true "task id"
// @Success 200 {object} schemas.Task
// @Failure 500 {object} echo.HTTPError
// @Router /task/{id} [get]
func GetTaskById(c echo.Context) error {
	taskID := c.Param("id")
	task, err := core.GetTaskById(taskID)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return c.JSON(http.StatusOK, &task)
}

// @Summary Create Task
// @Description create new task
// @ID create-task
// @Accept  json
// @Produce  json
// @Param task body schemas.Task true "task"
// @Success 200 {object} schemas.Task
// @Failure 500 {object} echo.HTTPError
// @Router /task [post]
func CreateTask(c echo.Context) error {
	task := new(schemas.Task)
	err := c.Bind(task)
	if err != nil {
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if task.Id == "" {
		task.Id = uuid.New().String()
	}
	valErr := schemas.Validator.Struct(task)
	if valErr != nil {
		log.Error(valErr.Error())
		return echo.NewHTTPError(http.StatusBadRequest, valErr.Error())
	}
	createdTask, createErr := core.CreateTask(*task)
	if createErr != nil {
		log.Error(createErr.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, createdTask)
}

// @Summary Delete Task
// @Description Delete task by ID
// @ID delete-task-by-id
// @Accept  json
// @Produce  json
// @Param id  path string true "task id"
// @Success 200
// @Failure 500 {object} echo.HTTPError
// @Router /task/{id} [delete]
func DeleteTaskById(c echo.Context) error {
	taskID := c.Param("id")
	err := core.DeleteTaskById(taskID)
	if err != nil {
		return err
	}
	return c.String(http.StatusNoContent, "")
}

// @Summary Update Task
// @Description Update task
// @ID update-task
// @Accept  json
// @Produce  json
// @Param id  path string false "task id"
// @Param task body schemas.Task true "task"
// @Success 204 {object} schemas.Task
// @Failure 500 {object} echo.HTTPError
// @Router /task/{id} [put]
func UpdateTaskById(c echo.Context) error {
	taskID := c.Param("id")
	_, readErr := core.GetTaskById(taskID)
	if readErr != nil {
		log.Error(readErr.Error())
		return readErr
	}
	task := new(schemas.Task)
	err := c.Bind(task)
	if err != nil {
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if task.Id == "" {
		task.Id = taskID
	}
	if task.Id != taskID {
		return echo.NewHTTPError(http.StatusBadRequest, "task id in body not the same as path parameter")
	}
	valErr := schemas.Validator.Struct(task)
	if valErr != nil {
		log.Error(valErr.Error())
		return echo.NewHTTPError(http.StatusBadRequest, valErr.Error())
	}
	updatedTask, updateErr := core.UpdateTask(*task)
	if updateErr != nil {
		log.Error(updateErr.Error())
		return updateErr
	}
	return c.JSON(http.StatusOK, updatedTask)
}
