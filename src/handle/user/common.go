package user

// User type
type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Career    string `json:"career"`
	AvatarURI string `json:"avatarUri"`
	Message   string `json:"message"`
}
