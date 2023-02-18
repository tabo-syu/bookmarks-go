package infrastructures

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabo-syu/bookmarks/interfaces/controllers"
	"github.com/tabo-syu/bookmarks/interfaces/gateways"
	"github.com/tabo-syu/bookmarks/interfaces/presenters"
	"github.com/tabo-syu/bookmarks/sqlc"
	"github.com/tabo-syu/bookmarks/usecases"
)

func NewServer(sqlc *sqlc.Queries) *http.Server {
	webapiPresenter := presenters.NewWebAPIPresenter()

	bookmarks := controllers.NewBookmarksController(
		usecases.NewBookmarksUsecase(
			gateways.NewBookmarksGateway(sqlc),
		),
		webapiPresenter,
	)
	tags := controllers.NewTagsController(
		usecases.NewTagsUsecase(
			gateways.NewTagsGateway(sqlc),
		),
		webapiPresenter,
	)

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		b := v1.Group("/bookmarks")
		{
			b.GET("", bookmarks.List)
			b.GET("/:bookmark_id", bookmarks.Get)
			b.POST("", bookmarks.Create)
			b.PUT("/:bookmark_id", bookmarks.Update)
			b.DELETE("/:bookmark_id", bookmarks.Delete)
		}

		t := v1.Group("/tags")
		{
			t.GET("", tags.List)
			t.GET("/:id", tags.Get)
			t.POST("", tags.Create)
			t.PUT("/:id", tags.Update)
			t.DELETE("/:id", tags.Delete)
		}

		// 	c := v1.Group("/comments")
		// 	{
		// 		c.GET("", comments.List)
		// 		c.POST("", comments.Create)
		// 		c.DELETE("", comments.Delete)
		// 	}
	}

	return &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
}
