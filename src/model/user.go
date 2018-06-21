package model

import (
	"errors"
	"fmt"
	"strings"

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

	var item = map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String(user.ID),
		},
		"email": {
			S: aws.String(user.Email),
		},
		"displayName": {
			S: aws.String(user.DisplayName),
		},
	}

	if user.Career != nil {
		item["career"] = &dynamodb.AttributeValue{S: aws.String(*user.Career)}
	}

	if user.AvatarURI != nil {
		item["avatarUri"] = &dynamodb.AttributeValue{S: aws.String(*user.AvatarURI)}
	}

	if user.Message != nil {
		item["message"] = &dynamodb.AttributeValue{S: aws.String(*user.Message)}
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(UserTableName),
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
func UpdateUserByID(user UserUpdate) (User, error) {

	svc, err := getSVC()

	if err != nil {
		return User{}, err
	}

	if user.Email == nil && user.DisplayName == nil && user.Career == nil && user.AvatarURI == nil && user.Message == nil {
		return User{}, errors.New("required new value")
	}

	expressionAttributeValues := map[string]*dynamodb.AttributeValue{}
	updateExpression := "SET "

	if user.Email != nil {
		expressionAttributeValues[":email"] = &dynamodb.AttributeValue{S: aws.String(*user.Email)}
		updateExpression += "email = :email, "
	}

	if user.DisplayName != nil {
		expressionAttributeValues[":displayName"] = &dynamodb.AttributeValue{S: aws.String(*user.DisplayName)}
		updateExpression += "displayName = :displayName, "
	}

	if user.Career != nil {
		expressionAttributeValues[":career"] = &dynamodb.AttributeValue{S: aws.String(*user.Career)}
		updateExpression += "career = :career, "
	}

	if user.AvatarURI != nil {
		expressionAttributeValues[":avatarUri"] = &dynamodb.AttributeValue{S: aws.String(*user.AvatarURI)}
		updateExpression += "avatarUri = :avatarUri, "
	}

	if user.Message != nil {
		expressionAttributeValues[":message"] = &dynamodb.AttributeValue{S: aws.String(*user.Message)}
		updateExpression += "message = :message, "
	}

	input := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(UserTableName),
		ExpressionAttributeValues: expressionAttributeValues,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(user.ID),
			},
		},
		ReturnValues:     aws.String("ALL_NEW"),
		UpdateExpression: aws.String(strings.TrimRight(updateExpression, ", ")),
	}

	result, err := svc.UpdateItem(input)

	if err != nil {
		return User{}, err
	}

	fmt.Print(result)

	item := User{}

	err = dynamodbattribute.UnmarshalMap(result.Attributes, &item)

	fmt.Printf("items: %+v", item)

	if err != nil {
		return User{}, err
	}

	return item, nil
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
		TableName: aws.String(UserTableName),
	}

	_, err = svc.DeleteItem(input)

	if err != nil {
		return err
	}

	return nil

}
