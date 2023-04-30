package main

import (
    "context"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/awslabs/aws-lambda-go-api-proxy/gin"
    "mood-api/handlers"
)

var ginLambda *ginadapter.GinLambda

const (
    healthPath = "/api/health"
    moodPathV1 = "/api/v1/mood"
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

    // Configure CORS
    config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://www.worldmoodtoday.com", "https://develop.worldmoodtoday.com","http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
    g.Use(cors.New(config))

    // Create Gin Lambda instance
    ginLambda = ginadapter.New(g)
}

func main() {
    lambda.Start(Handler)
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    return ginLambda.ProxyWithContext(ctx, request)
}
