package main

import (
	"fmt"
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/handle/user"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/handle/work"
)

type payload struct {
	Field     string          `json:"field"`
	Arguments json.RawMessage `json:"arguments"`
	Sub       string          `json:"identity"`
}

func router(payload payload) (interface{}, error) {
	fmt.Print("==== debug point ====")
	fmt.Print("+%v", payload.Sub)
	fmt.Print("==== debug point ====")
	switch payload.Field {
	case "createUser":
		var p user.CreateUser
		json.Unmarshal(payload.Arguments, &p)
		return user.CreateUserHandle(p, payload.Sub)
	case "deleteUser":
		var p user.DeleteUser
		json.Unmarshal(payload.Arguments, &p)
		return user.DeleteUserHandle(p, payload.Sub)
	case "createWork":
		var p work.CreateWork
		json.Unmarshal(payload.Arguments, &p)
		return work.CreateWorkHandle(p, payload.Sub)
	case "updateWork":
		var p work.UpdateWork
		json.Unmarshal(payload.Arguments, &p)
		return work.UpdateWorkHandle(p, payload.Sub)
	}
	return nil, errors.New("field is not found")
}

func main() {
	lambda.Start(router)
}
