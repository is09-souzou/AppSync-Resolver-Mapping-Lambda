package handle

import (
	"log"
	"encoding/json"
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

func printID(arg DeleteUser) {
	log.Printf("print ID %+v\n", arg.ID)
}

// HandleRequest Delete User Handle
func HandleRequest(arg DeleteUser) (interface{}, error) {

	list := []User{}

	list = append(list, User{"id1", "email1", "name1"})
	list = append(list, User{"id2", "email2", "name2"})

	printID(arg)

	log.Printf("list %+v\n", list)

	return User{arg.ID, "email2", "name2"}, nil
}
