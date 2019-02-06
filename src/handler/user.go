package handler

// User response user struct
type User struct {
	ID               string    `json:"id"`
	Email            *string   `json:"email"`
	DisplayName      string    `json:"displayName"`
	Career           *string   `json:"career"`
	AvatarURI        *string   `json:"avatarUri"`
	Message          *string   `json:"message"`
	SkillList        []string  `json:"skillList"`
	CreatedAt        int       `json:"createdAt"`
	FavoriteWorkList *[]string `json:"favoriteWorkList"`
}

// UserConnection response user connection struct
type UserConnection struct {
	Items             []User  `json:"items"`
	ExclusiveStartKey *string `json:"exclusiveStartKey"`
}
