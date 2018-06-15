package handle

import (
	"fmt"
	"encoding/json"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

type request struct {
	Field     string `json:"field"`
	Arguments json.RawMessage `json:"arguments"`
}

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
func HandleRequest(arg DeleteUser) (interface{}, error) {

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
				N: aws.String(arg.ID),
			},
		},
		TableName: aws.String("WorksTable"),
	}
	
	_, err = svc.DeleteItem(input)
	
	if err != nil {
		fmt.Println("Got error calling DeleteItem")
		fmt.Println(err.Error())
		return nil, err
	}
	
	fmt.Println("Deleted 'The Big New Movie' (2015)")
	
	return true, nil
}
