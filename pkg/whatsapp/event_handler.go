package whatsapp

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

type LambdaEventHandler struct {
}

func HandleLambdaEvent(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	body := request.Body
	fmt.Printf("Body: %s\n", body)
	return &events.APIGatewayProxyResponse{Body: body, StatusCode: 200}, nil
}
