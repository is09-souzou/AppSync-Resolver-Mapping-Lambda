package work

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/types"
)

// UpdateWorkHandle Update Work Handle
// Only the principal can be update
func UpdateWorkHandle(arg WorkUpdate, identity types.Identity) (Work, error) {

	work, err := model.GetWorkByID(arg.Work.ID)

	if err != nil {
		fmt.Println("Got error calling UpdateWorkHandle:")
		fmt.Println(err.Error())
		return Work{}, err
	}

	if work.UserID != identity.Sub {
		return Work{}, errors.New("Only the created user can be update")
	}

	newWork, err := model.UpdateWorkByID(model.WorkUpdate{
		ID:          arg.Work.ID,
		UserID:      nil,
		Title:       arg.Work.Title,
		Tags:        arg.Work.Tags,
		ImageURI:    arg.Work.ImageURI,
		Description: arg.Work.Description,
	})

	if err != nil {
		fmt.Println("Got error calling UpdateWorkHandle:")
		fmt.Println(err.Error())
		return Work{}, err
	}

	newWrokCreatedAt, err := strconv.Atoi(newWork.CreatedAt)

	if err != nil {
		fmt.Println("Got error calling UpdateWorkHandle:")
		fmt.Println(err.Error())
		return Work{}, err
	}

	result := Work{
		ID:          newWork.ID,
		UserID:      newWork.UserID,
		Title:       newWork.Title,
		Tags:        newWork.Tags,
		ImageURI:    newWork.ImageURI,
		Description: newWork.Description,
		CreatedAt:   newWrokCreatedAt,
	}

	return result, nil
}