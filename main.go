package main

import (
  "log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

  log.Printf("Remote Addr %s\n", request.RequestContext.RemoteAddr)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       request.RequestContext.RemoteAddr,
	}, nil

}

func main() {
	lambda.Start(handler)
}
