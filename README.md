# Bookmarks API

```bash
$ git clone https://github.com/tabo-syu/bookmarks.git
$ cd bookmarks
$ go run cmd/migrator/main.go
$ go run cmd/seeder/main.go
$ go run cmd/server/main.go
```

## v1

### bookmarks

```bash
# ListBookmarks
$ curl -H {Token} http://localhost:8080/v1/bookmarks | jq -r .

# GetBookmark
$ curl -H {Token} http://localhost:8080/v1/bookmarks/{bookmark_id} | jq -r .

# CreateBookmark
$ curl -X POST -H {Token} http://localhost:8080/v1/bookmarks -d '{"url":"https://tabo-syu.com","title":"bookmark title", "description": "bookmark desc"}' -v | jq -r .

# UpdateBookmark
$ curl -X PUT -H {Token} http://localhost:8080/v1/bookmarks/{bookmark_id} -d '{"url":"https://tabo-syu.com","title":"bookmark title", "description": "bookmark desc"}' -v | jq -r .

# DeleteBookmark
$ curl -X DELETE -H {Token} http://localhost:8080/v1/bookmarks/{bookmark_id} | jq -r .

# FindBookmarksByTag
$ curl -H {Token} http://localhost:8080/v1/tags/{tag_id}/bookmarks | jq -r .
```

### tags

```bash
# ListTags
$ curl -H {Token} http://localhost:8080/v1/tags | jq -r .

# GetTag
$ curl -H {Token} http://localhost:8080/v1/tags/{tag_id} | jq -r .

# CreateTag
$ curl -X POST -H {Token} http://localhost:8080/v1/tags -d '{"name":"tag name","color":"#123456"}' -v | jq -r .

# UpdateTag
$ curl -X PUT -H {Token} http://localhost:8080/v1/tags/{tag_id} -d '{"name":"tag name","color":"#123456"}' -v | jq -r .

# DeleteTag
$ curl -X DELETE -H {Token} http://localhost:8080/v1/tags/{tag_id} | jq -r .

# FindTagsByBookmark
$ curl -H {Token} http://localhost:8080/v1/bookmarks/{bookmark_id}/tags | jq -r .

# AddTagToBookmark
$ curl -X POST -H {Token} http://localhost:8080/v1/bookmarks/{bookmark_id}/tags/{tag_id} | jq -r .

# DeleteTagFromBookmark
$ curl -X DELETE -H {Token} http://localhost:8080/v1/bookmarks/{bookmark_id}/tags/{tag_id} | jq -r .
```

### comments

```bash
# ListComments
$ curl -H {Token} http://localhost:8080/v1/bookmarks/{bookmark_id}/comments | jq -r .

# GetComment
$ curl -H {Token} http://localhost:8080/v1/comments/{comment_id} | jq -r .

# CreateComment
$ curl -X POST -H {Token} http://localhost:8080/v1/bookmarks/{bookmark_id}/comments -d '{"body":"comment"}' -v | jq -r .

# DeleteComment
$ curl -X DELETE -H {Token} http://localhost:8080/v1/comments/{comment_id} | jq -r .
```
