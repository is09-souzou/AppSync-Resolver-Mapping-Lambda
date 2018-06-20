package work

import (
	"fmt"
	"time"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/define"

	"github.com/google/uuid"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
)

// CreateWorkHandle Create Work Handle
func CreateWorkHandle(arg WorkCreate, identity define.Identity) (WorkResult, error) {

	uuid, err := uuid.NewUUID()

	if err != nil {
		return WorkResult{}, err
	}

	id := uuid.String()
	createdAt := int(time.Now().Unix())

	if err := model.CreateWork(model.WorkCreate{
		ID:          id,
		UserID:      arg.Work.UserID,
		Title:       arg.Work.Title,
		Tags:        arg.Work.Tags,
		ImageURI:    arg.Work.ImageURI,
		Description: arg.Work.Description,
		CreatedAt:   createdAt,
	}); err != nil {
		fmt.Println("Got error calling CreateWorkHandle:")
		fmt.Println(err.Error())
		return WorkResult{}, err
	}

	return WorkResult{}, nil
}
