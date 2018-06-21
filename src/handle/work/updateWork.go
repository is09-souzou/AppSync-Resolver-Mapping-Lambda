package work

import (
	"errors"
	"fmt"
	"time"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/define"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
)

// UpdateWorkHandle Update Work Handle
// Only the principal can be update
func UpdateWorkHandle(arg WorkUpdate, identity define.Identity) (Work, error) {

	if (arg.Work.UserID != identity.Sub) {
		return Work{}, errors.New("Only the created user can be update")
	}

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
		return Work{}, err
	}

	// TODO input result value
	return Work{}, nil
}
