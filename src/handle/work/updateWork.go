package work

import (
	"time"
	"fmt"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/define"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
)

// UpdateWorkHandle Update Work Handle
func UpdateWorkHandle(arg WorkUpdate, identity define.Identity) (interface{}, error) {

	createdAt := int(time.Now().Unix())

	err := model.UpdateWorkByID(
		&arg.Work.ID,
		&arg.Work.UserID,
		arg.Work.Title,
		arg.Work.Tags,
		arg.Work.ImageURI,
		arg.Work.Description,
		&createdAt,
	)

	if err != nil {
		fmt.Println("Got error calling UpdateItem:")
		fmt.Println(err.Error())
		return nil, err
	}

	return arg.Work, nil
}
