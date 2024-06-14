package model

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age" validate:"required,min=1"`
	Email string `json:"email" validate:"required,email"`
}

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}
