package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/p12s/2gis-catalog-api"
	"net/http"
)

// @Summary Create company
// @Tags Company
// @Description Creating company with minimal required data
// @ID createCompany
// @Accept  json
// @Produce  json
// @Param input body common.CompanyCreateRequest true "Company created info"
// @Success 200 {string} string "id"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/company/ [post]
func (h *Handler) createCompany(c *gin.Context) {
	var input common.CompanyCreateRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Company.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
