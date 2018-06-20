package user

import (
	"fmt"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/define"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
)

// DeleteUserHandle Delete User Handle
func DeleteUserHandle(arg UserDelete, identity define.Identity) (UserResult, error) {

	err := model.DeleteUserByID(arg.ID)

	if err != nil {
		fmt.Println("Got error calling DeleteUserHandle:")
		fmt.Println(err.Error())
		return UserResult{}, err
	}

	// TODO input result value
	return UserResult{}, nil
}
