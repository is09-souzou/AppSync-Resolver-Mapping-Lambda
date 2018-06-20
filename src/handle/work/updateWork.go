package work

import (
	"fmt"
	"time"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/define"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
)

// UpdateWork type
type UpdateWork struct {
	Work Work `json:"work"`
}

// UpdateWorkHandle Update Work Handle
func UpdateWorkHandle(arg UpdateWork, identity define.Identity) (interface{}, error) {

	arg.Work.CreatedAt = int(time.Now().Unix())

	err := model.UpdateWorkByID(
		&arg.Work.ID,
		&arg.Work.UserID,
		&arg.Work.Title,
		&arg.Work.Tags,
		&arg.Work.ImageURI,
		&arg.Work.Description,
		&arg.Work.CreatedAt,
	)

	if err != nil {
		fmt.Println("Got error calling UpdateItem:")
		fmt.Println(err.Error())
		return nil, err
	}

	return arg.Work, nil
}
