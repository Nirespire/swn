package gpt

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/tidwall/gjson"
)

type openAIImageRequest struct {
	Prompt string `json:"prompt"`
	N      int    `json:"n"`
	Size   string `json:"size"`
}

func (r *ImageRequest) GenerateImage() string {

	data := openAIImageRequest{
		r.Prompt,
		1,
		"512x512",
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Failed to marshal OpenAI Image Request", err)
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/images/generations", body)
	if err != nil {
		log.Fatal("Failed to generate OpenAI API request", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", os.ExpandEnv("Bearer $OPENAI_KEY"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Failed to call OpenAI API", err)
	}

	defer resp.Body.Close()

	
	bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
		log.Fatal("Failed read OpenAI API response", err)
    }
    bodyString := string(bodyBytes)
	
	if(resp.StatusCode != 200) {
		log.Fatal("OpenAI API failure: ", bodyString)
	}

	value := gjson.Get(bodyString, "data.0.url")

	return value.String()
}
