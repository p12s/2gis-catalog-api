package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Get all company in a rubric
// @Tags Rubric
// @Description Get all rubric company (without pagination)
// @ID getAllRubricCompany
// @Produce  json
// @Param rubric_id query integer true "Rubric id"
// @Success 200 {object} []common.CompanyResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/rubric/all-company [get]
func (h *Handler) getAllRubricCompany(c *gin.Context) {
	getParams := c.Request.URL.Query()
	rubricId, err := strconv.Atoi(getParams["rubric_id"][0])
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid rubric_id param")
		return
	}

	items, err := h.services.CompanyRubric.GetAllRubricCompany(rubricId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}
