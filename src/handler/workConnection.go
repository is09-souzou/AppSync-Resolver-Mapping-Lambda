package handler

import (
	"fmt"
	"strconv"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/types"
)

// WorkConnectionArg work connection argument struct
type WorkConnectionArg struct {
	Limit             *int    `json:"limit"`
	ExclusiveStartKey *string `json:"exclusiveStartKey"`
	UserID            string  `json:"userId"`
}

// WorkConnectionHandle List Work Handle
func WorkConnectionHandle(arg WorkConnectionArg, identity types.Identity) (WorkConnection, error) {

	svc, err := model.GetSVC()

	if err != nil {
		return WorkConnection{}, err
	}

	limit := int64(10)
	if arg.Limit != nil {
		limit = int64(*arg.Limit)
	}

	var publicOnly = true
	if arg.UserID == identity.Sub {
		publicOnly = false
	}

	var workList model.ScanWorkListResult
	workList, err = model.ScanWorkListByUserID(svc, limit, arg.ExclusiveStartKey, arg.UserID, publicOnly)

	if err != nil {
		fmt.Println("Got error calling ListWorkHandle:")
		fmt.Println(err.Error())
		return WorkConnection{}, err
	}

	items := []Work{}

	for _, i := range workList.Items {
		item := Work{}

		item.ID = i.ID
		item.UserID = i.UserID
		item.Title = i.Title
		item.Tags = i.Tags
		item.ImageURL = i.ImageURL
		item.Description = i.Description
		createdAt, _ := strconv.Atoi(i.CreatedAt)
		item.CreatedAt = createdAt

		items = append(items, item)
	}

	return WorkConnection{items, workList.ExclusiveStartKey}, nil
}
