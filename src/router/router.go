package router

import (
	"encoding/json"
	"errors"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/define"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/handle/user"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/handle/work"
)

// Router Routing By Field
func Router(payload define.Payload) (interface{}, error) {
	switch payload.Field {
	case "createUser":
		var p user.UserCreate
		json.Unmarshal(payload.Arguments, &p)
		return user.CreateUserHandle(p, payload.Identity)
	case "deleteUser":
		var p user.UserDelete
		json.Unmarshal(payload.Arguments, &p)
		return user.DeleteUserHandle(p, payload.Identity)
	case "createWork":
		var p work.WorkCreate
		json.Unmarshal(payload.Arguments, &p)
		return work.CreateWorkHandle(p, payload.Identity)
	case "updateWork":
		var p work.WorkUpdate
		json.Unmarshal(payload.Arguments, &p)
		return work.UpdateWorkHandle(p, payload.Identity)
	}
	return nil, errors.New("field is not found")
}
