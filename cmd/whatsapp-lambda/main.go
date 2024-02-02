package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/peterehik/goland/pkg/whatsapp"
	"log"
)

func main() {
	lambda.Start(whatsapp.HandleLambdaEvent)
}

func init() {
	log.Printf("Initializing lambda")
}
