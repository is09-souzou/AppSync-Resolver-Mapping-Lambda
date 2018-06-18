package user

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DeleteUser type
type DeleteUser struct {
	ID string `json:"id"`
}

// DeleteUserHandle Delete User Handle
func DeleteUserHandle(arg DeleteUser) (interface{}, error) {

	session, err := session.NewSession(
		&aws.Config{Region: aws.String("ap-northeast-1")},
	)
	if err != nil {
		panic(err)
	}

	svc := dynamodb.New(session)

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(arg.ID),
			},
		},
		TableName: aws.String("portal-users"),
	}

	_, err = svc.DeleteItem(input)

	if err != nil {
		fmt.Println("Got error calling DeleteItem")
		fmt.Println(err.Error())
		return nil, err
	}

	return arg, nil
}
