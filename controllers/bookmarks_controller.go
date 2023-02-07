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

func (c *BookmarksController) List(g *gin.Context) {
}

func (c *BookmarksController) Create(g *gin.Context) {
}

func (c *BookmarksController) Delete(g *gin.Context) {
}
