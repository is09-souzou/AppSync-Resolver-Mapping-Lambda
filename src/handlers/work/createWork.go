package work

import (
	"errors"
	"fmt"
	"time"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/types"

	"github.com/google/uuid"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
)

// CreateWorkHandle Create Work Handle
// Can create only oneself
func CreateWorkHandle(arg WorkCreate, identity types.Identity) (Work, error) {
	if arg.Work.UserID != identity.Sub {
		return Work{}, errors.New("Can create only oneself")
	}

	uuid, err := uuid.NewUUID()

	if err != nil {
		return Work{}, err
	}

	id := uuid.String()
	createdAt := time.Now().Unix()

	if err := model.CreateWork(model.WorkCreate{
		ID:          id,
		UserID:      arg.Work.UserID,
		Title:       arg.Work.Title,
		Tags:        arg.Work.Tags,
		ImageURI:    arg.Work.ImageURI,
		Description: arg.Work.Description,
		CreatedAt:   fmt.Sprint(createdAt),
	}); err != nil {
		fmt.Println("Got error calling CreateWorkHandle:")
		fmt.Println(err.Error())
		return Work{}, err
	}

	result := Work{
		ID:          id,
		UserID:      arg.Work.UserID,
		Title:       arg.Work.Title,
		Tags:        arg.Work.Tags,
		ImageURI:    arg.Work.ImageURI,
		Description: arg.Work.Description,
		CreatedAt:   int(createdAt),
	}

	return result, nil
}
