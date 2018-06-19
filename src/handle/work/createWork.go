package work

import (
	"time"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/is09-souzou/AppSync-Resolver-Mapping-Lambda/src/handle/user"
)

// CreateWork type
type CreateWork struct {
	Work Work  `json:"work"`
}

// CreateWorkHandle Create Work Handle
func CreateWorkHandle(arg CreateWork) (interface{}, error) {

	session, err := session.NewSession(
		&aws.Config{Region: aws.String("ap-northeast-1")},
	)

	if err != nil {
		return nil, err
	}

	svc := dynamodb.New(session)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("portal-users"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(arg.Work.UserID),
			},
		},
	})

	user := user.User{}
	
	err = dynamodbattribute.UnmarshalMap(result.Item, &user)

	if err != nil {
		return nil, err
	}

	id, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	arg.Work.ID = id.String()
	arg.Work.CreatedAt = int(time.Now().Unix())
	fmt.Printf("print ID %+v\n", arg.Work)

	work, err := dynamodbattribute.MarshalMap(arg.Work)

	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      work,
		TableName: aws.String("portal-works"),
	}

	_, err = svc.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		return nil, err
	}

	return arg.Work, nil
}
