resource "aws_sfn_state_machine" "state_machine" {
  name     = local.state_machine_name
  role_arn = aws_iam_role.sfn_role.arn

  definition = <<EOF
{
  "Comment": "A state machine that invokes a lambda function to generate items and then processes each item with a map state, saving data to DynamoDB, followed by a completion task.",
  "StartAt": "GetTasks",
  "States": {
    "GetTasks": {
      "Type": "Task",
      "Resource": "${aws_lambda_function.item_generator.arn}",
      "Parameters": {
        "template.$": "$.template"
      },
      "Next": "InProcess"
    },
    "InProcess": {
      "Type": "Map",
      "ItemsPath": "$.body.items",
      "MaxConcurrency": 0,
      "Iterator": {
        "StartAt": "WaitForCompletion",
        "States": {
          "WaitForCompletion": {
            "Type": "Task",
            "Resource": "arn:aws:states:::lambda:invoke.waitForTaskToken",
            "Parameters": {
              "FunctionName": "${aws_lambda_function.send_token_to_dynamodb.arn}",
              "Payload": {
                "id.$": "$.id",
                "name.$": "$.name",
                "taskToken.$": "$$.Task.Token"
              }
            },
            "End": true
          }
        }
      },
      "Next": "Done"
    },
    "Done": {
      "Type": "Task",
      "Resource": "arn:aws:states:::lambda:invoke.waitForTaskToken",
      "Parameters": {
        "FunctionName": "${aws_lambda_function.send_token_to_dynamodb.arn}",
        "Payload": {
          "name": "done",
          "taskToken.$": "$$.Task.Token"
        }
      },
      "End": true
    }
  }
}
EOF
}