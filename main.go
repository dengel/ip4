package main

import (
  "log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

  log.Printf("Headers:\n")
  for key, value := range request.Headers {
    log.Printf("    %s: %s\n", key, value)
  }

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       request.Headers.Get("X-Forwarded-For")
	}, nil

}

func main() {
	lambda.Start(handler)
}
