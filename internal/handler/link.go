package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	paramURL = "url"

	shortURLField = "shortURL"
	longURLField  = "longURL"
)

type input struct {
	Link string `json:"URL"`
}

//todo bind json without struct??

func (h *Handler) getShort(c *gin.Context) {

	// проверить, что линк - линк

	var in input
	if err := c.BindJSON(&in); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	shortURL, err := h.services.GetShortURL(in.Link)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		shortURLField: shortURL,
	})
}

// Перецепиться на get параметр
func (h *Handler) getLong(c *gin.Context) {
	link := c.Param(paramURL)

	longURL, err := h.services.GetLongURL(link)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	//todo bad request if error==not found??
	c.JSON(http.StatusOK, map[string]interface{}{
		longURLField: longURL,
	})
}
