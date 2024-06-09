package handlers

import (
	"devquest-server/devquest/domain/models"
	"devquest-server/devquest/usecases"
	"devquest-server/devquest/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type TaskHttpHandler struct {
	taskUsecase usecases.TaskUsecase
}

func NewTaskHttpHandler(tUsecase usecases.TaskUsecase) *TaskHttpHandler {
	return &TaskHttpHandler{taskUsecase: tUsecase}
}

func (t *TaskHttpHandler) GetProjectTasks(w http.ResponseWriter, r *http.Request) {
	projectID, err := uuid.Parse(chi.URLParam(r, "project_id"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	projectTasks, err := t.taskUsecase.GetProjectTasks(projectID)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, projectTasks)
}

func (t *TaskHttpHandler) CreateNewTask(w http.ResponseWriter, r *http.Request) {
	projectID, err := uuid.Parse(chi.URLParam(r, "project_id"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	categoryID, err := uuid.Parse(r.URL.Query().Get("categoryID"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	var newTask models.CreateTaskDTO
	newTask.ProjectID = projectID
	newTask.CategoryID = categoryID

	err = utils.ReadJSON(w, r, &newTask)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = t.taskUsecase.CreateNewTask(newTask)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	res := utils.JSONResponse{
		Error: false,
		Message: "task successfully created",
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, res)
}

func (t *TaskHttpHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	taskID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	categoryID, err := uuid.Parse(r.URL.Query().Get("categoryID"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	var updatedTask models.UpdateTaskDTO
	updatedTask.CategoryID = categoryID
	
	err = utils.ReadJSON(w, r, &updatedTask)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = t.taskUsecase.UpdateTask(taskID, updatedTask)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	res := utils.JSONResponse{
		Error: false,
		Message: "task successfully updated",
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, res)
}

func (t *TaskHttpHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = t.taskUsecase.DeleteTask(taskID)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	res := utils.JSONResponse{
		Error: false,
		Message: "task successfully deleted",
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, res)
}

func (t *TaskHttpHandler) AcceptTask(w http.ResponseWriter, r *http.Request) {
	taskID, err := uuid.Parse(r.URL.Query().Get("taskID"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	developerID, err := uuid.Parse(r.URL.Query().Get("developerID"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = t.taskUsecase.AcceptTask(taskID, developerID)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	res := utils.JSONResponse{
		Error: false,
		Message: "task accepted by developer",
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, res)
}

func (t *TaskHttpHandler) CompleteTask(w http.ResponseWriter, r *http.Request) {
	taskID, err := uuid.Parse(r.URL.Query().Get("taskID"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	developerID, err := uuid.Parse(r.URL.Query().Get("developerID"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = t.taskUsecase.CompleteTask(taskID, developerID)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	res := utils.JSONResponse{
		Error: false,
		Message: "task completed by developer",
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, res)
}

func (t *TaskHttpHandler) CreateNewTaskCategory(w http.ResponseWriter, r *http.Request) {
	var newTaskCategory models.CreateTaskCategoryDTO
	err := utils.ReadJSON(w, r, &newTaskCategory)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = t.taskUsecase.CreateNewTaskCategory(newTaskCategory)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	res := utils.JSONResponse{
		Error: false,
		Message: "task category successfully created",
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, res)
}