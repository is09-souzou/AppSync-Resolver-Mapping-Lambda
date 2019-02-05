package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// PopularTagTableName DynamoDB Popular Tag Table Name
const PopularTagTableName = "portal-popular-tags"

// CreatePopularTag Create popular tag to DynamoDB
func CreatePopularTag(svc *dynamodb.DynamoDB, popularTag PopularTag) error {

	if popularTag.Name == "" {
		return errors.New("required Name in popularTag")
	}
	var item = map[string]*dynamodb.AttributeValue{
		"name": {
			S: aws.String(popularTag.Name),
		},
		"count": {
			N: aws.String(strconv.Itoa(popularTag.Count)),
		},
		"system": {
			S: aws.String("popular-tag"),
		},
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(PopularTagTableName),
	}

	_, err := svc.PutItem(input)

	if err != nil {
		return err
	}

	return nil
}

// GetPopularTagByName Get popular tag by Name from DynamoDB
func GetPopularTagByName(svc *dynamodb.DynamoDB, name string) (PopularTag, error) {

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(PopularTagTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(name),
			},
		},
	})

	if err != nil {
		return PopularTag{}, err
	}

	item := PopularTag{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		return PopularTag{}, err
	}

	return item, nil
}

// ScanPopularTagListResult result ScanPopularTagList
type ScanPopularTagListResult struct {
	Items             []PopularTag
	ExclusiveStartKey *string
}

// ScanPopularTagList Scan popular tag list from DynamoDB
func ScanPopularTagList(svc *dynamodb.DynamoDB, limit int64, exclusiveStartKey *string) (ScanPopularTagListResult, error) {

	params := &dynamodb.QueryInput{
		Limit:     &limit,
		TableName: aws.String(PopularTagTableName),
		KeyConditions: map[string]*dynamodb.Condition{
			"system": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("popular-tag"),
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
			"name": {
				S: aws.String(key.ID),
			},
			"createdAt": {
				S: aws.String(key.CreatedAt),
			},
			"system": {
				S: aws.String("popular-tag"),
			},
		}
	}

	result, err := svc.Query(params)

	if err != nil {
		return ScanPopularTagListResult{}, err
	}

	items := []PopularTag{}

	for _, i := range result.Items {
		item := PopularTag{}

		err := dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			return ScanPopularTagListResult{}, err
		}

		items = append(items, item)
	}

	var respExclusiveStartKey *string
	if result.LastEvaluatedKey != nil {

		exclusiveStartKey := ExclusiveStartKey{
			ID:        *result.LastEvaluatedKey["name"].S,
			CreatedAt: *result.LastEvaluatedKey["createdAt"].S,
		}
		byteExclusiveStartKey, err := json.Marshal(exclusiveStartKey)

		if err != nil {
			fmt.Println("Got error json Marshal exclusiveStartKey")
			fmt.Println(err.Error())
			return ScanPopularTagListResult{}, err
		}

		stringExclusiveStartKey := string(byteExclusiveStartKey)
		respExclusiveStartKey = &stringExclusiveStartKey
	}

	return ScanPopularTagListResult{items, respExclusiveStartKey}, nil
}

// UpdatePopularTagByName Update popular tag By Name to DynamoDB
func UpdatePopularTagByName(svc *dynamodb.DynamoDB, popularTag PopularTag, count string) (PopularTag, error) {

	if popularTag.Name == "" {
		return PopularTag{}, errors.New("required Name in popularTag")
	}

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(PopularTagTableName),
		ExpressionAttributeNames: map[string]*string{
			"#c": aws.String("count"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":c": {
				N: aws.String(count),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(popularTag.Name),
			},
		},
		ReturnValues:     aws.String("ALL_NEW"),
		UpdateExpression: aws.String("ADD #c :c"),
	}

	result, err := svc.UpdateItem(input)

	if err != nil {
		return PopularTag{}, err
	}

	item := PopularTag{}

	err = dynamodbattribute.UnmarshalMap(result.Attributes, &item)

	if err != nil {
		return PopularTag{}, err
	}

	return item, nil
}

// DeletePopularTagByName Delete DynamoDB popular tag By Name
func DeletePopularTagByName(svc *dynamodb.DynamoDB, name string) error {

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(name),
			},
		},
		TableName: aws.String(PopularTagTableName),
	}

	_, err := svc.DeleteItem(input)

	if err != nil {
		return err
	}

	return nil

}
