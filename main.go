package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func hello() (string, error) {
	return "Hello λ!", nil
}

func kinesisHandler(ctx context.Context, event events.KinesisEvent) (string, error) {
	for _, record := range event.Records {
		log.Println("---------")
		log.Printf("RecordID: %s\n", record.EventID)
		log.Printf("RecordID: %s\n", string(record.Kinesis.Data))
	}
	log.Println("called.", len(event.Records))
	return "Hello kinesis!", nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(kinesisHandler)
}
