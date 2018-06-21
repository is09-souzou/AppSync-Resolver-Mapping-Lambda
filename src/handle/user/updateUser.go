package user

import (
	"errors"
	"fmt"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/define"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
)

// UpdateWorkHandle Update Work Handle
// Can update only oneself
func UpdateWorkHandle(arg UserUpdate, identity define.Identity) (User, error) {

	if (arg.User.ID != identity.Sub) {
		return User{}, errors.New("Can update only oneself")
	}

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
