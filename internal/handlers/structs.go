package handlers

type EmailRequestBody struct {
	Email string `json:"email" binding:"required,email"`
}
