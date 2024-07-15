resource "aws_iam_role" "sfn_role" {
  name = local.state_machine_role_name

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Principal = {
          Service = "states.amazonaws.com"
        }
        Effect = "Allow"
      },
    ]
  })
}

resource "aws_iam_role_policy" "sfn_policy" {
  name = local.state_machine_policy_name
  role = aws_iam_role.sfn_role.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect : "Allow",
        Action : [
          "lambda:InvokeFunction"
        ],
        Resource : [
          aws_lambda_function.item_generator.arn,
          aws_lambda_function.send_token_to_dynamodb.arn
        ]
      },
      {
        Effect : "Allow",
        Action : "dynamodb:PutItem",
        Resource : aws_dynamodb_table.token_table.arn
      }
    ]
  })
}