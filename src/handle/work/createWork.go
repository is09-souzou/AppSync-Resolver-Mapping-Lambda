package work

import (
	"time"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

// CreateWork type
type CreateWork struct {
	ID          string   `json:"id"`
	UserID      string   `json:"userId"`
	Tags        []string `json:"tags"`
	CreatedAt   int      `json:"createdAt"`
	Title       string   `json:"title"`
	ImageURI    string   `json:"imageUri"`
	Description string   `json:"description"`
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

	id, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	arg.ID = id.String()
	arg.CreatedAt = int(time.Now().Unix())
	fmt.Printf("print ID %+v\n", arg.ID)

	work, err := dynamodbattribute.MarshalMap(arg)

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

	return arg, nil
}
