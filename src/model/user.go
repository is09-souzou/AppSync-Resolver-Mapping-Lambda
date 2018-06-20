package model

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// UserTableName DynamoDB User Table Name
const UserTableName = "portal-users"

// User DynamoDB User Struct
type User struct {
	ID        string
	Email     string
	Name      string
	Career    string
	AvatarURI string
	Message   string
}

// CreateUser Get user list By ID from DynamoDB
func CreateUser(
	id *string,
	email *string,
	name *string,
	career *string,
	avatarURI *string,
	message *string,
) error {

	svc, err := getSVC()

	if err != nil {
		return err
	}

	if id == nil && email == nil && name == nil && career == nil && avatarURI == nil && message == nil {
		return errors.New("required new value")
	}

	var item = map[string]*dynamodb.AttributeValue{}

	if id != nil {
		item["id"].S = aws.String(*id)
	}

	if email != nil {
		item["email"].S = aws.String(*email)
	}

	if name != nil {
		item["name"].S = aws.String(*name)
	}

	if career != nil {
		item["career"].S = aws.String(*career)
	}

	if avatarURI != nil {
		item["avatarURI"].S = aws.String(*avatarURI)
	}

	if message != nil {
		item["message"].S = aws.String(*message)
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

// GetUserByID Get user by ID from DynamoDB
func GetUserByID(id string) (User, error) {

	svc, err := getSVC()

	if err != nil {
		return User{}, err
	}

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(UserTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return User{}, err
	}

	item := User{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		return User{}, err
	}

	return item, nil
}

// GetUserList Get user list By ID from DynamoDB
func GetUserList() ([]User, error) {

	svc, err := getSVC()

	if err != nil {
		return []User{}, err
	}

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(UserTableName),
	})

	item := []User{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		return []User{}, err
	}

	return item, nil
}

// UpdateUserByID Update user By ID to DynamoDB
func UpdateUserByID(
	id *string,
	email *string,
	name *string,
	career *string,
	avatarURI *string,
	message *string,
) error {

	svc, err := getSVC()

	if err != nil {
		return err
	}

	if id == nil && email == nil && name == nil && career == nil && avatarURI == nil && message == nil {
		return errors.New("required new value")
	}

	var expressionAttributeValues = map[string]*dynamodb.AttributeValue{}

	if email != nil {
		expressionAttributeValues["email"].S = aws.String(*email)
	}

	if name != nil {
		expressionAttributeValues["name"].S = aws.String(*name)
	}

	if career != nil {
		expressionAttributeValues["career"].S = aws.String(*career)
	}

	if avatarURI != nil {
		expressionAttributeValues["avatarURI"].S = aws.String(*avatarURI)
	}

	if message != nil {
		expressionAttributeValues["message"].S = aws.String(*message)
	}

	input := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(UserTableName),
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
