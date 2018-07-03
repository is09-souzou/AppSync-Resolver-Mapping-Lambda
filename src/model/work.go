package model

import (
	"errors"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

// WorkTableName DynamoDB Work Table Name
const WorkTableName = "portal-works"

// CreateWork Create work to DynamoDB
func CreateWork(svc *dynamodb.DynamoDB, work WorkCreate) error {

	if work.ID == "" {
		return errors.New("required ID in work")
	}

	if work.UserID == "" {
		work.UserID = " "
	}

	if work.Title == "" {
		work.Title = " "
	}

	if work.UserID == "" {
		work.UserID = " "
	}

	if work.ImageURI == "" {
		work.ImageURI = " "
	}

	if work.Description == "" {
		work.Description = " "
	}

	if work.CreatedAt == "" {
		work.CreatedAt = " "
	}

	var item = map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String(work.ID),
		},
		"userId": {
			S: aws.String(work.UserID),
		},
		"title": {
			S: aws.String(work.Title),
		},
		"imageUri": {
			S: aws.String(work.ImageURI),
		},
		"description": {
			S: aws.String(work.Description),
		},
		"createdAt": {
			S: aws.String(work.CreatedAt),
		},
	}

	if work.Tags != nil {
		item["tags"] = &dynamodb.AttributeValue{SS: aws.StringSlice(*work.Tags)}
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(WorkTableName),
	}

	_, err := svc.PutItem(input)

	if err != nil {
		return err
	}

	return nil
}

// GetWorkByID Get work by ID from DynamoDB
func GetWorkByID(svc *dynamodb.DynamoDB, id string) (Work, error) {

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
func GetWorkList(svc *dynamodb.DynamoDB) ([]Work, error) {

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
func UpdateWorkByID(svc *dynamodb.DynamoDB, work WorkUpdate) (Work, error) {

	if work.UserID == nil && work.Title == nil && work.Tags == nil && work.ImageURI == nil && work.Description == nil && work.CreatedAt == nil {
		return Work{}, errors.New("required new value")
	}

	expressionAttributeValues := map[string]*dynamodb.AttributeValue{}
	updateExpression := "SET "

	if work.UserID != nil {
		if *work.UserID == "" {
			*work.UserID = " "
		}
		expressionAttributeValues[":userId"] = &dynamodb.AttributeValue{S: aws.String(*work.UserID)}
		updateExpression += "userId = :userId, "
	}

	if work.Title != nil {
		if *work.Title == "" {
			*work.Title = " "
		}
		expressionAttributeValues[":title"] = &dynamodb.AttributeValue{S: aws.String(*work.Title)}
		updateExpression += "title = :title, "
	}

	if work.Tags != nil {
		expressionAttributeValues[":tags"] = &dynamodb.AttributeValue{SS: aws.StringSlice(*work.Tags)}
		updateExpression += "tags = :tags, "
	}

	if work.ImageURI != nil {
		if *work.ImageURI == "" {
			*work.ImageURI = " "
		}
		expressionAttributeValues[":imageUri"] = &dynamodb.AttributeValue{S: aws.String(*work.ImageURI)}
		updateExpression += "imageUri = :imageUri, "
	}

	if work.Description != nil {
		if *work.Description == "" {
			*work.Description = " "
		}
		expressionAttributeValues[":description"] = &dynamodb.AttributeValue{S: aws.String(*work.Description)}
		updateExpression += "description = :description, "
	}

	if work.CreatedAt != nil {
		if *work.CreatedAt == "" {
			*work.CreatedAt = " "
		}
		expressionAttributeValues[":createdAt"] = &dynamodb.AttributeValue{N: aws.String(string(*work.CreatedAt))}
		updateExpression += "createdAt = :createdAt, "
	}

	input := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(WorkTableName),
		ExpressionAttributeValues: expressionAttributeValues,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(work.ID),
			},
		},
		ReturnValues:     aws.String("ALL_NEW"),
		UpdateExpression: aws.String(strings.TrimRight(updateExpression, ", ")),
	}

	result, err := svc.UpdateItem(input)

	if err != nil {
		return Work{}, err
	}

	item := Work{}

	err = dynamodbattribute.UnmarshalMap(result.Attributes, &item)

	if err != nil {
		return Work{}, err
	}

	return item, nil
}

// DeleteWorkByID Delete DynamoDB work By ID
func DeleteWorkByID(svc *dynamodb.DynamoDB, id string) error {

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(WorkTableName),
	}

	_, err := svc.DeleteItem(input)

	if err != nil {
		return err
	}

	return nil

}

// DeleteWorkByUserID Delete DynamoDB work By UserID
func DeleteWorkByUserID(svc *dynamodb.DynamoDB, userID string) error {

	filt := expression.Name("userId").Equal(expression.Value(userID))

	expr, err := expression.NewBuilder().WithFilter(filt).Build()

	if err != nil {
		return err
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String(WorkTableName),
	}

	result, err := svc.Scan(params)

	if err != nil {
		return err
	}

	writeRequest := []*dynamodb.WriteRequest{}

	if len(result.Items) != 0 {
		for _, i := range result.Items {
			item := Work{}
			err = dynamodbattribute.UnmarshalMap(i, &item)

			writeRequest = append(
				writeRequest,
				&dynamodb.WriteRequest{
					DeleteRequest: &dynamodb.DeleteRequest{
						Key: map[string]*dynamodb.AttributeValue{
							"id": {
								S: aws.String(item.ID),
							},
						},
					},
				},
			)

			if err != nil {
				return err
			}
		}

		input := &dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]*dynamodb.WriteRequest{
				WorkTableName: writeRequest,
			},
		}

		_, err = svc.BatchWriteItem(input)

		if err != nil {
			return err
		}

	}

	return nil
}
