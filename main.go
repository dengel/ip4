package main

import (
  "log"
  "strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

  source := strings.Split(request.Headers["X-Forwarded-For"], ",")[0]

  log.Printf("Connection from: %s\n", source)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       source,
	}, nil

}

func main() {
	lambda.Start(handler)
}
