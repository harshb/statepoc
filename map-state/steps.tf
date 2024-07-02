

resource "aws_sfn_state_machine" "sfn_state_machine" {
  name     = local.state_machine_name
  role_arn = aws_iam_role.iam_role.arn

  definition = <<EOF
{
  "Comment": "A state machine that iterates over a list of items and spawns a task for each item.",
  "StartAt": "StartState",
  "States": {
    "StartState": {
      "Type": "Task",
      "Resource": "arn:aws:states:::sqs:sendMessage.waitForTaskToken",
      "Parameters": {
        "QueueUrl": "${aws_sqs_queue.queue.url}",
        "MessageBody": {
          "TaskToken.$": "$$.Task.Token",
          "TaskName": "StartState"
        }
      },
      "Next": "PrepareState"
    },
    "PrepareState": {
      "Type": "Task",
      "Resource": "arn:aws:states:::sqs:sendMessage.waitForTaskToken",
      "Parameters": {
        "QueueUrl": "${aws_sqs_queue.queue.url}",
        "MessageBody": {
          "TaskToken.$": "$$.Task.Token",
          "TaskName": "PrepareState"
        }
      },
      "Next": "InProcessState"
    },
    "InProcessState": {
      "Type": "Task",
      "Resource": "${aws_lambda_function.item_generator.arn}",
      "Next": "ProcessItems"
    },
    "ProcessItems": {
      "Type": "Map",
      "ItemsPath": "$.Payload.items",
      "Iterator": {
        "StartAt": "SendMessage",
        "States": {
          "SendMessage": {
            "Type": "Task",
            "Resource": "arn:aws:states:::sqs:sendMessage.waitForTaskToken",
            "Parameters": {
              "QueueUrl": "${aws_sqs_queue.queue.url}",
              "MessageBody": {
                "TaskToken.$": "$$.Task.Token",
                "TaskName": "InProcessState"
              }
            },
            "End": true
          }
        }
      },
      "Next": "DoneState"
    },
    "DoneState": {
      "Type": "Task",
      "Resource": "arn:aws:states:::sqs:sendMessage.waitForTaskToken",
      "Parameters": {
        "QueueUrl": "${aws_sqs_queue.queue.url}",
        "MessageBody": {
          "TaskToken.$": "$$.Task.Token",
          "TaskName": "DoneState"
        }
      },
      "Next": "EndState"
    },
    "EndState": {
      "Type": "Task",
      "Resource": "arn:aws:states:::sqs:sendMessage.waitForTaskToken",
      "Parameters": {
        "QueueUrl": "${aws_sqs_queue.queue.url}",
        "MessageBody": {
          "TaskToken.$": "$$.Task.Token",
          "TaskName": "EndState"
        }
      },
      "End": true
    }
  }
}
EOF
}
