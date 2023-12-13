package handlers

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/components"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/core"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/schemas"
)

// GET /htmx
func TasksView(c echo.Context) error {
	backlogTaskList, progressTaskList, doneTaskList, echoHTTPErr := core.GetTasksInOrder()
	if echoHTTPErr != nil {
		return echoHTTPErr
	}

	component := components.TaskView(backlogTaskList, progressTaskList, doneTaskList)
	return component.Render(context.Background(), c.Response().Writer)
}

// DELETE /htmx/task/:id
func DeleteTaskView(c echo.Context) error {
	task_id := c.Param("id")
	echoHTTPErr := core.DeleteTaskById(task_id)
	if echoHTTPErr != nil {
		return echoHTTPErr
	}
	return c.String(http.StatusOK, "")
}

// POST /htmx/task/:id
func EditTaskView(c echo.Context) error {
	task_id := c.Param("id")
	task, echoHTTPErr := core.GetTaskById(task_id)
	if echoHTTPErr != nil {
		return echoHTTPErr
	}
	component := components.EditTask(task)
	return component.Render(context.Background(), c.Response().Writer)
}

// POST /htmx/task/empty/:status
func EmptyEditTaskView(c echo.Context) error {
	task_status := c.Param("status")
	task := schemas.Task{
		Id:          uuid.New().String(),
		Status:      task_status,
		Title:       "",
		Description: "",
	}
	component := components.EditTask(task)
	return component.Render(context.Background(), c.Response().Writer)
}

// DELETE /htmx/task/cancel/:id
func CancelEditTaskView(c echo.Context) error {
	task_id := c.Param("id")
	task, echoHTTPErr := core.GetTaskById(task_id)
	if echoHTTPErr != nil {
		if echoHTTPErr.Code == http.StatusNotFound {
			return c.String(http.StatusOK, "")
		}
		return echoHTTPErr
	}
	component := components.TaskSingleton(task)
	return component.Render(context.Background(), c.Response().Writer)
}

// PUT /htmx/task/:id
func UpdateTaskView(c echo.Context) error {
	task_id := c.Param("id")
	status := c.FormValue("status")
	title := c.FormValue("title")
	description := c.FormValue("description")
	new_task := schemas.Task{
		Id:          task_id,
		Title:       title,
		Description: description,
		Status:      status,
	}
	task, echoHTTPErr := core.UpsertTask(new_task)
	if echoHTTPErr != nil {
		return echoHTTPErr
	}
	component := components.TaskSingleton(task)
	return component.Render(context.Background(), c.Response().Writer)
}

// POST /htmx/sort/:status
func SortTaskView(c echo.Context) error {
	var sortTaskRequestParams schemas.SortTaskRequestParams
	err := c.Bind(&sortTaskRequestParams)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	taskSort, echoHTTPErr := core.GetTaskSortByStatus(sortTaskRequestParams.Status)
	if echoHTTPErr != nil {
		return echoHTTPErr
	}
	taskSort.Sorting_order = sortTaskRequestParams.TaskIds
	_, echoHTTPErr = core.UpdateTaskSort(taskSort)
	if echoHTTPErr != nil {
		return echoHTTPErr
	}
	tasks, echoHTTPErr := core.UpdateTasksStatusById(sortTaskRequestParams.TaskIds, sortTaskRequestParams.Status)
	if echoHTTPErr != nil {
		return echoHTTPErr
	}

	component := components.TaskList(tasks, sortTaskRequestParams.Status)
	return component.Render(context.Background(), c.Response().Writer)
}
