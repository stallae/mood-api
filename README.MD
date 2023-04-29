# Mood API

[![GitHub](https://img.shields.io/github/license/arthsouza-meli/mood-api)](https://github.com/arthsouza-meli/mood-api/blob/main/LICENSE)

Mood API is a simple RESTful API that provides the back-end services for the [World Mood Today](https://www.worldmoodtoday.com) web application. The API is responsible for collecting mood data from users, calculating the average mood for the day, and serving that information to the front-end.

## Live Demo

Visit [World Mood Today](https://www.worldmoodtoday.com) to see the live demo of the project that uses this API.

## Features

- RESTful API for submitting user moods
- Calculates the average mood for the day
- Serves the average mood and total number of entries to the front-end

## Technologies

- Back-end: Go
- Database: DynamoDB
- Deployment: Amazon S3, Amazon CloudFront, GitHub Actions , Api Gateway, Route 53 and etc.

## Getting Started

To set up the project locally for development, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/arthsouza-meli/mood-api.git
```
2. Install the dependencies and set up the development environment:
Follow the instructions in the project's main README to set up the API server and the MongoDB database.

Start the development server:
```
go run main.go
```
The API server will now run on http://localhost:8000.

## Deployment
This project is automatically deployed to Heroku. The deployment is configured in the Procfile and app.json files in the repository.

## Contributing
We welcome contributions to improve the project. If you have a feature request, bug report, or any suggestions, please open a new issue on the GitHub repository.

## About the Author
Arthur  is an experienced software engineer with a strong background in full-stack development. He is passionate about creating innovative solutions and has demonstrated expertise in multiple programming languages and frameworks. Arthur is a dedicated developer who continuously seeks to improve his skills and contribute to the tech community. To learn more about Arthur and his work, you can visit his LinkedIn profile.