package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// APIKey is the API key used for authentication.
const API_KEY = "YOUR_API_KEY"

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Extract x-signature-timestamp and x-signature-ed25519 values from headers
	// Concatenate x-signature-timestamp and request body

	timestamp := request.Headers["x-signature-timestamp"]
	signature := request.Headers["x-signature-ed25519"]
	bodyTime := timestamp + request.Body

	fmt.Println("Timestamp: ", timestamp)
	fmt.Println("Signature: ", signature)
	fmt.Println("BodyTime: ", bodyTime)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello, World!",
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
