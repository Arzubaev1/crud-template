package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create User godoc
// @ID create_user
// @Router /user [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.CreateUser true "CreateUserRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *handler) CreateUser(c *gin.Context) {

	var createUser models.CreateUser

	err := c.ShouldBindJSON(&createUser) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create user", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.User().Create(context.Background(), &createUser)
	if err != nil {
		h.handlerResponse(c, "storage.user.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.user.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Security ApiKeyAuth
// Get By ID User godoc
// @ID get_by_id_user
// @Router /user/{id} [GET]
// @Summary Get By ID User
// @Description Get By ID User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *handler) GetByIdUser(c *gin.Context) {
	valuer, exists := c.Get("user_id")
	if !exists {
		h.handlerResponse(c, "getById user", http.StatusInternalServerError, "login first")
		return
	}
	userData := valuer.(helper.TokenInfo)
	id := ""
	if len(userData.UserID) > 0 {
		id = userData.UserID
	} else {
		id = c.Param("id")
	}

	resp, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.user.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get user by id", http.StatusCreated, resp)
}

// Get List User godoc
// @ID get_list_user
// @Router /user [GET]
// @Summary Get List User
// @Description Get List User
// @Tags User
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *handler) GetListUser(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list user", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list user", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.User().GetList(context.Background(), &models.GetListUserRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.user.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list user response", http.StatusOK, resp)
}

// Update User godoc
// @ID update_user
// @Router /user/{id} [PUT]
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param user body models.UpdateUser true "UpdateUserRequest"
// @Success 202 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *handler) UpdateUser(c *gin.Context) {

	var updateUser models.UpdateUser

	id := c.Param("id")

	err := c.ShouldBindJSON(&updateUser)
	if err != nil {
		h.handlerResponse(c, "update user", http.StatusBadRequest, err.Error())
		return
	}

	updateUser.Id = id

	rowsAffected, err := h.strg.User().Update(context.Background(), &updateUser)
	if err != nil {
		h.handlerResponse(c, "storage.user.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.user.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.user.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update user", http.StatusAccepted, resp)
}

// DELETE User godoc
// @ID delete_user
// @Router /user/{id} [DELETE]
// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param user body models.UserPrimaryKey true "DeleteUserRequest"
// @Success 204 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *handler) DeleteUser(c *gin.Context) {

	id := c.Param("id")

	rowsAffected, err := h.strg.User().Delete(context.Background(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.user.delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.user.delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "delete user", http.StatusNoContent, nil)
}
