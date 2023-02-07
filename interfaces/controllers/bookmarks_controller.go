package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tabo-syu/bookmarks/usecases"
)

type BookmarksController struct {
	bookmarks *usecases.BookmarksUsecase
}

func NewBookmarksController(bookmarks *usecases.BookmarksUsecase) *BookmarksController {
	return &BookmarksController{bookmarks}
}

func (c *BookmarksController) List(g *gin.Context) {
	c.bookmarks.List(g.Request.Context(), g)
}

func (c *BookmarksController) Create(g *gin.Context) {
}

func (c *BookmarksController) Delete(g *gin.Context) {
}
