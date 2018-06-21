package work

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/define"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
)

// UpdateWorkHandle Update Work Handle
// Only the principal can be update
func UpdateWorkHandle(arg WorkUpdate, identity define.Identity) (Work, error) {

	work, err := model.GetWorkByID(arg.Work.ID)

	if err != nil {
		fmt.Println("Got error calling UpdateWorkHandle:")
		fmt.Println(err.Error())
		return Work{}, err
	}

	if work.UserID != identity.Sub {
		return Work{}, errors.New("Only the created user can be update")
	}

	if err := model.UpdateWorkByID(model.WorkUpdate{
		ID:          arg.Work.ID,
		UserID:      nil,
		Title:       arg.Work.Title,
		Tags:        arg.Work.Tags,
		ImageURI:    arg.Work.ImageURI,
		Description: arg.Work.Description,
	}); err != nil {
		fmt.Println("Got error calling UpdateWorkHandle:")
		fmt.Println(err.Error())
		return Work{}, err
	}

	createdAt, _ := strconv.Atoi(work.CreatedAt)

	result := Work{
		ID:          work.ID,
		UserID:      work.UserID,
		Title:       work.Title,
		Tags:        &work.Tags,
		ImageURI:    work.ImageURI,
		Description: work.Description,
		CreatedAt:   createdAt,
	}

	if arg.Work.Title != nil {
		result.Title = *arg.Work.Title
	}

	if arg.Work.Tags != nil {
		result.Tags = arg.Work.Tags
	}

	if arg.Work.ImageURI != nil {
		result.ImageURI = *arg.Work.ImageURI
	}

	if arg.Work.Description != nil {
		result.Description = *arg.Work.Description
	}

	return result, nil
}
