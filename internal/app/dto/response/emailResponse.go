package response

type EmailResponse struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	EmailMessage string `json:"emailMessage"`
}
