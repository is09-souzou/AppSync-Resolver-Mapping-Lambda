package handle

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

// Work type
type Work struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// WorkUpdateHandle Update Work Handle
func updateWorkHandle(arg UpdateWork) (interface{}, error) {

	session, err := session.NewSession(
		&aws.Config{Region: aws.String("ap-northeast-1")},
	)
	if err != nil {
		panic(err)
	}

	svc := dynamodb.New(session)

	input := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(arg.ID),
			},
		},
		TableName: aws.String("portal-works"),
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
