package gpt

type GPT interface {
    GenerateImage()
}

type ImageRequest struct {
	Prompt string
}