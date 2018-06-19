package work

// Work type
type Work struct {
	ID          string   `json:"id"`
	UserID      string   `json:"userId"`
	Tags        []string `json:"tags"`
	CreatedAt   int      `json:"createdAt"`
	Title       string   `json:"title"`
	ImageURI    string   `json:"imageUri"`
	Description string   `json:"description"`
}