package user

import (
	"fmt"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/define"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DeleteUserHandle Delete User Handle
func DeleteUserHandle(arg UserDelete, identity define.Identity) (interface{}, error) {

	session, err := session.NewSession(
		&aws.Config{Region: aws.String("ap-northeast-1")},
	)
	if err != nil {
		return nil, err
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
