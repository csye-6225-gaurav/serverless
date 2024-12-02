package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/csye-6225-gaurav/serverless/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func HandleUserEmailVerification(ctx context.Context, event events.SNSEvent) error {
	secretName := "lambda_secret_SG"
	region := "us-east-1"
	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Println(err)
	}
	svc := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}
	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		log.Println(err.Error())
	}
	var secretString string = *result.SecretString
	var message models.Message
	log.Println(event.Records[0].SNS.Message)
	err = json.Unmarshal([]byte(event.Records[0].SNS.Message), &message)
	if err != nil {
		log.Println("Failed to parse json")
	}

	from := mail.NewEmail("noreply", "noreply@gauravgunjal.me") // Change to your verified sender

	subject := "Verify Your Email Address"

	to := mail.NewEmail("New User", message.Email) // Change to your recipient
	url := fmt.Sprintf("https://%s?user=%s&token=%s", os.Getenv("URL"), message.Email, message.Token)
	emailBody := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	  <meta charset="UTF-8">
	  <meta name="viewport" content="width=device-width, initial-scale=1.0">
	  <title>Email Verification</title>
	</head>
	<body>
	  <p>Hi %s,</p>
	  <p>Thank you for signing up! To complete the registration process, please verify your email address.</p>
	  <p>
	    <a href="%s" style="color: #1a73e8; text-decoration: none;">Click here to verify your email</a>
	  </p>
	  <p>This link will expire in 2 hours, so be sure to verify your email soon.</p>
	  <p>If you did not create an account, please ignore this email or contact our support team if you have any questions.</p>
	  <p>Best regards,</p>
	</body>
	</html>
	`, to.Name, url)
	plainTextContent := ""
	log.Println("plainText")
	htmlContent := emailBody

	email := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	log.Println("message created")
	client := sendgrid.NewSendClient(secretString)
	log.Println("client created")
	response, err := client.Send(email)
	log.Println("email sent")
	if err != nil {

		log.Println(err)

	} else {

		log.Println(response.StatusCode)

	}

	return nil
}
