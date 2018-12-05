package model

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// GetSVC get DynamoDB SVC
func GetSVC() (*dynamodb.DynamoDB, error) {
	session, err := session.NewSession(
		&aws.Config{Region: aws.String("ap-northeast-1")},
	)

	if err != nil {
		return nil, err
	}

	svc := dynamodb.New(session)

	return svc, nil
}

// User DynamoDB Resut User Struct
type User struct {
	ID          string
	Email       *string
	DisplayName string
	Career      *string
	AvatarURI   *string
	Message     *string
	SkillList   []string
}

// UserCreate DynamoDB Create User Struct
type UserCreate struct {
	ID          string
	Email       *string
	DisplayName string
	Career      *string
	AvatarURI   *string
	Message     *string
	SkillList   []string
}

// UserUpdate DynamoDB Update User Struct
type UserUpdate struct {
	ID          string
	Email       *string
	DisplayName *string
	Career      *string
	AvatarURI   *string
	Message     *string
	SkillList   *[]string
}

// Work DynamoDB Result Work Struct
type Work struct {
	ID          string
	UserID      string
	Title       string
	Tags        *[]string
	ImageURL    *string
	Description string
	IsPublic    bool
	CreatedAt   string
}

// WorkCreate DynamoDB Create Work Struct
type WorkCreate struct {
	ID          string
	UserID      string
	Title       string
	Tags        *[]string
	ImageURL    *string
	Description string
	IsPublic    bool
	CreatedAt   string
}

// WorkUpdate DynamoDB Create Work Struct
type WorkUpdate struct {
	ID          string
	UserID      *string
	Title       *string
	Tags        *[]string
	ImageURL    *string
	Description *string
	IsPublic    *bool
	CreatedAt   *string
}
