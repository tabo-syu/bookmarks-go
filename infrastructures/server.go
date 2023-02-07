package infrastructures

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabo-syu/bookmarks/interfaces/controllers"
	"github.com/tabo-syu/bookmarks/interfaces/gateways"
	"github.com/tabo-syu/bookmarks/sqlc"
	"github.com/tabo-syu/bookmarks/usecases"
)

func NewServer(sqlc *sqlc.Queries) *http.Server {
	bookmarksGateway := gateways.NewBookmarksGateway(sqlc)
	// tagsRepo := gateways.NewTagsRepository(sqlc)
	// commentsRepo := gateways.NewCommentsRepository(sqlc)

	bookmarksUsecase := usecases.NewBookmarksUsecase(bookmarksGateway)

	bookmarks := controllers.NewBookmarksController(bookmarksUsecase)
	// tags := controllers.NewTagsController(tagsRepo)
	// comments := controllers.NewCommentsController(commentsRepo)

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		b := v1.Group("/bookmarks")
		{
			b.GET("", bookmarks.List)
			b.POST("", bookmarks.Create)
			// 		b.DELETE("", bookmarks.Delete)
		}

		// 	t := v1.Group("/tags")
		// 	{
		// 		t.GET("", tags.List)
		// 		t.POST("", tags.Create)
		// 		t.DELETE("", tags.Delete)
		// 	}

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