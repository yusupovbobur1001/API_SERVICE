package handler

import (
	"api_gateway/api/models"
	pb "api_gateway/genproto/task_service"
	"net/http"

	"github.com/gin-gonic/gin"
)


// @Summary Create task
// @Tags task
// @Security     ApiKeyAuth
// @Param body body task_service.TaskRequest true "Task data"
// @Success 200 {object} task_service.Task
// @Failure 500 {object} models.Response
// @Router /task_service/tasks [post]
func (h *Handler) CreateTask(ctx *gin.Context) {
	request := pb.TaskRequest{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		HandleResponse(ctx, h.log, "error is while create user by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}

	response, err := h.service.TaskService().CreateTask(ctx, &request)
	if err != nil {
		HandleResponse(ctx, h.log, "error is while update  user by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	HandleResponse(ctx, h.log, "SUCCESSS", http.StatusOK, response)
}


// @Summary Update task by id
// @Tags task
// @Security     ApiKeyAuth
// @Param id path string true "Task id"
// @Param body body models.UpdateTask true "Task data"
// @Success 200 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /task_service/tasks/{id} [put]
func (h *Handler) UpdateTask(ctx *gin.Context) {
	request := models.UpdateTask{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		HandleResponse(ctx, h.log, "error is while create user by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}

	id := ctx.Param("id")

	response, err := h.service.TaskService().UpdateTask(ctx, &pb.Task{
		Id:        id,
		UserId:    request.UserId,
		Title:     request.Title,
	})
	if err != nil {
		HandleResponse(ctx, h.log, "error is while update  user by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	HandleResponse(ctx, h.log, "SUCCESSS", http.StatusOK, response)
}


// @Summary Get task by id
// @Tags task
// @Security     ApiKeyAuth
// @Param id path string true "Task id"
// @Success 200 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /task_service/tasks/{id} [get]
func (h *Handler) GetTask(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := h.service.TaskService().UpdateTask(ctx, &pb.Task{Id: id})
	if err != nil {
		HandleResponse(ctx, h.log, "error is while update  user by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	HandleResponse(ctx, h.log, "SUCCESSS", http.StatusOK, response)
}


// @Summary Get all tasks
// @Tags task
// @Security     ApiKeyAuth
// @Param page query int false "Page number"
// @Param limit query int false "Limit"
// @Param search query string false "Search query"
// @Success 200 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /task_service/tasks [get]
func (h *Handler) GetAllTasks(ctx *gin.Context) {
	requset := pb.GetListRequest{}

	err := ctx.ShouldBindJSON(&requset)
	if err != nil {
		HandleResponse(ctx, h.log, "error is while create user by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}

	response, err := h.service.TaskService().GetAllTasks(ctx, &requset)
	if err != nil {
		HandleResponse(ctx, h.log, "error is while update  user by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	HandleResponse(ctx, h.log, "SUCCESSS", http.StatusOK, response)
}

// @Summary Delete task by id
// @Tags task
// @Security     ApiKeyAuth
// @Param id path string true "Task id"
// @Success 200 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /task_service/tasks/{id} [delete]
func (h *Handler) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := h.service.TaskService().DeleteTask(ctx, &pb.PrimaryKey{Id: id})
	if err != nil {
		HandleResponse(ctx, h.log, "error is while update  user by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	HandleResponse(ctx, h.log, "SUCCESSS", http.StatusOK, response)
}

