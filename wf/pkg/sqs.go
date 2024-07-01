package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sfn"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type MessageBody struct {
	TaskName      string                 `json:"TaskName"`
	TaskToken     string                 `json:"TaskToken"`
	Input         map[string]interface{} `json:"Input"`
	ReceiptHandle *string
}

func GetAllMessagesFromSQSQueue() ([]MessageBody, error) {
	sess, err := NewAWSSession()
	if err != nil {
		//fmt.Println("Error getting messages:", err)
		return nil, err
	}

	svc := sqs.New(sess)

	queueUrl := SqsQueueUrl

	receiveMessageInput := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueUrl),
		MaxNumberOfMessages: aws.Int64(10), // You can adjust this number based on your needs
		WaitTimeSeconds:     aws.Int64(1),  // Long polling
	}

	var messageBodies []MessageBody

	// Keep receiving messages until there are no more messages
	for {
		output, err := svc.ReceiveMessage(receiveMessageInput)
		if err != nil {
			return nil, err
		}

		if len(output.Messages) == 0 {
			break
		}

		for _, message := range output.Messages {
			var body MessageBody
			err := json.Unmarshal([]byte(*message.Body), &body)
			if err != nil {
				return nil, err
			}
			body.ReceiptHandle = message.ReceiptHandle
			messageBodies = append(messageBodies, body)
		}
	}

	return messageBodies, nil
}

//func SetTokens() error {
//	// Get all messages from the SQS queue
//	messages, err := GetAllMessagesFromSQSQueue()
//	if err != nil {
//		fmt.Println("Error getting messages:", err)
//		return err
//	}
//
//	// Create a new AWS session
//	sess, err := NewAWSSession()
//	if err != nil {
//		return err
//	}
//
//	// Create a new Step Functions service client
//	svc := sfn.New(sess)
//
//	// Create a new SQS service client
//	sqsSvc := sqs.New(sess)
//
//	// Iterate over each message
//	for _, message := range messages {
//		// Marshal the map into a JSON string
//		inputJson, err := json.Marshal(message.Input)
//		if err != nil {
//			fmt.Println("Error marshalling input:", err)
//			return err
//		}
//		// Send a task success signal to the step function
//		sendTaskSuccessInput := &sfn.SendTaskSuccessInput{
//			Output:    aws.String(string(inputJson)),
//			TaskToken: aws.String(message.TaskToken),
//		}
//		_, err = svc.SendTaskSuccess(sendTaskSuccessInput)
//		if err != nil {
//			fmt.Println("Error sending task success:", err)
//			return err
//		}
//		println("Task completed:", message.TaskName)
//
//		// Delete the message from the SQS queue
//		deleteMessageInput := &sqs.DeleteMessageInput{
//			QueueUrl:      aws.String(SqsQueueUrl), // replace with your queue URL
//			ReceiptHandle: message.ReceiptHandle,
//		}
//		_, err = sqsSvc.DeleteMessage(deleteMessageInput)
//		if err != nil {
//			fmt.Println("Error deleting message:", err)
//			return err
//		}
//		println("Message deleted:", *message.ReceiptHandle)
//	}
//
//	return nil
//}

func SetTokens() error {
	// Get all messages from the SQS queue
	messages, err := GetAllMessagesFromSQSQueue()
	if err != nil {
		fmt.Println("Error getting messages:", err)
		return err
	}

	// Create a new AWS session
	sess, err := NewAWSSession()
	if err != nil {
		return err
	}

	// Create a new Step Functions service client
	svc := sfn.New(sess)

	// Iterate over each message
	for _, message := range messages {
		// Marshal the map into a JSON string
		inputJson, err := json.Marshal(message.Input)
		if err != nil {
			fmt.Println("Error marshalling input:", err)
			return err
		}
		// Send a task success signal to the step function
		sendTaskSuccessInput := &sfn.SendTaskSuccessInput{
			Output:    aws.String(string(inputJson)),
			TaskToken: aws.String(message.TaskToken),
		}
		_, err = svc.SendTaskSuccess(sendTaskSuccessInput)
		if err != nil {
			fmt.Println("Error sending task success:", err)
			return err
		}
		println("Task completed:", message.TaskName)
	}

	return nil
}

func PurgeSQSQueue() error {
	sess, err := NewAWSSession()
	if err != nil {
		return err
	}

	svc := sqs.New(sess)

	_, err = svc.PurgeQueue(&sqs.PurgeQueueInput{
		QueueUrl: aws.String(SqsQueueUrl),
	})

	if err != nil {
		return err
	}

	return nil
}
