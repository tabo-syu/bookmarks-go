package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabo-syu/bookmarks/interfaces/presenters"
	"github.com/tabo-syu/bookmarks/usecases"
)

type BookmarkTagsController struct {
	bookmarkTags *usecases.BookmarkTagsUsecase
	writer       *presenters.WebAPIPresenter
}

func NewBookmarkTagsController(bookmarkTags *usecases.BookmarkTagsUsecase, writer *presenters.WebAPIPresenter) *BookmarkTagsController {
	return &BookmarkTagsController{bookmarkTags, writer}
}

func (c *BookmarkTagsController) List(g *gin.Context) {
	var req usecases.BookmarkTagsListRequest
	if err := g.ShouldBindUri(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	tags, err := c.bookmarkTags.List(g, &req)
	c.writer.Read(g, tags, err)
}

func (c *BookmarkTagsController) Add(g *gin.Context) {
	var req usecases.BookmarkTagsAddRequest
	if err := g.ShouldBindUri(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	bookmarks, err := c.bookmarkTags.Add(g, &req)
	c.writer.Create(g, bookmarks, err)
}

func (c *BookmarkTagsController) Remove(g *gin.Context) {
	var req usecases.BookmarkTagsRemoveRequest
	if err := g.ShouldBindUri(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	bookmarks, err := c.bookmarkTags.Remove(g, &req)
	c.writer.Delete(g, bookmarks, err)
}
