package model

import (
	"encoding/json"
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
func CreateUser(svc *dynamodb.DynamoDB, user UserCreate) error {

	if user.ID == "" {
		return errors.New("required ID in user")
	}

	var item = map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String(user.ID),
		},
		"displayName": {
			S: aws.String(user.DisplayName),
		},
		"system": {
			S: aws.String("user"),
		},
	}

	if user.Email != nil {
		if *user.Email == "" {
			*user.Email = " "
		}
		item["email"] = &dynamodb.AttributeValue{S: aws.String(*user.Email)}
	}

	if user.Career != nil {
		if *user.Career == "" {
			*user.Career = " "
		}
		item["career"] = &dynamodb.AttributeValue{S: aws.String(*user.Career)}
	}

	if user.AvatarURI != nil {
		if *user.AvatarURI == "" {
			*user.AvatarURI = " "
		}
		item["avatarUri"] = &dynamodb.AttributeValue{S: aws.String(*user.AvatarURI)}
	}

	if user.Message != nil {
		if *user.Message == "" {
			*user.Message = " "
		}
		item["message"] = &dynamodb.AttributeValue{S: aws.String(*user.Message)}
	}

	if user.SkillList != nil && len(user.SkillList) != 0 {
		item["skillList"] = &dynamodb.AttributeValue{SS: aws.StringSlice(user.SkillList)}
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(UserTableName),
	}

	_, err := svc.PutItem(input)

	if err != nil {
		return err
	}

	return nil
}

// GetUserByID Get user by ID from DynamoDB
func GetUserByID(svc *dynamodb.DynamoDB, id string) (User, error) {

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

// ScanUserListResult result ScanUserList
type ScanUserListResult struct {
	Items             []User
	ExclusiveStartKey *string
}

// ScanUserList Scan user list from DynamoDB
func ScanUserList(svc *dynamodb.DynamoDB, limit int64, exclusiveStartKey *string) (ScanUserListResult, error) {

	params := &dynamodb.QueryInput{
		Limit:     &limit,
		TableName: aws.String(UserTableName),
		KeyConditions: map[string]*dynamodb.Condition{
			"system": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("user"),
					},
				},
			},
		},
		IndexName:        aws.String("system-createdAt-index"),
		ScanIndexForward: aws.Bool(false),
	}

	if exclusiveStartKey != nil {

		jsonBytes := ([]byte)(*exclusiveStartKey)

		var key ExclusiveStartKey
		json.Unmarshal(jsonBytes, &key)

		params.ExclusiveStartKey = map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(key.ID),
			},
			"createdAt": {
				S: aws.String(key.CreatedAt),
			},
			"system": {
				S: aws.String("user"),
			},
		}
	}

	result, err := svc.Query(params)

	if err != nil {
		return ScanUserListResult{}, err
	}

	items := []User{}

	for _, i := range result.Items {
		item := User{}

		err := dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			return ScanUserListResult{}, err
		}

		items = append(items, item)
	}

	var respExclusiveStartKey *string
	if result.LastEvaluatedKey != nil {

		exclusiveStartKey := ExclusiveStartKey{
			ID:        *result.LastEvaluatedKey["id"].S,
			CreatedAt: *result.LastEvaluatedKey["createdAt"].S,
		}
		byteExclusiveStartKey, err := json.Marshal(exclusiveStartKey)

		if err != nil {
			fmt.Println("Got error json Marshal exclusiveStartKey")
			fmt.Println(err.Error())
			return ScanUserListResult{}, err
		}

		stringExclusiveStartKey := string(byteExclusiveStartKey)
		respExclusiveStartKey = &stringExclusiveStartKey
	}

	return ScanUserListResult{items, respExclusiveStartKey}, nil
}

// UpdateUserByID Update user By ID to DynamoDB
func UpdateUserByID(svc *dynamodb.DynamoDB, user UserUpdate) (User, error) {

	if user.Email == nil && user.DisplayName == nil && user.Career == nil && user.AvatarURI == nil && user.Message == nil {
		return User{}, errors.New("required new value")
	}

	expressionAttributeValues := map[string]*dynamodb.AttributeValue{}
	updateExpression := "SET "

	if user.Email != nil {
		if *user.Email == "" {
			*user.Email = " "
		}
		expressionAttributeValues[":email"] = &dynamodb.AttributeValue{S: aws.String(*user.Email)}
		updateExpression += "email = :email, "
	}

	if user.DisplayName != nil {
		if *user.DisplayName == "" {
			*user.DisplayName = " "
		}
		expressionAttributeValues[":displayName"] = &dynamodb.AttributeValue{S: aws.String(*user.DisplayName)}
		updateExpression += "displayName = :displayName, "
	}

	if user.Career != nil {
		if *user.Career == "" {
			*user.Career = " "
		}
		expressionAttributeValues[":career"] = &dynamodb.AttributeValue{S: aws.String(*user.Career)}
		updateExpression += "career = :career, "
	}

	if user.AvatarURI != nil {
		if *user.AvatarURI == "" {
			*user.AvatarURI = " "
		}
		expressionAttributeValues[":avatarUri"] = &dynamodb.AttributeValue{S: aws.String(*user.AvatarURI)}
		updateExpression += "avatarUri = :avatarUri, "
	}

	if user.Message != nil {
		if *user.Message == "" {
			*user.Message = " "
		}
		expressionAttributeValues[":message"] = &dynamodb.AttributeValue{S: aws.String(*user.Message)}
		updateExpression += "message = :message, "
	}

	if user.SkillList != nil && len(*user.SkillList) != 0 {
		expressionAttributeValues[":skillList"] = &dynamodb.AttributeValue{SS: aws.StringSlice(*user.SkillList)}
		updateExpression += "skillList = :skillList, "
	}

	if user.FavoriteWorkList != nil && len(*user.FavoriteWorkList) != 0 {
		expressionAttributeValues[":FavoriteWorkList"] = &dynamodb.AttributeValue{SS: aws.StringSlice(*user.FavoriteWorkList)}
		updateExpression += "FavoriteWorkList = :FavoriteWorkList, "
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

	item := User{}

	err = dynamodbattribute.UnmarshalMap(result.Attributes, &item)

	if err != nil {
		return User{}, err
	}

	return item, nil
}

// DeleteUserByID Delete DynamoDB user By ID
func DeleteUserByID(svc *dynamodb.DynamoDB, id string) error {

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(UserTableName),
	}

	_, err := svc.DeleteItem(input)

	if err != nil {
		return err
	}

	return nil

}
