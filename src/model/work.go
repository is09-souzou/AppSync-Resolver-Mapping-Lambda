package model

import (
	"errors"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// WorkTableName DynamoDB Work Table Name
const WorkTableName = "portal-works"

// CreateWork Create work to DynamoDB
func CreateWork(work WorkCreate) error {

	svc, err := getSVC()

	if err != nil {
		return err
	}

	var item = map[string]*dynamodb.AttributeValue{}

	item["id"].S = aws.String(work.ID)
	item["userId"].S = aws.String(work.UserID)
	item["title"].S = aws.String(work.Title)
	item["imageUri"].S = aws.String(work.ImageURI)
	item["description"].S = aws.String(work.Description)
	item["createdAt"].N = aws.String(strconv.Itoa(work.CreatedAt))

	if work.Tags != nil {
		item["tags"].SS = aws.StringSlice(*work.Tags)
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(WorkTableName),
	}

	_, err = svc.PutItem(input)

	if err != nil {
		return err
	}

	return nil
}

// GetWorkByID Get work by ID from DynamoDB
func GetWorkByID(id string) (Work, error) {

	svc, err := getSVC()

	if err != nil {
		return Work{}, err
	}

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(WorkTableName),
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
		TableName: aws.String(WorkTableName),
	})

	item := []Work{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		return []Work{}, err
	}

	return item, nil
}

// UpdateWorkByID Update work By ID to DynamoDB
func UpdateWorkByID(work WorkUpdate) error {

	svc, err := getSVC()

	if err != nil {
		return err
	}

	if work.UserID == nil && work.Title == nil && work.Tags == nil && work.ImageURI == nil && work.Description == nil && work.CreatedAt == nil {
		return errors.New("required new value")
	}

	var expressionAttributeValues = map[string]*dynamodb.AttributeValue{}


	if work.UserID != nil {
		expressionAttributeValues["userId"].S = aws.String(*work.UserID)
	}

	if work.Title != nil {
		expressionAttributeValues["title"].S = aws.String(*work.Title)
	}

	if work.Tags != nil {
		expressionAttributeValues["tags"].SS = aws.StringSlice(*work.Tags)
	}

	if work.ImageURI != nil {
		expressionAttributeValues["imageUri"].S = aws.String(*work.ImageURI)
	}

	if work.Description != nil {
		expressionAttributeValues["description"].S = aws.String(*work.Description)
	}

	if work.CreatedAt != nil {
		expressionAttributeValues["createdAt"].N = aws.String(strconv.Itoa(*work.CreatedAt))
	}

	input := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(WorkTableName),
		ExpressionAttributeValues: expressionAttributeValues,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(work.ID),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String(""),
	}

	_, err = svc.UpdateItem(input)

	if err != nil {
		return err
	}

	return nil
}
