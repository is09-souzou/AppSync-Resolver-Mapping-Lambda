package handler

import (
	"strconv"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/types"
)

// ListWork list work struct
type ListWork struct {
	Limit     *int             `json:"limit"`
	NextToken *string          `json:"nextToken"`
	Option    *WorkQueryOption `json:"option"`
}

// WorkQueryOption work query option struct
type WorkQueryOption struct {
	Tags   *[]string `json:"tags"`
	Word   *string   `json:"word"`
	UserID *string   `json:"userId"`
}

// ListWorkHandle List Work Handle
func ListWorkHandle(arg ListWork, identity types.Identity) (WorkConnection, error) {

	svc, err := model.GetSVC()

	if err != nil {
		return WorkConnection{}, err
	}

	workList, err := model.ScanWorkList(svc, int64(10), nil)

	if err != nil {
		return WorkConnection{}, err
	}

	items := []Work{}

	for _, i := range workList.Items {
		item := Work{}

		item.ID = i.ID
		item.UserID = i.UserID
		item.Title = i.Title
		item.Tags = i.Tags
		item.ImageURI = i.ImageURI
		item.Description = i.Description
		createdAt, _ := strconv.Atoi(i.CreatedAt)
		item.CreatedAt = createdAt

		_ = append(items, item)
	}

	return WorkConnection{items, workList.NextToken}, nil
}
