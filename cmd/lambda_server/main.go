package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/tabo-syu/bookmarks/infrastructures"
	"github.com/tabo-syu/bookmarks/sqlc"
)

var ginLambda *ginadapter.GinLambda

func init() {
	log.Printf("Gin cold start")
	db, err := infrastructures.NewSQLHandler()
	if err != nil {
		log.Fatal("DB connection failed", err)
	}
	sqlc := sqlc.New(db)

	server := infrastructures.NewServer(sqlc)
	ginLambda = ginadapter.New(server)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
