
data "archive_file" "send_token_lambda_zip" {
  type        = "zip"
  source_file = "${path.module}/save_token_lambda/lambda_function.py"
  output_path = "${path.module}/save_token_lambda/function.zip"
}

resource "aws_lambda_function" "send_token_to_dynamodb" {
  function_name = local.send_token_lambda_function_name
  handler       = "lambda_function.lambda_handler"
  runtime       = "python3.8"
  role          = aws_iam_role.save_token_lambda_exec.arn
  filename      = data.archive_file.send_token_lambda_zip.output_path
  source_code_hash = filebase64sha256(data.archive_file.send_token_lambda_zip.output_path)

}

resource "aws_iam_role" "save_token_lambda_exec" {
  name = local.save_token_lambda_exec_role_name

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

resource "aws_iam_role_policy" "save_token_lambda_exec_policy" {
  name = local.save_token_lambda_policy_name
  role = aws_iam_role.save_token_lambda_exec.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "dynamodb:*"
        ]
        Resource = aws_dynamodb_table.token_table.arn
      },
    ]
  })
}
