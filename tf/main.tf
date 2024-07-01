provider "aws" {
  region = var.region
  profile = var.profile
}
locals {
  step_function_execution_role_name = "wf-poc-step-function-role"
  sqs_queue_name = "wf-poc-queue"
}

//IAM Role for Step function
resource "aws_iam_role" "step_function_execution_role" {
  name               = local.step_function_execution_role_name
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "states.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "step_function_execution_role_policy" {
  name   = "MyPolicy"
  role   = aws_iam_role.step_function_execution_role.id
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": ["iam:*", "sqs:*"],
      "Resource": "*"
    }
  ]
}
EOF
}

//SQS Queue
resource "aws_sqs_queue" "my_queue" {
  name = local.sqs_queue_name
}