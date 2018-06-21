package model

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func getSVC() (*dynamodb.DynamoDB, error) {
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
	ID        string
	Email     string
	Name      string
	Career    *string
	AvatarURI *string
	Message   *string
}

// UserCreate DynamoDB Create User Struct
type UserCreate struct {
	ID        string
	Email     string
	Name      string
	Career    *string
	AvatarURI *string
	Message   *string
}

// UserUpdate DynamoDB Update User Struct
type UserUpdate struct {
	ID        string
	Email     *string
	Name      *string
	Career    *string
	AvatarURI *string
	Message   *string
}

// Work DynamoDB Result Work Struct
type Work struct {
	ID          string
	UserID      string
	Title       string
	Tags        []string
	ImageURI    string
	Description string
	CreatedAt   string
}

// WorkCreate DynamoDB Create Work Struct
type WorkCreate struct {
	ID          string
	UserID      string
	Title       string
	Tags        *[]string
	ImageURI    string
	Description string
	CreatedAt   string
}

// WorkUpdate DynamoDB Create Work Struct
type WorkUpdate struct {
	ID          string
	UserID      *string
	Title       *string
	Tags        *[]string
	ImageURI    *string
	Description *string
	CreatedAt   *string
}
