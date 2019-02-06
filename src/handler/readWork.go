package handler

import (
	"fmt"
	"strconv"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/types"
)

// ReadWork list work struct
type ReadWork struct {
	ID string `json:"id"`
}

// ReadWorkHandle List Work Handle
func ReadWorkHandle(arg ReadWork, identity types.Identity) (Work, error) {

	svc, err := model.GetSVC()

	if err != nil {
		return Work{}, err
	}

	work, err := model.GetWorkByID(svc, arg.ID, false)

	if err != nil {
		fmt.Println("Got error calling ReadWorkHandle:")
		fmt.Println(err.Error())
		return Work{}, err
	}

	if !work.IsPublic && work.UserID != identity.Sub {
		return Work{}, nil
	}

	wrokCreatedAt, err := strconv.Atoi(work.CreatedAt)

	if err != nil {
		fmt.Println("Got error calling ReadWorkHandle:")
		fmt.Println(err.Error())
		return Work{}, err
	}

	result := Work{
		ID:               work.ID,
		UserID:           work.UserID,
		Title:            work.Title,
		Tags:             work.Tags,
		ImageURL:         work.ImageURL,
		Description:      work.Description,
		IsPublic:         work.IsPublic,
		CreatedAt:        wrokCreatedAt,
		FavoriteUserList: work.FavoriteUserList,
	}

	return result, nil
}
