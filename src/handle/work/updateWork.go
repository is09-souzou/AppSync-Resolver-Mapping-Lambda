package work

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// UpdateWork type
type UpdateWork struct {
	Work Work `json:"word"`
}

// UpdateWorkHandle Update Work Handle
func UpdateWorkHandle(arg UpdateWork) (interface{}, error) {

	session, err := session.NewSession(
		&aws.Config{Region: aws.String("ap-northeast-1")},
	)
	if err != nil {
		return nil, err
	}

	svc := dynamodb.New(session)

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(""),
			},
			"userId": {
				S: aws.String("hashiserver@gmail.com"),
			},
		},
		TableName: aws.String("portal-works"),
		Key: map[string]*dynamodb.AttributeValue{
			"title": {
				S: aws.String("test"),
			},
			"tags": {
				SS: aws.StringSlice(arg.Work.Tags),
			},
			"imageUri": {
				S: aws.String("https://create"),
			},
			"description": {
				S: aws.String("testCode"),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String(""),
	}

	_, err = svc.UpdateItem(input)

	if err != nil {
		fmt.Println("Got error calling UpdateItem:")
		fmt.Println(err.Error())
		return nil, err
	}

	return true, nil
}
