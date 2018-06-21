package user

import (
	"fmt"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/define"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
)

// UpdateWorkHandle Update Work Handle
func UpdateWorkHandle(arg UserUpdate, identity define.Identity) (User, error) {

	err := model.UpdateUserByID(model.UserUpdate{
		ID:        identity.Sub,
		Email:     arg.User.Email,
		Name:      arg.User.Name,
		Career:    arg.User.Career,
		AvatarURI: arg.User.AvatarURI,
		Message:   arg.User.Message,
	})

	if err != nil {
		fmt.Println("Got error calling UpdateWorkHandle:")
		fmt.Println(err.Error())
		return User{}, err
	}

	user, err := model.GetUserByID(identity.Sub)

	result := User{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Career:    user.Career,
		AvatarURI: user.AvatarURI,
		Message:   user.Message,
	}

	return result, nil
}
