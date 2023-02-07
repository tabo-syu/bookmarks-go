package usecases

import "github.com/gin-gonic/gin"

type WebAPIOutput interface {
	Create(g *gin.Context, response interface{}, err error)
	Read(g *gin.Context, response interface{}, err error)
	Update(g *gin.Context, response interface{}, err error)
	Delete(g *gin.Context, response interface{}, err error)
}
