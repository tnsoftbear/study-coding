package types

type Task struct {
	Id      	string `json:"id"`
	Name		string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
}
