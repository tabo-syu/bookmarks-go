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

func (c *BookmarksController) Get(g *gin.Context) {
	var req usecases.BookmarkGetRequest
	if err := g.ShouldBindUri(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	bookmark, err := c.bookmarks.Get(g, &req)
	c.writer.Read(g, bookmark, err)
}

func (c *BookmarksController) List(g *gin.Context) {
	bookmarks, err := c.bookmarks.List(g)
	c.writer.Read(g, bookmarks, err)
}

func (c *BookmarksController) Create(g *gin.Context) {
	var req usecases.BookmarkCreateRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	bookmark, err := c.bookmarks.Create(g, &req)
	c.writer.Create(g, bookmark, err)
}

func (c *BookmarksController) Update(g *gin.Context) {
	var (
		reqURI  usecases.BookmarkUpdateURIRequest
		reqJSON usecases.BookmarkUpdateJSONRequest
	)
	if err := g.ShouldBindUri(&reqURI); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	if err := g.ShouldBindJSON(&reqJSON); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	req := usecases.BookmarkUpdateRequest{
		BookmarkUpdateURIRequest:  reqURI,
		BookmarkUpdateJSONRequest: reqJSON,
	}

	bookmark, err := c.bookmarks.Update(g, &req)
	c.writer.Update(g, bookmark, err)
}

func (c *BookmarksController) Delete(g *gin.Context) {
	var req usecases.BookmarkDeleteRequest
	if err := g.ShouldBindUri(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	bookmark, err := c.bookmarks.Delete(g, &req)
	c.writer.Delete(g, bookmark, err)
}
