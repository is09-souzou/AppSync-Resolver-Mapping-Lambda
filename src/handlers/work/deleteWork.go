package work

import (
	"errors"
	"fmt"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/types"
)

// DeleteWorkHandle Delete User Handle
// Can delete only oneself
func DeleteWorkHandle(arg WorkDelete, identity types.Identity) (Work, error) {
	work, err := model.GetWorkByID(arg.ID)

	if err != nil {
		return Work{}, err
	}

	if work.UserID != identity.Sub {
		return Work{}, errors.New("Can delete only oneself")
	}

	err = model.DeleteWorkByID(arg.ID)

	if err != nil {
		fmt.Println("Got error calling DeleteUserHandle:")
		fmt.Println(err.Error())
		return Work{}, err
	}

	return Work{ID: arg.ID}, nil
}
