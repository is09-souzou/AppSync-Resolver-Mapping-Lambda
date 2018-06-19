package model

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const UserTableName = "portal-users"

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
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

// UpdateUserById Update user list By ID to DynamoDB
func UpdateUserById(
	id *string,
	email *string,
	name *string,
) error {

	svc, err := getSVC()

	if err != nil {
		return err
	}

	if (id == nil && email == nil && name == nil) {
		return errors.New("required new value")
	}


	var key = map[string]*dynamodb.AttributeValue{}

	if (id != nil) {
		key["id"].S = aws.String(*id)
	}

	if (email != nil) {
		key["email"].S = aws.String(*email)
	}

	if (name != nil) {
		key["name"].S = aws.String(*name)
	}

	input := &dynamodb.UpdateItemInput{
		TableName:        aws.String(UserTableName),
		Key:              key,
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String(""),
	}

	_, err = svc.UpdateItem(input)

	if err != nil {
		return err
	}

	return nil
}
