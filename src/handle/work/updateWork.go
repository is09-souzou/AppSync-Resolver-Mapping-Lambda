package work

import (
	"fmt"
	"time"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/define"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
)

// UpdateWorkHandle Update Work Handle
func UpdateWorkHandle(arg WorkUpdate, identity define.Identity) (WorkResult, error) {

	createdAt := int(time.Now().Unix())

	err := model.UpdateWorkByID(model.WorkUpdate{
		ID:          arg.Work.ID,
		UserID:      nil,
		Title:       arg.Work.Title,
		Tags:        arg.Work.Tags,
		ImageURI:    arg.Work.ImageURI,
		Description: arg.Work.Description,
		CreatedAt:   &createdAt,
	})

	if err != nil {
		fmt.Println("Got error calling UpdateWorkHandle:")
		fmt.Println(err.Error())
		return WorkResult{}, err
	}

	// TODO input result value
	return WorkResult{}, nil
}
