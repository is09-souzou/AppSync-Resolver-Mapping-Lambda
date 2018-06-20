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

// Work DynamoDB Work Struct
type Work struct {
	ID          string
	UserID      string
	Title       string
	Tag         []string
	ImageURI    string
	Description string
	CreatedAt   int
}

// CreateUser Get user list By ID from DynamoDB
func CreateWork(
	id *string,
	userID *string,
	title *string,
	tag *[]string,
	imageURI *string,
	description *string,
	createdAt *int,
) error {

	svc, err := getSVC()

	if err != nil {
		return err
	}

	if id == nil && userID == nil && title == nil && tag == nil && imageURI == nil && description == nil && createdAt == nil {
		return errors.New("required new value")
	}

	var item = map[string]*dynamodb.AttributeValue{}

	if id != nil {
		item["userId"].S = aws.String(*userID)
	}

	if userID != nil {
		item["userId"].S = aws.String(*userID)
	}

	if title != nil {
		item["title"].S = aws.String(*title)
	}

	if tag != nil {
		item["tag"].SS = aws.StringSlice(*tag)
	}

	if imageURI != nil {
		item["imageUri"].S = aws.String(*imageURI)
	}

	if description != nil {
		item["description"].S = aws.String(*description)
	}

	if createdAt != nil {
		item["createdAt"].N = aws.String(strconv.Itoa(*createdAt))
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("Movies"),
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
func UpdateWorkByID(
	id *string,
	userID *string,
	title *string,
	tag *[]string,
	imageURI *string,
	description *string,
	createdAt *int,
) error {

	svc, err := getSVC()

	if err != nil {
		return err
	}

	if id == nil && userID == nil && title == nil && tag == nil && imageURI == nil && description == nil && createdAt == nil {
		return errors.New("required new value")
	}

	var expressionAttributeValues = map[string]*dynamodb.AttributeValue{}

	if userID != nil {
		expressionAttributeValues["userId"].S = aws.String(*userID)
	}

	if title != nil {
		expressionAttributeValues["title"].S = aws.String(*title)
	}

	if tag != nil {
		expressionAttributeValues["tag"].SS = aws.StringSlice(*tag)
	}

	if imageURI != nil {
		expressionAttributeValues["imageUri"].S = aws.String(*imageURI)
	}

	if description != nil {
		expressionAttributeValues["description"].S = aws.String(*description)
	}

	if createdAt != nil {
		expressionAttributeValues["createdAt"].N = aws.String(strconv.Itoa(*createdAt))
	}

	input := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(WorkTableName),
		ExpressionAttributeValues: expressionAttributeValues,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(*id),
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
