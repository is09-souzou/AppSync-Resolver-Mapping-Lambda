package handler

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/types"
)

// WorkUpdateBase WorkUpdate work struct
type WorkUpdateBase struct {
	ID          string    `json:"id"`
	UserID      *string   `json:"userId"`
	Title       *string   `json:"title"`
	Tags        *[]string `json:"tags"`
	ImageURIs   *[]string `json:"imageUris"`
	Description *string   `json:"description"`
}

// WorkUpdate update work struct
type WorkUpdate struct {
	Work WorkUpdateBase `json:"work"`
}

// UpdateWorkHandle Update Work Handle
// Only the principal can be update
func UpdateWorkHandle(arg WorkUpdate, identity types.Identity) (Work, error) {

	svc, err := model.GetSVC()

	if err != nil {
		return Work{}, err
	}

	work, err := model.GetWorkByID(svc, arg.Work.ID)

	if err != nil {
		fmt.Println("Got error calling UpdateWorkHandle:")
		fmt.Println(err.Error())
		return Work{}, err
	}

	if work.UserID != identity.Sub {
		return Work{}, errors.New("Only the created user can be update")
	}

	newWork, err := model.UpdateWorkByID(
		svc,
		model.WorkUpdate{
			ID:          arg.Work.ID,
			UserID:      nil,
			Title:       arg.Work.Title,
			Tags:        arg.Work.Tags,
			ImageURIs:   arg.Work.ImageURIs,
			Description: arg.Work.Description,
		},
	)

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
		ImageURIs:   newWork.ImageURIs,
		Description: newWork.Description,
		CreatedAt:   newWrokCreatedAt,
	}

	return result, nil
}
