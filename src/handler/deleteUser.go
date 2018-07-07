package handler

import (
	"errors"
	"fmt"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/types"
	"golang.org/x/sync/errgroup"
)

// UserDelete delete user struct
type UserDelete struct {
	ID string `json:"id"`
}

// DeleteUserHandle Delete User Handle
// Can delete only oneself
func DeleteUserHandle(arg UserDelete, identity types.Identity) (User, error) {
	if arg.ID != identity.Sub {
		return User{}, errors.New("Can delete only oneself")
	}

	eg := errgroup.Group{}

	svc, err := model.GetSVC()

	if err != nil {
		return User{}, err
	}

	eg.Go(func() error { return model.DeleteUserByID(svc, arg.ID) })
	eg.Go(func() error { return model.DeleteWorkByUserID(svc, arg.ID) })

	if err := eg.Wait(); err != nil {
		fmt.Println("Got error calling DeleteUserHandle:")
		fmt.Println(err.Error())
		return User{}, err
	}

	return User{ID: arg.ID}, nil
}
