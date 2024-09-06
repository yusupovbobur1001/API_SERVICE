package handler

import (
	"api_gateway/api/models"
	pb "api_gateway/genproto/task_service"
	"net/http"

	"github.com/gin-gonic/gin"
)


// @Summary Create detail
// @Tags detail
// @Security     ApiKeyAuth
// @Param body body task_service.DetailRequest true "Detail data"
// @Success 200 {object} task_service.Detail
// @Failure 500 {object} models.Response
// @Router /task_service/details [post]
func (h *Handler) CreateDetail(ctx *gin.Context) {

	request := pb.DetailRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		HandleResponse(ctx, h.log, "error is while create user by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}

	response, err := h.service.DetailService().CreateDetail(ctx, &request)
	if err != nil {
		HandleResponse(ctx, h.log, "error is while update  user by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	HandleResponse(ctx, h.log, "SUCCESSS", http.StatusOK, response)
}


// @Summary Update detail by id
// @Tags detail
// @Security     ApiKeyAuth
// @Param id path string true "Detail id"
// @Param body body models.UpdateDetail true "Detail data"
// @Success 200 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /task_service/details/{id} [put]
func (h *Handler) UpdateDetail(ctx *gin.Context) {
	request := models.UpdateDetail{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		HandleResponse(ctx, h.log, "error is while create user by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}

	id := ctx.Param("id")

	response, err := h.service.DetailService().UpdateDetail(ctx, &pb.Detail{
		Id:          id,
		TaskId:      request.TaskId,
		Description: request.Description,
		Status:      request.Status,
		Priority:    request.Priority,
	})
	if err != nil {
		HandleResponse(ctx, h.log, "error is while update  user by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	HandleResponse(ctx, h.log, "SUCCESSS", http.StatusOK, response)
}


// @Summary Get detail by id
// @Tags detail
// @Security     ApiKeyAuth
// @Param id path string true "Detail id"
// @Success 200 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /task_service/details/{id} [get]
func (h *Handler) GetDetail(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := h.service.DetailService().GetDetail(ctx, &pb.PrimaryKey{Id: id})
	if err != nil {
		HandleResponse(ctx, h.log, "error is while update  user by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	HandleResponse(ctx, h.log, "SUCCESSS", http.StatusOK, response)
}


// @Summary Get all details
// @Tags detail
// @Security     ApiKeyAuth
// @Param page query int false "Page number"
// @Param limit query int false "Limit"
// @Param search query string false "Search query"
// @Success 200 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /task_service/details [get]
func (h *Handler) GetAllDetail(ctx *gin.Context) {
	requset := pb.GetListRequest{}

	err := ctx.ShouldBindJSON(&requset)
	if err != nil {
		HandleResponse(ctx, h.log, "error is while create user by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}

	response, err := h.service.DetailService().GetAllDetail(ctx, &requset)
	if err != nil {
		HandleResponse(ctx, h.log, "error is while update  user by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	HandleResponse(ctx, h.log, "SUCCESSS", http.StatusOK, response)
}


// @Summary Delete detail by id
// @Tags detail
// @Security     ApiKeyAuth
// @Param id path string true "Detail id"
// @Success 200 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /task_service/details/{id} [delete]
func (h *Handler) DeleteDetail(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := h.service.DetailService().DeleteDetail(ctx, &pb.PrimaryKey{Id: id})
	if err != nil {
		HandleResponse(ctx, h.log, "error is while update  user by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	HandleResponse(ctx, h.log, "SUCCESSS", http.StatusOK, response)
}




