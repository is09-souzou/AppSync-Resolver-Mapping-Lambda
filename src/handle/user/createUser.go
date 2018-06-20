package user

import (
	"fmt"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/define"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
)

// CreateUserHandle Create User Handle
func CreateUserHandle(arg UserCreate, identity define.Identity) (UserResult, error) {

	err := model.CreateUser(model.UserCreate{
		ID:        identity.Sub,
		Email:     arg.User.Email,
		Name:      arg.User.Name,
		Career:    arg.User.Career,
		AvatarURI: arg.User.AvatarURI,
		Message:   arg.User.Message,
	})

	if err != nil {
		fmt.Println("Got error calling CreateUserHandle:")
		fmt.Println(err.Error())
		return UserResult{}, err
	}

	// TODO input result value
	return UserResult{}, nil
}
