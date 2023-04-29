package main

import (
	"context"

	"mood-api/handlers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

const (
	healthPath = "/api/health"
	moodPathV1   = "/api/v1/mood"
)

func init() {
	g := gin.Default()

	// Define health route
	g.GET(healthPath, func(c *gin.Context) {
		handlers.PingHandler(c.Writer, c.Request)
	})

	g.POST(moodPathV1, func(c *gin.Context) {
		handlers.MoodHandler(c.Writer, c.Request)
	})
	g.GET(moodPathV1, func(c *gin.Context) {
		handlers.MoodHandler(c.Writer, c.Request)
	})

	// Create Gin Lambda instance
	ginLambda = ginadapter.New(g)
}

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}
