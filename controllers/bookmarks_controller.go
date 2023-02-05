package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tabo-syu/bookmarks/services"
)

type BookmarksController struct {
	bookmarks *services.BookmarksService
}

func NewBookmarksController(bookmarks *services.BookmarksService) *BookmarksController {
	return &BookmarksController{bookmarks}
}

func (c *BookmarksController) List(ctx *gin.Context) {
}

func (c *BookmarksController) Create(ctx *gin.Context) {
}

func (c *BookmarksController) Delete(ctx *gin.Context) {
}
