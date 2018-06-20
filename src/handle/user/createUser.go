package user

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/define"
)

//CreateUser type
type CreateUser struct {
	User User `json:"User"`
}

// UserTableName DynamoDB User Table Name
const UserTableName = "portal-users"

// CreateUserHandle Create User Handle
func CreateUserHandle(arg CreateUser, identity define.Identity) (interface{}, error) {

	session, err := session.NewSession(
		&aws.Config{Region: aws.String("ap-northeast-1")},
	)

	if err != nil {
		return nil, err
	}

	svc := dynamodb.New(session)

	arg.User.ID = identity.Sub
	fmt.Printf("print ID %+v\n", arg.User)

	user, err := dynamodbattribute.MarshalMap(arg.User)
	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      user,
		TableName: aws.String(UserTableName),
	}

	_, err = svc.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		return nil, err
	}

	return arg.User, nil
}
