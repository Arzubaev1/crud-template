package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create saleProduct godoc
// @ID create_saleProduct
// @Router /saleProduct [POST]
// @Summary Create saleProduct
// @Description Create saleProduct
// @Tags saleProduct
// @Accept json
// @Procedure json
// @Param saleProduct body models.CreateSaleProduct true "CreateSaleProductRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateSaleProduct(c *gin.Context) {
	var (
		user_id string

		createSaleProduct models.CreateSaleProduct
	)

	valuer, exists := c.Get("user_id")
	if !exists {
		h.handlerResponse(c, "getById user", http.StatusInternalServerError, "login first")
		return
	}
	userData := valuer.(helper.TokenInfo)
	user_id = userData.UserID
	err := c.ShouldBindJSON(&createSaleProduct)
	if err != nil {
		h.handlerResponse(c, "error sales shouldBindJSON", http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.strg.SaleProduct().Create(c.Request.Context(), &createSaleProduct)
	if err != nil {
		h.handlerResponse(c, "error salesProduct createSaleProduct", http.StatusBadRequest, err.Error())
		return
	}
	resp, err := h.strg.SaleProduct().GetByID(c.Request.Context(), &models.SaleProductPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.salesProduct.getById", http.StatusInternalServerError, err.Error())
		return
	}
	body := map[string]interface{}{}
	body = map[string]interface{}{
		"sale_id": id,
		"user_id": user_id,
		"total":   resp.TotalPrice,
		"count":   resp.Count,
	}
	_, err = helper.DoRequest("http://localhost:8080/sale", "POST", body)

	h.handlerResponse(c, "create saleProduct  resposne", http.StatusCreated, resp)
}

func (h *handler) GetByIdSaleProduct(c *gin.Context) {

	var id string = c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.SaleProduct().GetByID(c.Request.Context(), &models.SaleProductPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.product.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id product resposne", http.StatusOK, resp)

}
