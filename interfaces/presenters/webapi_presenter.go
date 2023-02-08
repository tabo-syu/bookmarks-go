package presenters

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabo-syu/bookmarks/interfaces/gateways"
	"github.com/tabo-syu/bookmarks/usecases"
)

type WebAPIPresenter struct{}

func NewWebAPIPresenter() usecases.WebAPIOutput {
	return &WebAPIPresenter{}
}

func (p *WebAPIPresenter) Create(g *gin.Context, response any, err error) {
	if err == nil {
		g.JSON(http.StatusCreated, response)

		return
	}

	var invalid *gateways.ValidationError
	if errors.As(err, &invalid) {
		g.JSON(http.StatusBadRequest, p.error(err))
	} else {
		g.JSON(http.StatusInternalServerError, p.error(err))
	}
}

func (p *WebAPIPresenter) Read(g *gin.Context, response any, err error) {
	if err == nil {
		g.JSON(http.StatusOK, response)

		return
	}

	var missing *gateways.MissingEntityError
	if errors.As(err, &missing) {
		g.JSON(http.StatusNotFound, p.error(err))
	} else {
		g.JSON(http.StatusInternalServerError, p.error(err))
	}
}

func (p *WebAPIPresenter) Update(g *gin.Context, response any, err error) {
	if err == nil {
		g.JSON(http.StatusOK, response)

		return
	}

	var (
		missing *gateways.MissingEntityError
		invalid *gateways.ValidationError
	)
	if errors.As(err, &missing) {
		g.JSON(http.StatusNotFound, p.error(err))
	} else if errors.As(err, &invalid) {
		g.JSON(http.StatusBadRequest, p.error(err))
	} else {
		g.JSON(http.StatusInternalServerError, p.error(err))
	}
}

func (p *WebAPIPresenter) Delete(g *gin.Context, response any, err error) {
	if err == nil {
		g.JSON(http.StatusOK, response)

		return
	}

	var missing *gateways.MissingEntityError
	if errors.As(err, &missing) {
		g.JSON(http.StatusNotFound, p.error(err))
	} else {
		g.JSON(http.StatusInternalServerError, p.error(err))
	}
}

func (p *WebAPIPresenter) error(err error) gin.H {
	return gin.H{"error": err.Error()}
}
