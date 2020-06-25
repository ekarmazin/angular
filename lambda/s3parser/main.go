package main

import (
	"bytes"
	"context"
	"encoding/json"
	"sort"
	"time"
	//"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
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

// M as an alias for map
type M map[string]string

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	location, _ := time.LoadLocation("America/New_York")
	time := time.Now().AddDate(0, 0, -5)
	startAfter := time.Format("200601021504")

	svc := s3.New(cfg)
	inputR := &s3.ListObjectsV2Input{
		Bucket:     aws.String("ss-stage-robot-assets"),
		MaxKeys:    aws.Int64(50),
		StartAfter: aws.String("report-" + startAfter),
	}
	inputL := &s3.ListObjectsV2Input{
		Bucket:     aws.String("ss-stage-robot-assets"),
		MaxKeys:    aws.Int64(50),
		StartAfter: aws.String("log-" + startAfter),
	}

	reqR := svc.ListObjectsV2Request(inputR)
	reqL := svc.ListObjectsV2Request(inputL)
	resultR, err := reqR.Send(context.Background())
	resultL, err := reqL.Send(context.Background())

	resR := resultR.Contents
	resL := resultL.Contents

	var dataSlice []M

	// Fill out the slice with required data. Must be in type 'string'
	for i := range resultR.Contents {
		dataSlice = append(dataSlice, M{"Name": *resR[i].Key, "Time": resR[i].LastModified.In(location).Format("Jan 02 2006 15:04") + " EST", "URL": "https://robot.assets.staging.sweet.io/" + *resR[i].Key})
		dataSlice = append(dataSlice, M{"Name": *resL[i].Key, "Time": resL[i].LastModified.In(location).Format("Jan 02 2006 15:04") + " EST", "URL": "https://robot.assets.staging.sweet.io/" + *resL[i].Key})
	}

	// Descending sort by date old to new
	sort.Slice(dataSlice, func(i, j int) bool { return dataSlice[i]["Time"] > dataSlice[j]["Time"] })

	body, err := json.Marshal(dataSlice)
	if err != nil {
		return Response{StatusCode: 404}, err
	}

	//Escape fancy chars to keep JSON clean
	var buf bytes.Buffer
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
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
	lambda.Start(Handler)
}
