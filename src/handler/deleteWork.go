package handler

import (
	"errors"
	"fmt"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/types"
)

// WorkDelete delete work struct
type WorkDelete struct {
	ID string `json:"id"`
}

// DeleteWorkHandle Delete Work Handle
// Can delete only oneself
func DeleteWorkHandle(arg WorkDelete, identity types.Identity) (Work, error) {

	svc, err := model.GetSVC()

	if err != nil {
		return Work{}, err
	}

	work, err := model.GetWorkByID(svc, arg.ID, false)

	if err != nil {
		return Work{}, err
	}

	if work.UserID != identity.Sub {
		return Work{}, errors.New("Can delete only oneself")
	}

	err = model.DeleteWorkByID(svc, arg.ID)

	if err != nil {
		fmt.Println("Got error calling DeleteUserHandle:")
		fmt.Println(err.Error())
		return Work{}, err
	}

	// Delete or Decrement popular tags
	if work.Tags != nil {
		for _, i := range *work.Tags {
			tag, err := model.GetPopularTagByName(svc, i)
			if err == nil && tag.Name != "" {
				result, err := model.UpdatePopularTagByName(svc, tag, "-1")
				if err != nil {
					fmt.Println("Got error calling UpdatePopularTag:")
					fmt.Println(err.Error())
					return Work{}, err
				}
				if result.Count == 0 {
					err := model.DeletePopularTagByName(svc, result.Name)
					if err != nil {
						fmt.Println("Got error calling DeletePopularTag:")
						fmt.Println(err.Error())
						return Work{}, err
					}
				}
			}
		}
	}

	return Work{ID: arg.ID}, nil
}
