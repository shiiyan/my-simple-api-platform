package handlers

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetSummaryResponse struct {
	User  User  `json:"user"`
	Order []any `json:"order"`
}
