package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/schemas"
)

// Get a list of all task ids sorts separated by status
// returns backlogSortOrder, progressSortOrder, doneSortOrder, error
func GetTaskSort() ([]string, []string, []string, *echo.HTTPError) {
	endpoint := fmt.Sprintf("%v/items/task_sorting", DirectusHost)
	res, httpErr := http.Get(endpoint)
	// error handling for http request
	if httpErr != nil {
		return []string{}, []string{}, []string{}, echo.NewHTTPError(res.StatusCode, httpErr.Error())
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	// error handling for anything above 2xx response
	if res.StatusCode > 299 {
		return []string{}, []string{}, []string{}, echo.NewHTTPError(res.StatusCode, httpErr.Error())
	}
	var httpResponse map[string][]schemas.TaskSort
	jsonErr := json.Unmarshal(body, &httpResponse)
	// error handling for json unmarshaling
	if jsonErr != nil {
		return []string{}, []string{}, []string{}, echo.NewHTTPError(http.StatusInternalServerError, jsonErr.Error())
	}

	backlogSortOrder, progressSortOrder, doneSortOrder := []string{}, []string{}, []string{}
	for _, taskSort := range httpResponse["data"] {
		if taskSort.Status == "backlog" {
			backlogSortOrder = taskSort.Sorting_order
		} else if taskSort.Status == "progress" {
			progressSortOrder = taskSort.Sorting_order
		} else {
			doneSortOrder = taskSort.Sorting_order
		}
	}
	return backlogSortOrder, progressSortOrder, doneSortOrder, nil
}

func GetTaskSortByStatus(status string) (schemas.TaskSort, *echo.HTTPError) {
	endpoint := fmt.Sprintf("%v/items/task_sorting?filter[status][_eq]=%v", DirectusHost, status)
	res, httpErr := http.Get(endpoint)
	// error handling for http request
	if httpErr != nil {
		return schemas.TaskSort{}, echo.NewHTTPError(res.StatusCode, httpErr.Error())
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	// error handling for anything above 2xx response
	if res.StatusCode > 299 {
		return schemas.TaskSort{}, echo.NewHTTPError(res.StatusCode, httpErr.Error())
	}
	var httpResponse map[string][]schemas.TaskSort
	jsonErr := json.Unmarshal(body, &httpResponse)
	// error handling for json unmarshaling
	if jsonErr != nil {
		return schemas.TaskSort{}, echo.NewHTTPError(http.StatusInternalServerError, jsonErr.Error())
	}
	if len(httpResponse["data"]) == 0 {
		return schemas.TaskSort{}, echo.NewHTTPError(http.StatusNotFound, "Item not found")
	}

	return httpResponse["data"][0], nil
}

func UpdateTaskSort(taskSort schemas.TaskSort) (schemas.TaskSort, *echo.HTTPError) {
	endpoint := fmt.Sprintf("%v/items/task_sorting/%v", DirectusHost, taskSort.Id)
	reqBody, _ := json.Marshal(taskSort)
	req, httpErr := http.NewRequest(http.MethodPatch, endpoint, bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	if httpErr != nil {
		return taskSort, echo.NewHTTPError(http.StatusInternalServerError, httpErr.Error())
	}
	client := &http.Client{}
	res, httpErr := client.Do(req)
	if httpErr != nil {
		return taskSort, echo.NewHTTPError(http.StatusInternalServerError, httpErr.Error())
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return taskSort, echo.NewHTTPError(res.StatusCode, string(body))
	}
	var taskResponse map[string]schemas.TaskSort
	jsonErr := json.Unmarshal(body, &taskResponse)
	// error handling for json unmarshaling
	if jsonErr != nil {
		return taskSort, echo.NewHTTPError(http.StatusInternalServerError, jsonErr.Error())
	}
	return taskResponse["data"], nil
}

func UpdateTaskSortByTasks(status string, tasks []schemas.Task) (schemas.TaskSort, *echo.HTTPError) {
	taskSort, echoHTTPErr := GetTaskSortByStatus(status)
	if echoHTTPErr != nil {
		return taskSort, echoHTTPErr
	}
	sortOrder := []string{}
	for _, task := range tasks {
		sortOrder = append(sortOrder, task.Id)
	}
	taskSort.Sorting_order = sortOrder
	updatedTaskSort, echoHTTPErr := UpdateTaskSort(taskSort)
	return updatedTaskSort, echoHTTPErr
}

func GetTasks() ([]schemas.Task, *echo.HTTPError) {
	endpoint := fmt.Sprintf("%v/items/task", DirectusHost)
	res, httpErr := http.Get(endpoint)
	// error handling for http request
	if httpErr != nil {
		return []schemas.Task{}, echo.NewHTTPError(res.StatusCode, httpErr.Error())
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	// error handling for anything above 2xx response
	if res.StatusCode > 299 {
		return []schemas.Task{}, echo.NewHTTPError(res.StatusCode, string(body))
	}
	var tasksResponse map[string][]schemas.Task
	jsonErr := json.Unmarshal(body, &tasksResponse)
	// error handling for json unmarshaling
	if jsonErr != nil {
		return []schemas.Task{}, echo.NewHTTPError(http.StatusInternalServerError, jsonErr.Error())
	}

	return tasksResponse["data"], nil
}

func FilterTaskById(task_id string, tasks []schemas.Task) schemas.Task {
	var filteredTask schemas.Task

	for _, task := range tasks {
		if task.Id == task_id {
			return task
		}
	}

	return filteredTask
}

func GetTasksInOrder() ([]schemas.Task, []schemas.Task, []schemas.Task, *echo.HTTPError) {
	tasks, echoHTTPErr := GetTasks()
	if echoHTTPErr != nil {
		return []schemas.Task{}, []schemas.Task{}, []schemas.Task{}, echoHTTPErr
	}
	backlogTaskSort, progressTaskSort, doneTaskSort, echoHTTPErr := GetTaskSort()
	if echoHTTPErr != nil {
		return []schemas.Task{}, []schemas.Task{}, []schemas.Task{}, echoHTTPErr
	}

	backlogTasks, progressTasks, doneTasks := []schemas.Task{}, []schemas.Task{}, []schemas.Task{}

	if len(backlogTaskSort)+len(progressTaskSort)+len(doneTaskSort) != len(tasks) {
		log.Error("Task Sorting not the same length as the total number of tasks!")
		for _, task := range tasks {
			if task.Status == "backlog" {
				backlogTasks = append(backlogTasks, task)
			} else if task.Status == "progress" {
				progressTasks = append(progressTasks, task)
			} else {
				doneTasks = append(doneTasks, task)
			}
		}
		var wg sync.WaitGroup
		wg.Add(3)
		goUpdateTaskSortByTasks := func(status string, tasks []schemas.Task) {
			defer wg.Done()
			UpdateTaskSortByTasks(status, tasks)
		}
		go goUpdateTaskSortByTasks("backlog", backlogTasks)
		go goUpdateTaskSortByTasks("progress", progressTasks)
		go goUpdateTaskSortByTasks("done", doneTasks)
		wg.Wait()
		return backlogTasks, progressTasks, doneTasks, nil

	}

	backlogTaskChan := make(chan []schemas.Task)
	progressTaskChan := make(chan []schemas.Task)
	doneTaskChan := make(chan []schemas.Task)
	goUpdateTasksBySortingOrder := func(taskSort []string, tasks []schemas.Task, taskChan chan []schemas.Task) {
		sortedTasks := []schemas.Task{}
		for _, taskId := range taskSort {
			sortedTasks = append(sortedTasks, FilterTaskById(taskId, tasks))
		}
		taskChan <- sortedTasks
	}
	go goUpdateTasksBySortingOrder(backlogTaskSort, tasks, backlogTaskChan)
	go goUpdateTasksBySortingOrder(progressTaskSort, tasks, progressTaskChan)
	go goUpdateTasksBySortingOrder(doneTaskSort, tasks, doneTaskChan)

	backlogTasks = <-backlogTaskChan
	progressTasks = <-progressTaskChan
	doneTasks = <-doneTaskChan

	return backlogTasks, progressTasks, doneTasks, nil
}

func GetTaskById(task_id string) (schemas.Task, *echo.HTTPError) {
	endpoint := fmt.Sprintf("%v/items/task?filter[id][_eq]=%v", DirectusHost, task_id)
	res, httpErr := http.Get(endpoint)
	// error handling for http request
	if httpErr != nil {
		return schemas.Task{}, echo.NewHTTPError(http.StatusInternalServerError, httpErr.Error())
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return schemas.Task{}, echo.NewHTTPError(res.StatusCode, string(body))
	}

	var taskResponse map[string][]schemas.Task
	jsonErr := json.Unmarshal(body, &taskResponse)
	// error handling for json unmarshaling
	if jsonErr != nil {
		return schemas.Task{}, echo.NewHTTPError(http.StatusInternalServerError, jsonErr.Error())
	}
	if len(taskResponse["data"]) == 0 {
		return schemas.Task{}, echo.NewHTTPError(http.StatusNotFound, "task not found")
	}

	return taskResponse["data"][0], nil
}

func DeleteTaskById(task_id string) *echo.HTTPError {
	log.Debugf("Deleting task id: %v...", task_id)
	task, echoHTTPErr := GetTaskById(task_id)
	if echoHTTPErr != nil {
		return echoHTTPErr
	}

	endpoint := fmt.Sprintf("%v/items/task/%v", DirectusHost, task_id)

	req, httpErr := http.NewRequest(http.MethodDelete, endpoint, nil)
	if httpErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, httpErr.Error())
	}
	client := &http.Client{}
	res, httpErr := client.Do(req)
	if httpErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, httpErr.Error())
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	if res.StatusCode != 204 {
		return echo.NewHTTPError(res.StatusCode, string(body))
	}

	taskSort, echoHTTPErr := GetTaskSortByStatus(task.Status)
	sortOrder := []string{}
	if echoHTTPErr != nil {
		return echoHTTPErr
	}
	for _, taskId := range taskSort.Sorting_order {
		if taskId != task_id {
			sortOrder = append(sortOrder, taskId)
		}
	}
	taskSort.Sorting_order = sortOrder
	_, echoHTTPErr = UpdateTaskSort(taskSort)
	if echoHTTPErr != nil {
		return echoHTTPErr
	}

	return nil
}

func UpdateTask(task schemas.Task) (schemas.Task, *echo.HTTPError) {
	endpoint := fmt.Sprintf("%v/items/task/%v", DirectusHost, task.Id)
	reqBody, _ := json.Marshal(task)
	req, httpErr := http.NewRequest(http.MethodPatch, endpoint, bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	if httpErr != nil {
		return task, echo.NewHTTPError(http.StatusInternalServerError, httpErr.Error())
	}
	client := &http.Client{}
	res, httpErr := client.Do(req)
	if httpErr != nil {
		return task, echo.NewHTTPError(http.StatusInternalServerError, httpErr.Error())
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return task, echo.NewHTTPError(res.StatusCode, string(body))
	}
	var taskResponse map[string]schemas.Task
	jsonErr := json.Unmarshal(body, &taskResponse)
	// error handling for json unmarshaling
	if jsonErr != nil {
		return task, echo.NewHTTPError(http.StatusInternalServerError, jsonErr.Error())
	}
	return taskResponse["data"], nil
}

func UpdateTasksStatusById(task_ids []string, status string) ([]schemas.Task, *echo.HTTPError) {
	endpoint := fmt.Sprintf("%v/items/task", DirectusHost)
	data := map[string]interface{}{
		"keys": task_ids,
		"data": map[string]interface{}{
			"status": status,
		},
	}
	reqBody, _ := json.Marshal(data)
	req, httpErr := http.NewRequest(http.MethodPatch, endpoint, bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	if httpErr != nil {
		return []schemas.Task{}, echo.NewHTTPError(http.StatusInternalServerError, httpErr.Error())
	}
	client := &http.Client{}
	res, httpErr := client.Do(req)
	if httpErr != nil {
		return []schemas.Task{}, echo.NewHTTPError(http.StatusInternalServerError, httpErr.Error())
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return []schemas.Task{}, echo.NewHTTPError(res.StatusCode, string(body))
	}
	var taskResponse map[string][]schemas.Task
	jsonErr := json.Unmarshal(body, &taskResponse)
	// error handling for json unmarshaling
	if jsonErr != nil {
		return []schemas.Task{}, echo.NewHTTPError(http.StatusInternalServerError, jsonErr.Error())
	}

	taskMapping := map[string]schemas.Task{}
	for _, task := range taskResponse["data"] {
		taskMapping[task.Id] = task
	}
	updatedTasks := []schemas.Task{}
	for _, taskId := range task_ids {
		updatedTasks = append(updatedTasks, taskMapping[taskId])
	}
	return updatedTasks, nil
}

func CreateTask(task schemas.Task) (schemas.Task, *echo.HTTPError) {
	endpoint := fmt.Sprintf("%v/items/task", DirectusHost)
	reqBody, _ := json.Marshal(task)
	req, httpErr := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	if httpErr != nil {
		return task, echo.NewHTTPError(http.StatusInternalServerError, httpErr.Error())
	}
	client := &http.Client{}
	res, httpErr := client.Do(req)
	if httpErr != nil {
		return task, echo.NewHTTPError(http.StatusInternalServerError, httpErr.Error())
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return task, echo.NewHTTPError(res.StatusCode, string(body))
	}
	var taskResponse map[string]schemas.Task
	jsonErr := json.Unmarshal(body, &taskResponse)
	// error handling for json unmarshaling
	if jsonErr != nil {
		return task, echo.NewHTTPError(http.StatusInternalServerError, jsonErr.Error())
	}

	taskSort, echoHTTPErr := GetTaskSortByStatus(task.Status)
	sortOrder := []string{task.Id}
	if echoHTTPErr != nil {
		return task, echoHTTPErr
	}
	taskSort.Sorting_order = append(sortOrder, taskSort.Sorting_order[:]...)
	_, echoHTTPErr = UpdateTaskSort(taskSort)

	if echoHTTPErr != nil {
		return task, echoHTTPErr
	}

	return taskResponse["data"], nil
}

func UpsertTask(task schemas.Task) (schemas.Task, *echo.HTTPError) {
	_, echoHTTPErr := GetTaskById(task.Id)
	if echoHTTPErr != nil {
		if echoHTTPErr.Code == http.StatusNotFound {
			newTask, echoHTTPErr := CreateTask(task)
			return newTask, echoHTTPErr
		}
		return task, echoHTTPErr
	}
	newTask, echoHTTPErr := UpdateTask(task)
	return newTask, echoHTTPErr
}
