package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabo-syu/bookmarks/interfaces/presenters"
	"github.com/tabo-syu/bookmarks/usecases"
)

type CommentsController struct {
	comments *usecases.CommentsUsecase
	writer   *presenters.WebAPIPresenter
}

func NewCommentsController(comments *usecases.CommentsUsecase, writer *presenters.WebAPIPresenter) *CommentsController {
	return &CommentsController{comments, writer}
}

func (c *CommentsController) Get(g *gin.Context) {
	var req usecases.CommentGetRequest
	if err := g.ShouldBindUri(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	comment, err := c.comments.Get(g, &req)
	c.writer.Read(g, comment, err)
}

func (c *CommentsController) List(g *gin.Context) {
	var req usecases.CommentsListRequest
	if err := g.ShouldBindUri(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	comments, err := c.comments.List(g, &req)
	c.writer.Read(g, comments, err)
}

func (c *CommentsController) Create(g *gin.Context) {
	var (
		reqURI  usecases.CommentCreateURIRequest
		reqJSON usecases.CommentCreateJSONRequest
	)
	if err := g.ShouldBindUri(&reqURI); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	if err := g.ShouldBindJSON(&reqJSON); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	req := usecases.CommentCreateRequest{
		CommentCreateURIRequest:  reqURI,
		CommentCreateJSONRequest: reqJSON,
	}

	comment, err := c.comments.Create(g, &req)
	c.writer.Create(g, comment, err)
}

func (c *CommentsController) Delete(g *gin.Context) {
	var req usecases.CommentDeleteRequest
	if err := g.ShouldBindUri(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	comment, err := c.comments.Delete(g, &req)
	c.writer.Delete(g, comment, err)
}
