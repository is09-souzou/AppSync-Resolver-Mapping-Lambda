package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/router"
)

func main() {
	lambda.Start(router.Router)
}
