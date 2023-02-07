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

}

func (p *WebAPIPresenter) Read(g *gin.Context, response any, err error) {
	if err == nil {
		g.JSON(http.StatusOK, response)

		return
	}

	var missing *gateways.MissingEntityError
	if errors.As(err, &missing) {
		g.JSON(http.StatusNotFound, err.Error())
	} else {
		g.JSON(http.StatusInternalServerError, err.Error())
	}
}

func (p *WebAPIPresenter) Update(g *gin.Context, response any, err error) {

}

func (p *WebAPIPresenter) Delete(g *gin.Context, response any, err error) {

}
