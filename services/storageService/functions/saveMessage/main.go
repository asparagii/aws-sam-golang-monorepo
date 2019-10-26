package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type input struct {
	Message string `json:"message:"`
}

func handler(in input) error {
	fmt.Printf("%s", in.Message)
	return nil
}

func main() {
	lambda.Start(handler)
}
