package handler

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

const (
	paramURL = "url"

	shortURLField = "shortURL"
	longURLField  = "longURL"
	messageField  = "message"
)

type input struct {
	Link string `json:"URL"`
}

func (h *Handler) getShort(c *gin.Context) {

	var in input
	if err := c.BindJSON(&in); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if _, err := url.ParseRequestURI(in.Link); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			messageField: "this link is not a valid URL",
		})
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

func (h *Handler) getLong(c *gin.Context) {
	link := c.Param(paramURL)
	if len(link) != 10 {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			messageField: "the short link must consist of 10 characters",
		})
		return
	}
	long, err := h.services.GetLongURL(link)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if long == "" {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			messageField: "there is no such short link yet",
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		longURLField: long,
	})
}
