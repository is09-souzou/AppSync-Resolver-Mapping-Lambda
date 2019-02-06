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
	ImageURL    *string   `json:"imageUrl"`
	IsPublic    *bool     `json:"isPublic"`
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

	work, err := model.GetWorkByID(svc, arg.Work.ID, false)

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
			ImageURL:    arg.Work.ImageURL,
			IsPublic:    arg.Work.IsPublic,
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

	// Delete or Decrement popular tags of oldWork
	if work.Tags != nil {
		for _, i := range *work.Tags {
			tag, err := model.GetPopularTagByName(svc, i)
			if err == nil && tag.Name != "" {
				result, err := model.UpdatePopularTagByName(svc, tag, "-1")
				if err != nil {
					fmt.Println("Got error calling UpdatePopularTag:")
					fmt.Println(err.Error())
					return Work{}, err
				}
				if result.Count == 0 {
					err := model.DeletePopularTagByName(svc, result.Name)
					if err != nil {
						fmt.Println("Got error calling DeletePopularTag:")
						fmt.Println(err.Error())
						return Work{}, err
					}
				}
			}
		}
	}

	result := Work{
		ID:          newWork.ID,
		UserID:      newWork.UserID,
		Title:       newWork.Title,
		Tags:        newWork.Tags,
		ImageURL:    newWork.ImageURL,
		Description: newWork.Description,
		IsPublic:    newWork.IsPublic,
		CreatedAt:   newWrokCreatedAt,
	}

	// Add or Inclement popular tags of newWork
	if newWork.Tags != nil {
		for _, i := range *newWork.Tags {
			tag, err := model.GetPopularTagByName(svc, i)
			if err != nil || tag.Name == "" {
				if err := model.CreatePopularTag(
					svc,
					model.PopularTag{
						Name:  i,
						Count: 1,
					},
				); err != nil {
					fmt.Println("Got error calling CreatePopularTag:")
					fmt.Println(err.Error())
					return Work{}, err
				}
			} else {
				_, err := model.UpdatePopularTagByName(svc, tag, "1")
				if err != nil {
					fmt.Println("Got error calling UpdatePopularTag:")
					fmt.Println(err.Error())
					return Work{}, err
				}
			}
		}
	}

	return result, nil
}
