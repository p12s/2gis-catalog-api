package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/p12s/2gis-catalog-api"
	"net/http"
)

func (h *Handler) createCity(c *gin.Context) {
	var input common.City
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.City.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
