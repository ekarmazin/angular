package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

type myReturn struct {
	Response string `json:"response"`
}

func handle(ctx context.Context, name events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Print("Request body: ", name)
	log.Print("context ", ctx)
	headers := map[string]string{"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Headers": "Origin, X-Requested-With, Content-Type, Accept"}

	code := 200
	response, error := json.Marshal(myReturn{Response: "Hello, " + name.Body})
	if error != nil {
		log.Println(error)
		response = []byte("Internal Server Error")
		code = 500
	}

	return events.APIGatewayProxyResponse{code, headers, string(response), false}, nil
}

func main() {
	lambda.Start(handle)
}
