package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request map[string]interface{}) (string, error) {
	// simply return string representation of the request
	return fmt.Sprintf("Response: %v", request), nil
}

func main() {
	lambda.Start(handler)
}
