package work

import (
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

	if err := model.CreateWork(
		&id,
		&arg.Work.UserID,
		&arg.Work.Title,
		arg.Work.Tags,
		&arg.Work.ImageURI,
		&arg.Work.Description,
		&createdAt,
	); err != nil {
		return WorkResult{}, err
	}

	return WorkResult{}, nil
}
