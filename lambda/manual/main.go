package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest

type Data struct {
	Email  string `json:"email"`
	Status string `json:"status"`
}

func handle(ctx context.Context, req Request) (Response, error) {
	resData := Data{}

	code := 200
	err := json.Unmarshal([]byte(req.Body), &resData)
	if err != nil {
		log.Println(err)
		code = 500
	}

	resData.Status = "Success"

	body, err := json.Marshal(resData)
	if err != nil {
		log.Println(err)
		body = []byte("Internal Server Error")
		code = 500
	}

	var buf bytes.Buffer
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      code,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*", // <-- CORS
		},
	}
	return resp, nil

}

func main() {
	lambda.Start(handle)
}
