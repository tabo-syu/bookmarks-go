package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabo-syu/bookmarks/interfaces/presenters"
	"github.com/tabo-syu/bookmarks/usecases"
)

type TagsController struct {
	tags   *usecases.TagsUsecase
	writer *presenters.WebAPIPresenter
}

func NewTagsController(tags *usecases.TagsUsecase, writer *presenters.WebAPIPresenter) *TagsController {
	return &TagsController{tags, writer}
}

func (c *TagsController) Get(g *gin.Context) {
	var req usecases.TagGetRequest
	if err := g.ShouldBindUri(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	tag, err := c.tags.Get(g, &req)
	c.writer.Read(g, tag, err)
}

func (c *TagsController) List(g *gin.Context) {
	tags, err := c.tags.List(g)
	c.writer.Read(g, tags, err)
}

func (c *TagsController) Create(g *gin.Context) {
	var req usecases.TagCreateRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	tag, err := c.tags.Create(g, &req)
	c.writer.Create(g, tag, err)
}

func (c *TagsController) Update(g *gin.Context) {
	var (
		reqURI  usecases.TagUpdateURIRequest
		reqJSON usecases.TagUpdateJSONRequest
	)
	if err := g.ShouldBindUri(&reqURI); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	if err := g.ShouldBindJSON(&reqJSON); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	req := usecases.TagUpdateRequest{
		TagUpdateURIRequest:  reqURI,
		TagUpdateJSONRequest: reqJSON,
	}

	tag, err := c.tags.Update(g, &req)
	c.writer.Update(g, tag, err)
}

func (c *TagsController) Delete(g *gin.Context) {
	var req usecases.TagDeleteRequest
	if err := g.ShouldBindUri(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	tag, err := c.tags.Delete(g, &req)
	c.writer.Delete(g, tag, err)
}
