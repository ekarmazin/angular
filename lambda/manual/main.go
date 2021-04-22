package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"log"
	"os"
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
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}
	svc := ecs.New(cfg)
	resData := Data{}
	code := 200

	log.Print("Request FE: ", req.Body) //{"content":"{ \"email\": \"eugene@sweet.io\" }

	// Parse the request data
	err = json.Unmarshal([]byte(req.Body), &resData)
	if err != nil {
		log.Println(err)
		code = 500
	}
	email := resData.Email

	log.Print(resData)

	// ECS Task parameters
	params := &ecs.RunTaskInput{
		TaskDefinition: aws.String(os.Getenv("TASK_DEFINITION")), // Required
		Cluster:        aws.String(os.Getenv("CLUSTER_NAME")),
		Count:          aws.Int64(1),
		NetworkConfiguration: &ecs.NetworkConfiguration{
			AwsvpcConfiguration: &ecs.AwsVpcConfiguration{
				AssignPublicIp: ecs.AssignPublicIpEnabled,
				SecurityGroups: []string{os.Getenv("SECURITY_GROUP")},
				Subnets:        []string{os.Getenv("SUBNET_IDS")},
			},
		},
		Overrides: &ecs.TaskOverride{
			ContainerOverrides: []ecs.ContainerOverride{
				{Environment: []ecs.KeyValuePair{
					{Name: aws.String("EMAIL"),
						Value: aws.String(email)},
				},
					Name: aws.String("robot-framework"),
				},
			},
		},
		StartedBy:       aws.String("On-Demand"),
		PlatformVersion: aws.String("1.4.0"),
	}

	// Run a task and push errors if any to logs
	log.Print("ECS Task definition: ", os.Getenv("TASK_DEFINITION"))
	ecsReq := svc.RunTaskRequest(params)
	res, err := ecsReq.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecs.ErrCodeServerException:
				fmt.Println(ecs.ErrCodeServerException, aerr.Error())
			case ecs.ErrCodeException:
				fmt.Println(ecs.ErrCodeException, aerr.Error())
			case ecs.ErrCodeInvalidParameterException:
				fmt.Println(ecs.ErrCodeInvalidParameterException, aerr.Error())
			case ecs.ErrCodeClusterNotFoundException:
				fmt.Println(ecs.ErrCodeClusterNotFoundException, aerr.Error())
			case ecs.ErrCodeUnsupportedFeatureException:
				fmt.Println(ecs.ErrCodeUnsupportedFeatureException, aerr.Error())
			case ecs.ErrCodePlatformUnknownException:
				fmt.Println(ecs.ErrCodePlatformUnknownException, aerr.Error())
			case ecs.ErrCodePlatformTaskDefinitionIncompatibilityException:
				fmt.Println(ecs.ErrCodePlatformTaskDefinitionIncompatibilityException, aerr.Error())
			case ecs.ErrCodeAccessDeniedException:
				fmt.Println(ecs.ErrCodeAccessDeniedException, aerr.Error())
			case ecs.ErrCodeBlockedException:
				fmt.Println(ecs.ErrCodeBlockedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	}

	// Prepare response body in json structure
	body, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		body = []byte("Internal Server Error")
		code = 500
	}
	//Escape fancy chars to keep JSON clean
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
