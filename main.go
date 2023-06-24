package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type RequestBody struct {
	Name string `json:"name"`
}
type ResponseBody struct {
	Message string `json:"message"`
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %d.\n", len(request.Body))

	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %s: %s\n", key, value)
	}

	var requestBody RequestBody
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	msg := fmt.Sprintf("Hello %v", requestBody.Name)
	responseBody := ResponseBody{
		Message: msg,
	}
	jbytes, err := json.Marshal(responseBody)
	if err != nil {
		responseBody.Message = fmt.Sprintf("Internal Server Error")
	}
	jstr := string(jbytes)

	return events.APIGatewayProxyResponse{Body: jstr, StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
