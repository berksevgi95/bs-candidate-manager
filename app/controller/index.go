package controller

import (
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)

type Controller struct { 
}

func NewController() *Controller {
	return &Controller{}
}

// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} object models.Post
// @Header 200 {string} Token "qwerty"
// @Failure default {object} object httputil.DefaultError
// @Router /accounts/{id} [get]
func (c *Controller) ShowAccount(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.Post{Title: "asd", Body : "cscc"})
}

// ListAccounts godoc
// @Summary List accounts
// @Description get accounts
// @Accept  json
// @Produce  json
// @Success 200 {array} object models.Post
// @Header 200 {string} Token "qwerty"
// @Failure default {object} object httputil.DefaultError
// @Router /accounts [get]
func (c *Controller) ListAccounts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.Post{Title: "asd", Body : "cscc"})
}
