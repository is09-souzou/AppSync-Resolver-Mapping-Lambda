// https://github.com/sbstjn/go-appsync-graphql-cloudformation/blob/master/main.go
package main

import (
  "log"
	"time"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Post struct {
	URL   string    `json:"url"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}

func handleRequest(req events.APIGatewayProxyRequest) (interface{}, error) {

  // list := []Post{}
  
  // list = append(list, Post{"test", "test2", "test3"})
  // list = append(list, Post{"test", "test2", "test3"})
  log.Print(req)

	return nil, nil
}

func main() {
	lambda.Start(handleRequest)
}
