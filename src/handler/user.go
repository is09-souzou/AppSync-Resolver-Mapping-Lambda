package handler

// User response user struct
type User struct {
	ID          string  `json:"id"`
	Email       *string `json:"email"`
	DisplayName string  `json:"displayName"`
	Career      *string `json:"career"`
	AvatarURI   *string `json:"avatarUri"`
	Message     *string `json:"message"`
}
