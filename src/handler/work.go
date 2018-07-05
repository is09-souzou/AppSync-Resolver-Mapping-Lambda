package handler

// Work response work struct
type Work struct {
	ID          string    `json:"id"`
	UserID      string    `json:"userId"`
	Title       string    `json:"title"`
	Tags        *[]string `json:"tags"`
	ImageURI    string    `json:"imageUri"`
	Description string    `json:"description"`
	CreatedAt   int       `json:"createdAt"`
}
