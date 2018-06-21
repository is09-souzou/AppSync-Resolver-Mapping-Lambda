package work

import (
	"fmt"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/define"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
)

// DeleteWorkHandle Delete User Handle
func DeleteWorkHandle(arg WorkDelete, identity define.Identity) (Work, error) {

	err := model.DeleteWorkByID(arg.ID)

	if err != nil {
		fmt.Println("Got error calling DeleteUserHandle:")
		fmt.Println(err.Error())
		return Work{}, err
	}

	return Work{ID: arg.ID}, nil
}