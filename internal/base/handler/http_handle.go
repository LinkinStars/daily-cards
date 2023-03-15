package handler

import (
	"errors"
	"net/http"

	"github.com/LinkinStars/go-scaffold/logger"
	"github.com/LinkinStars/go-scaffold/mistake"
	"github.com/gin-gonic/gin"
)

func HandleResponse(c *gin.Context, err error, data interface{}) {
	if err == nil {
		c.JSON(http.StatusOK, SucceedRespBodyAndData("success", data))
		return
	}

	var myErr *mistake.Error
	if !errors.As(err, &myErr) {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, ErrRespBody("Internal Server Error"))
	}

	if mistake.IsBadRequest(myErr) {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, ErrRespBody(myErr.Message))
		return
	}

	if mistake.IsUnauthorized(myErr) {
		respBody := NewRespBody(-1, http.StatusText(http.StatusUnauthorized), "")
		c.JSON(http.StatusUnauthorized, respBody)
		return
	}

	if mistake.IsInternalServer(myErr) {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, ErrRespBody(myErr.Message))
		return
	}

	logger.Error(err)
	c.JSON(http.StatusInternalServerError, ErrRespBody("Internal Server Error"))
}
