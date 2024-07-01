package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sfn"
	"io/ioutil"
)

type Task struct {
	TaskName string `json:"taskName"`
}

type Workflow struct {
	WfName string `json:"wfName"`
	Tasks  []Task `json:"tasks"`
}

func CreateActivityStateMachine() error {
	// Read tasks.json
	data, err := ioutil.ReadFile("tasks.json")
	if err != nil {
		return err
	}

	queueUrl := SqsQueueUrl

	// Unmarshal JSON data
	var workflow Workflow
	err = json.Unmarshal(data, &workflow)
	if err != nil {
		return err
	}

	// Create a new AWS session
	sess, err := NewAWSSession()
	svc := sfn.New(sess)

	// Define state machine
	stateMachine := map[string]interface{}{
		"Comment": "A WF for EAC Activities.",
		"StartAt": workflow.Tasks[0].TaskName, // Set the StartAt field to the first task
		"States":  make(map[string]interface{}),
	}

	// Create activities and send messages to SQS
	for i, task := range workflow.Tasks {
		// Create activity
		createActivityResponse, err := svc.CreateActivity(&sfn.CreateActivityInput{
			Name: aws.String(task.TaskName),
		})
		if err != nil {
			return err
		}
		fmt.Printf("Activity ARN: %s\n", *createActivityResponse.ActivityArn)

		state := map[string]interface{}{
			"Type":     "Task",
			"Resource": "arn:aws:states:::sqs:sendMessage.waitForTaskToken",
			"Parameters": map[string]interface{}{
				"QueueUrl": queueUrl,
				"MessageBody": map[string]string{
					"Input.$":     "$",
					"TaskToken.$": "$$.Task.Token",
					"TaskName":    task.TaskName,
				},
			},
		}

		// If it's not the last task, add the Next field
		if i != len(workflow.Tasks)-1 {
			state["Next"] = workflow.Tasks[i+1].TaskName
		} else {
			state["End"] = true
		}

		stateMachine["States"].(map[string]interface{})[task.TaskName] = state
	}

	// Convert the state machine to JSON
	stateMachineDefinition, err := json.Marshal(stateMachine)
	if err != nil {
		return err
	}

	roleArn := StepFunctionExecutionRole

	// Create state machine
	createStateMachineResponse, err := svc.CreateStateMachine(&sfn.CreateStateMachineInput{
		Definition: aws.String(string(stateMachineDefinition)),
		Name:       aws.String(ActivityStateMachineName),
		RoleArn:    aws.String(roleArn),
	})
	if err != nil {
		return err
	}

	// Start state machine
	_, err = svc.StartExecution(&sfn.StartExecutionInput{
		StateMachineArn: createStateMachineResponse.StateMachineArn,
	})
	if err != nil {
		return err
	}

	return nil
}

func DeleteStateMachineAndActivities() error {
	stateMachineArn, err := GetStateMachineArnByName()
	if err != nil {
		fmt.Println(err)
		return err

	}
	err = DeleteAllActivitiesA(stateMachineArn)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = StopExecutionsAndDeleteStateMachine(stateMachineArn)
	if err != nil {
		fmt.Println(err)
		return err

	}
	fmt.Println("Deleted state machine and activities")
	return nil
}

func DeleteAllActivitiesA(stateMachineArn string) error {
	// Create a new AWS session
	sess, err := NewAWSSession()
	if err != nil {
		return err
	}
	svc := sfn.New(sess)

	// List all activities
	listActivitiesOutput, err := svc.ListActivities(&sfn.ListActivitiesInput{})
	if err != nil {
		return err
	}

	// Delete each activity
	for _, activity := range listActivitiesOutput.Activities {
		_, err := svc.DeleteActivity(&sfn.DeleteActivityInput{
			ActivityArn: activity.ActivityArn,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
func StopExecutionsAndDeleteStateMachine(stateMachineArn string) error {
	// Create a new AWS session
	sess, err := NewAWSSession()
	if err != nil {
		return err
	}
	svc := sfn.New(sess)

	// List all executions
	listExecutionsOutput, err := svc.ListExecutions(&sfn.ListExecutionsInput{
		StateMachineArn: aws.String(stateMachineArn),
	})
	if err != nil {
		return err
	}

	// Stop each execution
	for _, execution := range listExecutionsOutput.Executions {
		_, err := svc.StopExecution(&sfn.StopExecutionInput{
			ExecutionArn: execution.ExecutionArn,
		})
		if err != nil {
			return err
		}
	}

	// Delete the state machine
	_, err = svc.DeleteStateMachine(&sfn.DeleteStateMachineInput{
		StateMachineArn: aws.String(stateMachineArn),
	})
	if err != nil {
		return err
	}

	return nil
}
func GetStateMachineArnByName() (string, error) {
	stateMachineName := ActivityStateMachineName
	// Create a new AWS session
	sess, err := NewAWSSession()
	if err != nil {
		return "", err
	}
	svc := sfn.New(sess)

	// List all state machines
	listStateMachinesOutput, err := svc.ListStateMachines(&sfn.ListStateMachinesInput{})
	if err != nil {
		return "", err
	}

	// Find the state machine with the given name
	for _, stateMachine := range listStateMachinesOutput.StateMachines {
		if *stateMachine.Name == stateMachineName {
			return *stateMachine.StateMachineArn, nil
		}
	}

	return "", fmt.Errorf("State machine with name %s not found", stateMachineName)
}
