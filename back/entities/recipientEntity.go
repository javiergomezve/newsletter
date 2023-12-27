package entities

type Recipient struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Status   string `json:"status"`
}
