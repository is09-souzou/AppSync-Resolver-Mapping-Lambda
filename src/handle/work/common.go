package work

// User type
type Work struct {
	ID         string `json:"id"`
	UserID     string `json:"userid"`
	Title      string `json:"title"`
	ImageUri   string `json:"imageuri"`
	Desciption string `json:"description"`
}
