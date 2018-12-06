package handler

import (
	"errors"
	"fmt"
	"time"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/types"

	"github.com/google/uuid"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/model"
)

// WorkCreateBase WorkCreate work struct
type WorkCreateBase struct {
	UserID      string    `json:"userId"`
	Title       string    `json:"title"`
	Tags        *[]string `json:"tags"`
	ImageURL    *string   `json:"imageUrl"`
	IsPublic    bool      `json:"isPublic"`
	Description string    `json:"description"`
}

// WorkCreate create work struct
type WorkCreate struct {
	Work WorkCreateBase `json:"work"`
}

// CreateWorkHandle Create Work Handle
// Can create only oneself
func CreateWorkHandle(arg WorkCreate, identity types.Identity) (Work, error) {
	if arg.Work.UserID != identity.Sub {
		return Work{}, errors.New("Can create only oneself")
	}

	svc, err := model.GetSVC()

	if err != nil {
		return Work{}, err
	}

	user, err := model.GetUserByID(svc, identity.Sub)

	if err != nil {
		return Work{}, err
	}

	if user.ID == "" {
		return Work{}, errors.New("User is not existed")
	}

	uuid, err := uuid.NewUUID()

	if err != nil {
		return Work{}, err
	}

	id := uuid.String()
	createdAt := time.Now().Unix()

	if err := model.CreateWork(
		svc,
		model.WorkCreate{
			ID:          id,
			UserID:      arg.Work.UserID,
			Title:       arg.Work.Title,
			Tags:        arg.Work.Tags,
			ImageURL:    arg.Work.ImageURL,
			Description: arg.Work.Description,
			IsPublic:    arg.Work.IsPublic,
			CreatedAt:   fmt.Sprint(createdAt),
		},
	); err != nil {
		fmt.Println("Got error calling CreateWorkHandle:")
		fmt.Println(err.Error())
		return Work{}, err
	}

	result := Work{
		ID:          id,
		UserID:      arg.Work.UserID,
		Title:       arg.Work.Title,
		Tags:        arg.Work.Tags,
		ImageURL:    arg.Work.ImageURL,
		Description: arg.Work.Description,
		IsPublic:    arg.Work.IsPublic,
		CreatedAt:   int(createdAt),
	}

	return result, nil
}
