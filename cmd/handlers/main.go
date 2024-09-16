package main

import (
	"context"
	"crypto/ed25519"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bsdlp/discord-interactions-go/interactions"
)

// APIKey is the API key used for authentication.
const API_KEY = "YOUR_API_KEY"

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Extract x-signature-timestamp and x-signature-ed25519 values from headers
	// Concatenate x-signature-timestamp and request body

	verified := interactions.Verify(r, ed25519.PublicKey(discordPubkey))

	fmt.Println("Marshalling json: ")

	jason_marsh, _ := json.Marshal(request)
	fmt.Println(string(jason_marsh))

	timestamp := request.Headers["x-signature-timestamp"]
	signature := request.Headers["x-signature-ed25519"]
	bodyTime := timestamp + request.Body

	fmt.Println("Timestamp: ", timestamp)
	fmt.Println("Signature: ", signature)
	fmt.Println("BodyTime: ", bodyTime)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "He's GHET!",
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
