package handler

import (
	"errors"
	"fmt"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/types"
)

// UserUpdateBase UserUpdate user struct
type UserUpdateBase struct {
	ID               string    `json:"id"`
	Email            *string   `json:"email"`
	DisplayName      *string   `json:"displayName"`
	Career           *string   `json:"career"`
	AvatarURI        *string   `json:"avatarUri"`
	Message          *string   `json:"message"`
	SkillList        *[]string `json:"skillList"`
	FavoriteWorkList *[]string `json:"favoriteWorkList"`
}

// UserUpdate create user struct
type UserUpdate struct {
	User UserUpdateBase `json:"user"`
}

// UpdateUserHandle Update User Handle
// Can update only oneself
func UpdateUserHandle(arg UserUpdate, identity types.Identity) (User, error) {

	svc, err := model.GetSVC()

	if err != nil {
		return User{}, err
	}

	if arg.User.ID != identity.Sub {
		return User{}, errors.New("Can update only oneself")
	}

	newUser, err := model.UpdateUserByID(
		svc,
		model.UserUpdate{
			ID:               identity.Sub,
			Email:            arg.User.Email,
			DisplayName:      arg.User.DisplayName,
			Career:           arg.User.Career,
			AvatarURI:        arg.User.AvatarURI,
			Message:          arg.User.Message,
			SkillList:        arg.User.SkillList,
			FavoriteWorkList: arg.User.FavoriteWorkList,
		},
	)

	if err != nil {
		fmt.Println("Got error calling UpdateUserHandle:")
		fmt.Println(err.Error())
		return User{}, err
	}

	result := User{
		ID:               newUser.ID,
		Email:            newUser.Email,
		DisplayName:      newUser.DisplayName,
		Career:           newUser.Career,
		AvatarURI:        newUser.AvatarURI,
		Message:          newUser.Message,
		SkillList:        newUser.SkillList,
		FavoriteWorkList: newUser.FavoriteWorkList,
	}

	return result, nil
}
