package handler

import (
	"fmt"
	"strconv"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/types"
)

// UserConnectionArg user connection argument struct
type UserConnectionArg struct {
	Limit             *int    `json:"limit"`
	ExclusiveStartKey *string `json:"exclusiveStartKey"`
}

// UserConnectionHandle List User Handle
func UserConnectionHandle(arg UserConnectionArg, identity types.Identity) (UserConnection, error) {

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
