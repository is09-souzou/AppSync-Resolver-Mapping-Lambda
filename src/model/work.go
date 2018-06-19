package model

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Work struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// GetWorkByID Create Work Handle
func GetWorkByID(id string) (Work, error) {

	session, err := session.NewSession(
		&aws.Config{Region: aws.String("ap-northeast-1")},
	)

	if err != nil {
		return Work{}, err
	}

	svc := dynamodb.New(session)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("portal-users"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	item := Work{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		return Work{}, err
	}

	return item, nil
}
