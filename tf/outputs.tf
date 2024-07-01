output "aws_iam_role" {
  description = "The IAM role for the step function"
  value       = aws_iam_role.step_function_execution_role.arn
}

output "aws_iam_role_policy" {
  description = "The IAM role policy for the step function"
  value       = aws_iam_role_policy.step_function_execution_role_policy.id
}

output "aws_sqs_queue" {
  description = "The SQS queue for the workflow"
  value       = aws_sqs_queue.my_queue.url
}