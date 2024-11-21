package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/csye-6225-gaurav/serverless/handlers"
)

func main() {
	log.Println("main start")
	lambda.Start(handlers.HandleUserEmailVerification)
}
