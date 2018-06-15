package main

import (
	"errors"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/handle"
	"encoding/json"
)

type payload struct {
	Field     string `json:"field"`
	Arguments json.RawMessage `json:"arguments"`
}

func router(payload payload) (interface{}, error) {
	switch payload.Field {
		case "deleteUser":
			var p handle.DeleteUser
			json.Unmarshal(payload.Arguments, &p)
			return handle.HandleRequest(p)
	}
	return nil, errors.New("field is not found")
}

func main() {
	lambda.Start(router)
}
