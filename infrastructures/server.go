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
	bookmarksGateway := gateways.NewBookmarksGateway(sqlc)
	tagsGateway := gateways.NewTagsGateway(sqlc)
	commentsGateway := gateways.NewCommentsGateway(sqlc)
	bookmarkTagsGateway := gateways.NewBookmarkTagsGateway(sqlc)

	webapiPresenter := presenters.NewWebAPIPresenter()

	bookmarks := controllers.NewBookmarksController(
		usecases.NewBookmarksUsecase(bookmarksGateway),
		webapiPresenter,
	)
	tags := controllers.NewTagsController(
		usecases.NewTagsUsecase(tagsGateway),
		webapiPresenter,
	)
	comments := controllers.NewCommentsController(
		usecases.NewCommentsUsecase(commentsGateway),
		webapiPresenter,
	)
	bookmarkTags := controllers.NewBookmarkTagsController(
		usecases.NewBookmarkTagsUsecase(bookmarksGateway, tagsGateway, bookmarkTagsGateway),
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

			b.GET("/:bookmark_id/tags", bookmarkTags.List)
			b.POST("/:bookmark_id/tags/:tag_id", bookmarkTags.Add)
			b.DELETE("/:bookmark_id/tags/:tag_id", bookmarkTags.Remove)

			b.GET("/:bookmark_id/comments", comments.List)
			b.POST("/:bookmark_id/comments", comments.Create)
		}

		t := v1.Group("/tags")
		{
			t.GET("", tags.List)
			t.GET("/:tag_id", tags.Get)
			t.POST("", tags.Create)
			t.PUT("/:tag_id", tags.Update)
			t.DELETE("/:tag_id", tags.Delete)
			// t.GET("/:tag_id/bookmarks/", tagBookmarks.List)
		}

		c := v1.Group("/comments")
		{
			c.GET("/:comment_id", comments.Get)
			c.DELETE("/:comment_id", comments.Delete)
		}
	}

	return &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
}
