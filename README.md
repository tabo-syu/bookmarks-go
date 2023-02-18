# v1

## bookmarks

```bash
# ListBookmarks
$ curl http://localhost:8080/v1/bookmarks | jq -r .

# GetBookmark
$ curl http://localhost:8080/v1/bookmarks/{bookmark_id} | jq -r .

# CreateBookmark
$ curl -X POST http://localhost:8080/v1/bookmarks -d '{"url":"https://tabo-syu.com","title":"bookmark title", "description": "bookmark desc"}' -v | jq -r .

# UpdateBookmark
$ curl -X PUT http://localhost:8080/v1/bookmarks/{bookmark_id} -d '{"url":"https://tabo-syu.com","title":"bookmark title", "description": "bookmark desc"}' -v | jq -r .

# DeleteBookmark
$ curl -X DELETE http://localhost:8080/v1/bookmarks/{bookmark_id} | jq -r .
```

## tags

```bash
# ListTags
$ curl http://localhost:8080/v1/tags | jq -r .

# GetTag
$ curl http://localhost:8080/v1/tags/{tag_id} | jq -r .

# CreateTag
$ curl -X POST http://localhost:8080/v1/tags -d '{"name":"tag name","color":"#123456"}' -v | jq -r .

# UpdateTag
$ curl -X PUT http://localhost:8080/v1/tags/{tag_id} -d '{"name":"tag name","color":"#123456"}' -v | jq -r .

# DeleteTag
$ curl -X DELETE http://localhost:8080/v1/tags/{tag_id} | jq -r .
```

## comments

```bash
# ListComments
$ curl http://localhost:8080/v1/bookmarks/{bookmark_id}/comments | jq -r .

# GetComment
$ curl http://localhost:8080/v1/comments/{comment_id} | jq -r .

# CreateComment
$ curl -X POST http://localhost:8080/v1/bookmarks/{bookmark_id} -d '{"body":"comment"}' -v | jq -r .

# DeleteComment
$ curl -X DELETE http://localhost:8080/v1/comments/{comment_id} | jq -r .
```