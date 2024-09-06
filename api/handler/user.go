package handler

import (
	"api_gateway/api/models"
	pb "api_gateway/genproto/auth_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

 
// @Summary Get user by id
// @Tags user
// @Security     ApiKeyAuth
// @Param id path string true "User id"
// @Success 200 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /user_service/users/{id} [get]
func (h *Handler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := h.service.UserService().GetUser(ctx, &pb.PrimaryKey{Id: id})
	if err != nil {
		HandleResponse(ctx, h.log, "error is while get by id  user by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}

	HandleResponse(ctx, h.log, "SUCCESS", http.StatusOK, response)
}


// @Summary Get all users
// @Tags user
// @Security     ApiKeyAuth
// @Param page query int false "Page number"
// @Param limit query int false "Limit"
// @Param search query string false "Search query"
// @Success 200 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /user_service/users [get]
func (h *Handler) GetAllUsers(ctx *gin.Context) {
	request := pb.GetListRequest{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		HandleResponse(ctx, h.log, "error is while converting page. -~~~~~~ERROR====", http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.service.UserService().GetAllUsers(ctx, &request)
	if err != nil {
		HandleResponse(ctx, h.log, "error is while update  user by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	HandleResponse(ctx, h.log, "SUCCESSS", http.StatusOK, response)
}


// @Summary Delete user by id
// @Tags user
// @Security     ApiKeyAuth
// @Param id path string true "User id"
// @Success 200 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /user_service/users/{id} [delete]
func (h *Handler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := h.service.UserService().DeleteUser(ctx, &pb.PrimaryKey{Id: id})
	if err != nil {
		HandleResponse(ctx, h.log, "error is while update  user by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	HandleResponse(ctx, h.log, "SUCCESSS", http.StatusOK, response)
}


// @Summary Update user by id
// @Tags user
// @Security     ApiKeyAuth
// @Param id path string true "User id"
// @Param body body models.User true "User data"
// @Success 200 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /user_service/users/{id} [put]
func (h *Handler) UpdateUser(ctx *gin.Context) {

	request := models.User{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		HandleResponse(ctx, h.log, "error is while converting page. -~~~~~~ERROR====", http.StatusBadRequest, err.Error())
		return
	}
	id := ctx.Param("id")

	response, err := h.service.UserService().UpdateUser(ctx, &pb.User{
		Id:           id,
		UserName:     request.Email,
		Email:        request.Email,
		PasswordHash: request.PasswordHash,
		Role:         request.Role,
	})

	if err != nil {
		HandleResponse(ctx, h.log, "error is while update  user by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	HandleResponse(ctx, h.log, "SUCCESSS", http.StatusOK, response)
}