package main

import (
	"bytes"
	"context"
	"encoding/json"
	//"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	//"log"

	//"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/external"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// M as an alias for map interface
type M map[string]interface{}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	var buf bytes.Buffer
	var dataSlice []M

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := s3.New(cfg)
	input := &s3.ListObjectsInput{
		Bucket:  aws.String("qs-production-angular-dev"),
		MaxKeys: aws.Int64(10),
	}

	req := svc.ListObjectsRequest(input)
	result, err := req.Send(context.Background())
	res := result.Contents

	for i := range result.Contents {
		dataSlice = append(dataSlice, M{"Name": res[i].Key, "Time": res[i].LastModified, "URL": "https://assets.karmazin.me/" + *res[i].Key})
	}

	body, err := json.Marshal(dataSlice)
	if err != nil {
		return Response{StatusCode: 404}, err
	}

	//Escape fancy chars to keep JSON clean
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
