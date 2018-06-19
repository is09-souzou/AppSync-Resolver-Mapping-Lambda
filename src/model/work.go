package model

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Work struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// GetWorkByID Get work by ID from DynamoDB
func GetWorkByID(id string) (Work, error) {

	svc, err := getSVC()

	if err != nil {
		return Work{}, err
	}

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("portal-works"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return Work{}, err
	}

	item := Work{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		return Work{}, err
	}

	return item, nil
}

// GetWorkList Get work list By ID from DynamoDB
func GetWorkList() ([]Work, error) {

	svc, err := getSVC()

	if err != nil {
		return []Work{}, err
	}

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("portal-works"),
	})

	item := []Work{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		return []Work{}, err
	}

	return item, nil
}
