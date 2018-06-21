package user

import (
	"fmt"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/types"
	"golang.org/x/sync/errgroup"
)

// DeleteUserHandle Delete User Handle
func DeleteUserHandle(arg UserDelete, identity types.Identity) (User, error) {
	eg := errgroup.Group{}

	eg.Go(func() error { return model.DeleteUserByID(arg.ID) })
	eg.Go(func() error { return model.DeleteWorkByUserID(arg.ID) })

	if err := eg.Wait(); err != nil {
		fmt.Println("Got error calling DeleteUserHandle:")
		fmt.Println(err.Error())
		return User{}, err
	}

	return User{ID: arg.ID}, nil
}