package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabo-syu/bookmarks/interfaces/presenters"
	"github.com/tabo-syu/bookmarks/usecases"
)

type BookmarksController struct {
	bookmarks *usecases.BookmarksUsecase
	writer    *presenters.WebAPIPresenter
}

func NewBookmarksController(bookmarks *usecases.BookmarksUsecase, writer *presenters.WebAPIPresenter) *BookmarksController {
	return &BookmarksController{bookmarks, writer}
}

func (c *BookmarksController) List(g *gin.Context) {
	bookmarks, err := c.bookmarks.List(g)
	c.writer.Read(g, bookmarks, err)
}

func (c *BookmarksController) Create(g *gin.Context) {
	var req usecases.BookmarkCreateRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		// FIXME: usecase 層を介して presenter でレスポンスを返却したほうが良い？
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	bookmark, err := c.bookmarks.Create(g, &req)
	c.writer.Create(g, bookmark, err)
}

func (c *BookmarksController) Delete(g *gin.Context) {
}
