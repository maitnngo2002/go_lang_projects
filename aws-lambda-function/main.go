/*
This program is a  serveless  tech where you create functions and host them on the cloud
You only pay for usage

With AWS Lambda, you can run code without provisioning or managing servers

Downside of aws lambda: cold start - when someone's using aws lambda, the lambda takes a while to start

*/

package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json: "What is your name?"`
	Age int `json: "How old are you?"`
}

type MyResponse struct {
	Message string `json: "Answer:"`
}

func HandleLambdaEvent(event MyEvent)(MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old", event.Name, event.Age)}, nil
}
func main() {
	lambda.Start(HandleLambdaEvent)
}