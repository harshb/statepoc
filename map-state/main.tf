provider "aws" {
  region = var.region
  profile = var.profile
}
locals{
  queue_name = "my-queue"
  state_machine_name = "my-state-machine"

}

resource "aws_sqs_queue" "queue" {
  name = local.queue_name
}
resource "aws_iam_role" "iam_role" {
  name = "my-iam-role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "states.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "iam_role_policy" {
  name = "my-iam-role-policy"
  role = aws_iam_role.iam_role.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "sqs:SendMessage"
      ],
      "Resource": "${aws_sqs_queue.queue.arn}"
    }
  ]
}
EOF
}