package main

import "fmt"
import "github.com/aws/aws-lambda-go/lambda"

type input struct {
	Message string `json:"message:"`
	Age     int    `json:"age:"`
}

type output struct {
	Test string `json:"response:"`
}

func handler(in input) (output, error) {
	fmt.Printf("Received message: %+v", in)
	return output{
		Test: in.Message,
	}, nil
}

func main() {
	lambda.Start(handler)
}
