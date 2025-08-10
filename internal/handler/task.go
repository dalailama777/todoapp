package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	. "todo-app/internal/model"
	. "todo-app/internal/repository"
)

type CreateTaskModel struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// getTasksHandler возвращает список всех задач
// @Summary Получить список задач
// @Tags tasks
// @Produce json
// @Success 200 {array} Task
// @Router /tasks [get]
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks, err := GetTasks()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(tasks)
}

// createTaskHandler добавляет новую задачу
// @Summary Создать новую задачу
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body CreateTaskModel true "Новая задача"
// @Success 200 {object} Task
// @Failure 400 {string} string "Bad Request"
// @Router /tasks [post]
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task Task
	var taskModel CreateTaskModel
	err := json.NewDecoder(r.Body).Decode(&taskModel)
	task.Done = taskModel.Done
	task.Title = taskModel.Title
	if err != nil {
		http.Error(w, "incorrect json", http.StatusBadRequest)
		return
	}
	if err := task.TaskValidation(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	InsertTask(task)

	json.NewEncoder(w).Encode(task)
}

// getTaskByIDHandler возвращает задачу по ID
// @Summary Получить задачу по ID
// @Tags tasks
// @Produce json
// @Param id path int true "ID задачи"
// @Success 200 {object} Task
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Router /tasks/{id} [get]
func GetTaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Incorrect id", http.StatusBadRequest)
		return
	}

	task, err := GetTask(id)
	if err != nil {
		http.Error(w, "Id not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(task)
}

// updateTaskHandler обновляет задачу по ID
// @Summary Обновить задачу по ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Param task body Task true "Обновлённая задача"
// @Success 200 {string} string "task was updated"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Router /tasks/{id} [put]
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Incorrect id", http.StatusBadRequest)
		return
	}
	var task Task
	var taskModel CreateTaskModel
	err = json.NewDecoder(r.Body).Decode(&taskModel)
	task.Done = taskModel.Done
	task.Title = taskModel.Title
	if err != nil {
		http.Error(w, "Incorrect json", http.StatusBadRequest)
		return
	}

	if err := task.TaskValidation(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = UpdateTask(task, id)
	if err != nil {
		http.Error(w, "Id not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode("task was updated")
}

// deleteTaskHandler удаляет задачу по ID
// @Summary Удалить задачу по ID
// @Tags tasks
// @Produce json
// @Param id path int true "ID задачи"
// @Success 200 {string} string "task was deleted"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Router /tasks/{id} [delete]
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Incorrect id", http.StatusBadRequest)
		return
	}

	err = DeleteTask(id)
	if err != nil {
		http.Error(w, "Id not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode("task was deleted")
}
