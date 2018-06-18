package handle

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

// User type
type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// HandleRequest Delete User Handle
func deleteUserHandle(arg DeleteUser) (interface{}, error) {

	// list := []User{}

	// list = append(list, User{"id1", "email1", "name1"})
	// list = append(list, User{"id2", "email2", "name2"})
	// log.Printf("list %+v\n", list)

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

	fmt.Println("Deleted")

	return arg, nil
}
