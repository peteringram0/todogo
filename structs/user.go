package structs

type User struct {
	Name    string         `json:"name"`
	Picture string         `json:"picture"`
	Email   string         `json:"email"`
	Tasks   TaskCollection `json:"tasks"`
}
