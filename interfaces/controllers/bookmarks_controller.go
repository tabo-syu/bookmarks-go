package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabo-syu/bookmarks/domain"
	"github.com/tabo-syu/bookmarks/usecases"
)

type BookmarksController struct {
	bookmarks *usecases.BookmarksUsecase
}

func NewBookmarksController(bookmarks *usecases.BookmarksUsecase) *BookmarksController {
	return &BookmarksController{bookmarks}
}

func (c *BookmarksController) List(g *gin.Context) {
	c.bookmarks.List(g)
}

func (c *BookmarksController) Create(g *gin.Context) {
	var req domain.BookmarkCreateRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		// FIXME: usecase 層を介して presenter でレスポンスを返却したほうが良い？
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	c.bookmarks.Create(g, &req)
}

func (c *BookmarksController) Delete(g *gin.Context) {
}
