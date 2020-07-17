package handlers

type Parameters struct {
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Name string `json:"name"`
}

type Response struct {
	Comment string `json:"comment"`
	Status  string `json:"status"` // succeeded, failed, canceled
}
