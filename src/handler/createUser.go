package handler

import (
	"fmt"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/types"
)

// UserCreateBase UserCreate user struct
type UserCreateBase struct {
	Email       *string 	`json:"email"`
	DisplayName string  	`json:"displayName"`
	Career      *string 	`json:"career"`
	AvatarURI   *string 	`json:"avatarUri"`
	Message     *string 	`json:"message"`
	SkillList	[]string	`json:"skillList"`
}

// UserCreate create user struct
type UserCreate struct {
	User UserCreateBase `json:"user"`
}

// CreateUserHandle Create User Handle
func CreateUserHandle(arg UserCreate, identity types.Identity) (User, error) {

	svc, err := model.GetSVC()

	if err != nil {
		return User{}, err
	}

	err = model.CreateUser(
		svc,
		model.UserCreate{
			ID:          identity.Sub,
			Email:       arg.User.Email,
			DisplayName: arg.User.DisplayName,
			Career:      arg.User.Career,
			AvatarURI:   arg.User.AvatarURI,
			Message:     arg.User.Message,
			SkillList:	arg.User.SkillList,
		},
	)

	if err != nil {
		fmt.Println("Got error calling CreateUserHandle:")
		fmt.Println(err.Error())
		return User{}, err
	}

	result := User{
		ID:          identity.Sub,
		Email:       arg.User.Email,
		DisplayName: arg.User.DisplayName,
		Career:      arg.User.Career,
		AvatarURI:   arg.User.AvatarURI,
		Message:     arg.User.Message,
		SkillList:	arg.User.SkillList,
	}

	return result, nil
}
