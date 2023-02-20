package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabo-syu/bookmarks/interfaces/presenters"
	"github.com/tabo-syu/bookmarks/usecases"
)

type TagBookmarksController struct {
	tagBookmarks *usecases.TagBookmarksUsecase
	writer       *presenters.WebAPIPresenter
}

func NewTagBookmarksController(tagBookmarks *usecases.TagBookmarksUsecase, writer *presenters.WebAPIPresenter) *TagBookmarksController {
	return &TagBookmarksController{tagBookmarks, writer}
}

func (c *TagBookmarksController) List(g *gin.Context) {
	var req usecases.TagBookmarksListRequest
	if err := g.ShouldBindUri(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	bookmarks, err := c.tagBookmarks.List(g, &req)
	c.writer.Read(g, bookmarks, err)
}
