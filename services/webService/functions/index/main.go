package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type response struct {
	Message string `json:"message:"`
}

func handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Received event: %v", event)

	response, err := json.Marshal(response{"Everything went right!"})
	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("Unable to marshal response message")
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(response),
	}, nil
}

func main() {
	lambda.Start(handler)
}
