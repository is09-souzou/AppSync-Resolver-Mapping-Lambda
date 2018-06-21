package user

import (
	"fmt"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/types"
)

// CreateUserHandle Create User Handle
func CreateUserHandle(arg UserCreate, identity types.Identity) (User, error) {

	fmt.Print("sub", identity.Sub)

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
		return User{}, err
	}

	result := User{
		ID:        identity.Sub,
		Email:     arg.User.Email,
		Name:      arg.User.Name,
		Career:    arg.User.Career,
		AvatarURI: arg.User.AvatarURI,
		Message:   arg.User.Message,
	}

	return result, nil
}
