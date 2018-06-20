package user

import (
	"fmt"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/define"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
)

// UpdateUser type
type UpdateUser struct {
	User User `json:"User"`
}

// UpdateWorkHandle Update Work Handle
func UpdateWorkHandle(arg UpdateUser, identity define.Identity) (interface{}, error) {

	err := model.UpdateUserByID(
		&identity.Sub,
		&arg.User.Email,
		&arg.User.Name,
		&arg.User.Career,
		&arg.User.AvatarURI,
		&arg.User.Message,
	)

	if err != nil {
		fmt.Println("Got error calling UpdateItem:")
		fmt.Println(err.Error())
		return nil, err
	}

	return arg.User, nil
}
