package user

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//Createuser type
type createUser struct {
	User User `json:"user"`
}

// CreateUserHandle Create User Handle
func createUserHandle(arg createUser) (interface{}, error) {

	session, err := session.NewSession(
		&aws.Config{Region: aws.String("ap-northeast-1")},
	)

	if err != nil {
		return nil, err
	}

	svc := dynamodb.New(session)

	arg.user.id = ""
	fmt.Println("print ID %+v\n", arg.User)

	usr, err := dynamodbAttribute.MarshalMap(arg.User)
	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
		return nil, err
	}

	imput := &dynamodb.PutItemInput{
		Item:      user,
		TableName: aws.String("portal-users"),
	}

	_, err = svc.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		return nil, err
	}

	return arg.user, nil
}
