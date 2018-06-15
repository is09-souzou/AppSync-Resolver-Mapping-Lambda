package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/handle"
	"encoding/json"
)

type payload struct {
	Field     string `json:"field"`
	Arguments json.RawMessage `json:"arguments"`
}

func router(payload payload) {
	switch payload.Field {
		case "deleteUser":
			var p handle.DeleteUser
			json.Unmarshal(payload.Arguments, &p)
			handle.HandleRequest(p)
	}
}

func main() {
	lambda.Start(router)
}
