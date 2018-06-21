package user

// User response user struct
type User struct {
	ID          string  `json:"id"`
	Email       string  `json:"email"`
	DisplayName string  `json:"displayName"`
	Career      *string `json:"career"`
	AvatarURI   *string `json:"avatarUri"`
	Message     *string `json:"message"`
}

// UserCreateBase UserCreate user struct
type UserCreateBase struct {
	Email       string  `json:"email"`
	DisplayName string  `json:"displayName"`
	Career      *string `json:"career"`
	AvatarURI   *string `json:"avatarUri"`
	Message     *string `json:"message"`
}

// UserCreate create user struct
type UserCreate struct {
	User UserCreateBase `json:"user"`
}

// UserUpdateBase UserUpdate user struct
type UserUpdateBase struct {
	ID          string  `json:"id"`
	Email       *string `json:"email"`
	DisplayName *string `json:"displayName"`
	Career      *string `json:"career"`
	AvatarURI   *string `json:"avatarUri"`
	Message     *string `json:"message"`
}

// UserUpdate create user struct
type UserUpdate struct {
	User UserUpdateBase `json:"user"`
}

// UserDelete delete user struct
type UserDelete struct {
	ID string `json:"id"`
}
