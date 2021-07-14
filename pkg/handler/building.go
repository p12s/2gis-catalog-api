package handler

import (
	"github.com/gin-gonic/gin"
	common "github.com/p12s/2gis-catalog-api"
	"net/http"
	"strconv"
)

// @Summary Get all company in a building
// @Tags Building
// @Description Get all company registered in a building
// @ID getAllCompany
// @Produce  json
// @Param city_id query integer true "City id"
// @Param street_id query integer true "Street id"
// @Param house query integer true "House id"
// @Success 200 {object} []common.CompanyResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/building/all-company [get]
func (h *Handler) getAllCompany(c *gin.Context) {
	getParams := c.Request.URL.Query()
	cityId, err := strconv.Atoi(getParams["city_id"][0])
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid city_id param")
		return
	}
	streetId, err := strconv.Atoi(getParams["street_id"][0])
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid street_id param")
		return
	}
	house, err := strconv.Atoi(getParams["house"][0])
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid house param")
		return
	}

	items, err := h.services.Building.GetAllCompany(cityId, streetId, house)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

// @Summary Add building
// @Tags Building
// @Description Add building into database
// @ID add
// @Produce  json
// @Param input body common.BuildingCreateRequest true "Building created info"
// @Success 200 {string} string "id"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/building/ [post]
func (h *Handler) add(c *gin.Context) {
	var input common.BuildingCreateRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Building.CreateNew(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
