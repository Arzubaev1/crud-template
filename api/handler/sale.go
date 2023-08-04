package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetByID sale godoc
// @ID get_by_id_sale
// @Router /sale/{id} [GET]
// @Summary Get By ID Sale
// @Description Get By ID Sale
// @Tags Sale
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdSale(c *gin.Context) {

	var id string = c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Sale().GetByID(c.Request.Context(), &models.SalePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.sale.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id sale resposne", http.StatusOK, resp)
}

// GetList sale godoc
// @ID get_list_sale
// @Router /sale [GET]
// @Summary Get List Sale
// @Description Get List Sale
// @Tags Sale
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListSale(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list sale offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list sale limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Sale().GetList(c.Request.Context(), &models.SaleRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.sale.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list sale resposne", http.StatusOK, resp)
}
