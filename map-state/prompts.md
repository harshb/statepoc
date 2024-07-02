Generate terraform code for a Step Function that uses a Map state to iterate over a list of 
items and spawn a task for each item.
These are the steps to generate the code:
1. start Activity state that uses the SQS:SendMessage API to send a message to an SQS queue.
2. prepare Activity state that uses the SQS:SendMessage API to send a message to an SQS queue
3. in-process task uses a Map Task state that generates or fetches the list of item from a lambda;uses the SQS:SendMessage API to send a message to an SQS queue
  
4. done  state that uses the SQS:SendMessage API to send a message to an SQS queue.
5. end state that uses the SQS:SendMessage API to send a message to an SQS queue.
each task sends a "TaskToken.$": "$$.Task.Token", "TaskName":    task.TaskName, to the queue
 

The in-process task uses a Task state that generates or fetches the list of item from a lambda, 
then uses  the Map state to iterate over the list of items and spawn a task for each item.
Example of a map state
```yaml
{
  "Comment": "A dynamic task spawning example using Map state",
  "StartAt": "GenerateItems",
  "States": {
    "GenerateItems": {
      "Type": "Task",
      "Resource": "arn:aws:lambda:REGION:ACCOUNT_ID:function:GenerateItemsFunction",
      "Next": "ProcessItems"
    },
    "ProcessItems": {
      "Type": "Map",
      "ItemsPath": "$.items",
      "Iterator": {
        "StartAt": "ProcessItem",
        "States": {
          "ProcessItem": {
            "Type": "Task",
            "Resource": "arn:aws:lambda:REGION:ACCOUNT_ID:function:ProcessItemFunction",
            "Next": "NextTask"
          },
          "NextTask": {
            "Type": "Task",
            "Resource": "arn:aws:lambda:REGION:ACCOUNT_ID:function:NextFunction",
            "End": true
          }
        }
      },
      "End": true
    }
  }
}

```

Each task is an Activity state that uses the SQS:SendMessage API to send a message to an SQS queue.
Example SQS:SendMessage API:
```yaml
{
  "Comment": "Step Functions workflow to send a message to SQS including a task token",
  "StartAt": "SendMessageToSQS",
  "States": {
    "SendMessageToSQS": {
      "Type": "Task",
      "Resource": "arn:aws:states:::sqs:sendMessage.waitForTaskToken",
      "Parameters": {
        "QueueUrl": "https://sqs.REGION.amazonaws.com/ACCOUNT_ID/QUEUE_NAME",
        "MessageBody": {
          "taskToken.$": "$$.Task.Token",
          "input.$": "$"
        }
      },
      "Next": "WaitForCompletion"
    },
    "WaitForCompletion": {
      "Type": "Wait",
      "Seconds": 300,
      "Next": "TaskCompleted"
    },
    "TaskCompleted": {
      "Type": "Pass",
      "Result": "Task Completed Successfully",
      "End": true
    }
  }
}
```
