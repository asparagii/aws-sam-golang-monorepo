package main

import "fmt"
import "github.com/aws/aws-lambda-go/lambda"

type input struct {
	message string `json:"message:"`
}

func handler(in input) error {
	fmt.Printf("%s", in.message)
	return nil
}

func main() {
	lambda.Start(handler)
}
