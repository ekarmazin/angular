package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	"sort"
	//"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/external"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// M as an alias for map
type M map[string]string

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	var dataSlice []M

	svc := cloudwatchevents.New(cfg)
	input := &cloudwatchevents.ListRulesInput{
		NamePrefix: aws.String("robot"),
	}

	req := svc.ListRulesRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		panic("No result came" + err.Error())
	}
	resultData := result.Rules[0].ScheduleExpression

	// Fill out the slice with required data. Must be in type 'string'
	//for i := range result.Contents {
	//	dataSlice = append(dataSlice, M{"Name": *res[i].Key, "Time": res[i].LastModified.String(), "URL": "https://robot.assets.staging.sweet.io/" + *res[i].Key})
	//}

	// Descending sort by date old to new
	sort.Slice(dataSlice, func(i, j int) bool { return dataSlice[i]["Time"] > dataSlice[j]["Time"] })

	body, err := json.Marshal(resultData)
	if err != nil {
		return Response{StatusCode: 404}, err
	}

	//Escape fancy chars to keep JSON clean
	var buf bytes.Buffer
	json.HTMLEscape(&buf, body)

	repl := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*", // <-- CORS
		},
	}
	return repl, nil
}

func main() {
	lambda.Start(Handler)
}
