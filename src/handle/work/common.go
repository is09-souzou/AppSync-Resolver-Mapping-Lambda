package work

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

// WorkCreateBase WorkCreate work struct
type WorkCreateBase struct {
	UserID      string    `json:"userId"`
	Title       string    `json:"title"`
	Tags        *[]string `json:"tags"`
	ImageURI    string    `json:"imageUri"`
	Description string    `json:"description"`
}

// WorkCreate create work struct
type WorkCreate struct {
	Work WorkCreateBase `json:"work"`
}

// WorkUpdateBase WorkUpdate work struct
type WorkUpdateBase struct {
	ID          string    `json:"id"`
	UserID      string    `json:"userId"`
	Title       *string   `json:"title"`
	Tags        *[]string `json:"tags"`
	ImageURI    *string   `json:"imageUri"`
	Description *string   `json:"description"`
}

// WorkUpdate update work struct
type WorkUpdate struct {
	Work WorkUpdateBase `json:"work"`
}

// WorkDelete delete work struct
type WorkDelete struct {
	ID string `json:"id"`
}
