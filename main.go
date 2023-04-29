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
	moodPath   = "/api/mood"
)

func init() {
	g := gin.Default()

	// Define health route
	g.GET(healthPath, func(c *gin.Context) {
		handlers.PingHandler(c.Writer, c.Request)
	})

	// Define mood routes
	g.POST(moodPath, func(c *gin.Context) {
		handlers.MoodHandler(c.Writer, c.Request)
	})
	g.GET(moodPath, func(c *gin.Context) {
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
