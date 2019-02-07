package handler

import (
	"fmt"
	"strconv"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/types"
)

// ListUser list user struct
type ListUser struct {
	Limit             *int             `json:"limit"`
	ExclusiveStartKey *string          `json:"exclusiveStartKey"`
	Option            *UserQueryOption `json:"option"`
}

// UserQueryOption user query option struct
type UserQueryOption struct {
	UserID *string `json:"userId"`
}

// ListUserHandle List User Handle
func ListUserHandle(arg ListUser, identity types.Identity) (UserConnection, error) {

	svc, err := model.GetSVC()

	if err != nil {
		return UserConnection{}, err
	}

	limit := int64(10)
	if arg.Limit != nil {
		limit = int64(*arg.Limit)
	}

	var userList model.ScanUserListResult
	userList, err = model.ScanUserList(svc, limit, arg.ExclusiveStartKey)

	if err != nil {
		fmt.Println("Got error calling ListUserHandle:")
		fmt.Println(err.Error())
		return UserConnection{}, err
	}

	items := []User{}

	for _, i := range userList.Items {
		item := User{}

		item.ID = i.ID
		item.AvatarURI = i.AvatarURI
		item.Career = i.Career
		item.DisplayName = i.DisplayName
		item.Email = i.Email
		item.Message = i.Message
		item.SkillList = i.SkillList
		createdAt, _ := strconv.Atoi(i.CreatedAt)
		item.CreatedAt = createdAt

		items = append(items, item)
	}

	return UserConnection{items, userList.ExclusiveStartKey}, nil
}
