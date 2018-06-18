package work

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// UpdateWork type
type UpdateWork struct {
	ID string `json:"id"`
}

// UpdateWorkHandle Update Work Handle
func UpdateWorkHandle(arg UpdateWork) (interface{}, error) {

	session, err := session.NewSession(
		&aws.Config{Region: aws.String("ap-northeast-1")},
	)
	if err != nil {
		panic(err)
	}

	svc := dynamodb.New(session)

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(arg.ID),
			},
		},
		TableName: aws.String("portal-works"),
		Key: map[string]*dynamodb.AttributeValue{
			"userId": {
				S: aws.String(arg.ID),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set info.rating = :r"),
	}

	_, err = svc.UpdateItem(input)

	if err != nil {
		fmt.Println("Got error calling UpdateItem")
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println("update 'The Big New Movie' (2015)")

	return true, nil
}
