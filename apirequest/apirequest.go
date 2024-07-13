package apirequest

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GetApiResponse() {
	ctx := context.Background()

	// Access API key as an environment variable
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Access the gemini model using the gemini 1.5 flash
	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text("Write a story about a magic backpack."))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Candidates[0].Content.Parts[0])
}

func main() {

}
