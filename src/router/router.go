package router

import (
	"encoding/json"
	"errors"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/handler"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/types"
)

// Router Routing By Field
func Router(payload types.Payload) (interface{}, error) {
	switch payload.Field {
	// GraphQL Queries
	case "listWorks":
		var p handler.ListWork
		json.Unmarshal(payload.Arguments, &p)
		return handler.ListWorkHandle(p, payload.Identity)
	// GraphQL Mutations
	case "createUser":
		var p handler.UserCreate
		json.Unmarshal(payload.Arguments, &p)
		return handler.CreateUserHandle(p, payload.Identity)
	case "deleteUser":
		var p handler.UserDelete
		json.Unmarshal(payload.Arguments, &p)
		return handler.DeleteUserHandle(p, payload.Identity)
	case "updateUser":
		var p handler.UserUpdate
		json.Unmarshal(payload.Arguments, &p)
		return handler.UpdateUserHandle(p, payload.Identity)
	case "createWork":
		var p handler.WorkCreate
		json.Unmarshal(payload.Arguments, &p)
		return handler.CreateWorkHandle(p, payload.Identity)
	case "deleteWork":
		var p handler.WorkDelete
		json.Unmarshal(payload.Arguments, &p)
		return handler.DeleteWorkHandle(p, payload.Identity)
	case "updateWork":
		var p handler.WorkUpdate
		json.Unmarshal(payload.Arguments, &p)
		return handler.UpdateWorkHandle(p, payload.Identity)
	}
	return nil, errors.New("field is not found")
}
