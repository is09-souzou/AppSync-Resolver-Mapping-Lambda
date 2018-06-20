package model

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// UserTableName DynamoDB User Table Name
const UserTableName = "portal-users"

// CreateUser Create user to DynamoDB
func CreateUser(user UserCreate) error {

	svc, err := getSVC()

	if err != nil {
		return err
	}

	var item = map[string]*dynamodb.AttributeValue{}

	item["id"].S = aws.String(user.ID)
	item["email"].S = aws.String(user.Email)
	item["name"].S = aws.String(user.Name)

	if user.Career != nil {
		item["career"].S = aws.String(*user.Career)
	}

	if user.AvatarURI != nil {
		item["avatarURI"].S = aws.String(*user.AvatarURI)
	}

	if user.Message != nil {
		item["message"].S = aws.String(*user.Message)
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
func UpdateUserByID(user UserUpdate) error {

	svc, err := getSVC()

	if err != nil {
		return err
	}

	if user.Email == nil && user.Name == nil && user.Career == nil && user.AvatarURI == nil && user.Message == nil {
		return errors.New("required new value")
	}

	var expressionAttributeValues = map[string]*dynamodb.AttributeValue{}

	if user.Email != nil {
		expressionAttributeValues["email"].S = aws.String(*user.Email)
	}

	if user.Name != nil {
		expressionAttributeValues["name"].S = aws.String(*user.Name)
	}

	if user.Career != nil {
		expressionAttributeValues["career"].S = aws.String(*user.Career)
	}

	if user.AvatarURI != nil {
		expressionAttributeValues["avatarURI"].S = aws.String(*user.AvatarURI)
	}

	if user.Message != nil {
		expressionAttributeValues["message"].S = aws.String(*user.Message)
	}

	input := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(UserTableName),
		ExpressionAttributeValues: expressionAttributeValues,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(user.ID),
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

// DeleteUserByID Delete DynamoDB user By ID
func DeleteUserByID(id string) error {

	svc, err := getSVC()

	if err != nil {
		return err
	}

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String("portal-users"),
	}

	_, err = svc.DeleteItem(input)

	if err != nil {
		return err
	}

	return nil

}
