package store

import (
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type MoodEntry struct {
	Day          int    `json:"Day"`
	Mood         string `json:"Mood"`
	PartitionKey string `json:"PartitionKey"` // Add this line
}

type MoodDetails struct {
	TotalEntries int                `json:"totalEntries"`
	Percentages  map[string]float64 `json:"percentages"`
	AverageMood  string             `json:"averageMood"`
}

const entriesTableName = "MoodEntries"

func createDynamoDBClient() *dynamodb.DynamoDB {
	region := "us-east-1" // Set the region directly

	awsConfig := &aws.Config{
		Region: aws.String(region),
		//Credentials: credentials.NewStaticCredentials("fakeAccessKeyId", "fakeSecretAccessKey", ""),
	}

	sess := session.Must(session.NewSession(awsConfig))

	dynamoDBClient := dynamodb.New(sess)

	return dynamoDBClient
}

func CalculateMoodDetailsForDate(date string) MoodDetails {
	dynamoDBClient := createDynamoDBClient()
	log.Printf("Input date: %s", date)

	// Get the Unix timestamp for the start and end of the current day
	startOfDay := int(time.Now().UTC().Truncate(24 * time.Hour).Unix())
	endOfDay := startOfDay + 24*60*60 - 1

	input := &dynamodb.QueryInput{
		TableName:              aws.String(entriesTableName),
		KeyConditionExpression: aws.String("#partitionKey = :pk AND #day BETWEEN :start AND :end"),
		ExpressionAttributeNames: map[string]*string{
			"#partitionKey": aws.String("PartitionKey"),
			"#day":          aws.String("Day"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":pk": {
				S: aws.String("mood_entry"),
			},
			":start": {
				N: aws.String(strconv.Itoa(startOfDay)),
			},
			":end": {
				N: aws.String(strconv.Itoa(endOfDay)),
			},
		},
	}
	result, err := dynamoDBClient.Query(input)
	if err != nil {
		// Handle the error
		log.Printf("Error querying DynamoDB table: %v", err)
		return MoodDetails{}
	}

	log.Printf("Query result for date %s: %v", date, result)

	var moodEntries []MoodEntry
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &moodEntries)
	if err != nil {
		// Handle the error
		log.Printf("Error unmarshalling DynamoDB items: %v", err)
		return MoodDetails{}
	}

	// Count the occurrences of each mood type
	moodCounts := map[string]int{}
	for _, entry := range moodEntries {
		moodCounts[entry.Mood]++
	}

	// Calculate the percentage of entries for each mood type
	totalEntries := len(moodEntries)
	percentages := map[string]float64{}
	for mood, count := range moodCounts {
		percentages[mood] = float64(count) / float64(totalEntries) * 100
	}

	// Determine the mood type with the highest percentage
	var highestMood string
	highestPercentage := 0.0
	for mood, percentage := range percentages {
		if percentage > highestPercentage {
			highestMood = mood
			highestPercentage = percentage
		}
	}

	moodDetails := MoodDetails{
		TotalEntries: totalEntries,
		Percentages:  percentages,
		AverageMood:  highestMood,
	}

	return moodDetails
}

func SaveMoodData(mood string) {
	dynamoDBClient := createDynamoDBClient()

	moodEntry := MoodEntry{
		Day:          int(time.Now().Unix()),
		Mood:         mood,
		PartitionKey: "mood_entry", // Add this line
	}
	av, err := dynamodbattribute.MarshalMap(moodEntry)
	if err != nil {
		log.Printf("Error marshalling mood entry to DynamoDB attribute value: %v", err)
		return
	}

	log.Printf("Marshalled item: %v", av) // Add this line to print the marshaled item

	_, err = dynamoDBClient.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(entriesTableName),
		Item:      av,
	})

	if err != nil {
		log.Printf("Error putting mood entry to DynamoDB table: %v", err)
		return
	}

	log.Printf("Successfully saved mood entry to DynamoDB table")
}
