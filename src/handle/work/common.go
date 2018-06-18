package work

// User type
type Work struct {
	ID         string   `json:"id"`
	UserID     string   `json:"userId"`
	Tags       []string `json:"tags`
	CreatedAt  int      `json:"createdAt"`
	Title      string   `json:"title"`
	ImageUri   string   `json:"imageuri"`
	Desciption string   `json:"description"`
}
