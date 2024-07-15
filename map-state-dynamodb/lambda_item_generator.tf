
data "archive_file" "lambda_zip" {
  type        = "zip"
  source_file = "${path.module}/map_lambda/lambda_function.py"
  output_path = "${path.module}/map_lambda/function.zip"
}


resource "aws_lambda_function" "item_generator" {
  function_name =  local.map_lambda_function_name
  handler       = "lambda_function.lambda_handler"
  runtime       = "python3.8"

  role          = aws_iam_role.lambda_exec.arn

  filename      = data.archive_file.lambda_zip.output_path
  source_code_hash = filebase64sha256(data.archive_file.lambda_zip.output_path)


}

resource "aws_iam_role" "lambda_exec" {
  name = local.map_lambda_role_name

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
resource "aws_iam_role_policy" "lambda_cust_policy" {
  name = local.map_lambda_policy_name
  role = aws_iam_role.lambda_exec.id

  policy = <<-EOF
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Action": [
          "logs:*","s3:*","dynamodb:*","cloudwatch:**","sns:*","lambda:*"
        ],
        "Effect": "Allow",
        "Resource": "*"
      }
    ]
  }
  EOF
}

