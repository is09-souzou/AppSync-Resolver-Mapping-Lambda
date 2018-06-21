package user

import (
	"errors"
	"fmt"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/types"
)

// UpdateUserHandle Update User Handle
// Can update only oneself
func UpdateUserHandle(arg UserUpdate, identity types.Identity) (User, error) {

	if arg.User.ID != identity.Sub {
		return User{}, errors.New("Can update only oneself")
	}

	newUser, err := model.UpdateUserByID(model.UserUpdate{
		ID:          identity.Sub,
		Email:       arg.User.Email,
		DisplayName: arg.User.DisplayName,
		Career:      arg.User.Career,
		AvatarURI:   arg.User.AvatarURI,
		Message:     arg.User.Message,
	})

	if err != nil {
		fmt.Println("Got error calling UpdateUserHandle:")
		fmt.Println(err.Error())
		return User{}, err
	}

	result := User{
		ID:          newUser.ID,
		Email:       newUser.Email,
		DisplayName: newUser.DisplayName,
		Career:      newUser.Career,
		AvatarURI:   newUser.AvatarURI,
		Message:     newUser.Message,
	}

	return result, nil
}
