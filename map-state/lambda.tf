
data "archive_file" "lambda_zip" {
  type        = "zip"
  source_file = "${path.module}/lambda/generator.py"
  output_path = "${path.module}/lambda/function.zip"
}

resource "aws_lambda_function" "item_generator" {
  function_name = "itemGenerator"
  handler       = "lambda_function.lambda_handler"
  runtime       = "python3.8"

  role          = aws_iam_role.lambda_exec.arn

  filename      = data.archive_file.lambda_zip.output_path

  source_code_hash = filebase64sha256(data.archive_file.lambda_zip.output_path)

  environment {
    variables = {
      foo = "bar"
    }
  }
}

resource "aws_iam_role" "lambda_exec" {
  name = "lambda_exec_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
        Effect = "Allow"
      },
    ]
  })
}

# resource "aws_iam_policy" "lambda_exec_policy" {
#   name        = "lambda_exec_policy"
#   description = "Policy for allowing Lambda function to access necessary services"
#
#   policy = <<EOF
# {
#   "Version": "2012-10-17",
#   "Statement": [
#     {
#       "Effect": "Allow",
#       "Action": [
#         "logs:CreateLogGroup",
#         "logs:CreateLogStream",
#         "logs:PutLogEvents"
#       ],
#       "Resource": "arn:aws:logs:*:*:*"
#     },
#     {
#       "Effect": "Allow",
#       "Action": [
#         "s3:GetObject",
#         "s3:PutObject"
#       ],
#       "Resource": "arn:aws:s3:::your_bucket_name/*"
#     }
#   ]
# }
# EOF
# }
#
# resource "aws_iam_role_policy_attachment" "lambda_exec_policy_attachment" {
#   role       = aws_iam_role.lambda_exec.name
#   policy_arn = aws_iam_policy.lambda_exec_policy.arn
# }