resource "aws_sfn_state_machine" "state_machine" {
  name     = local.state_machine_name
  role_arn = aws_iam_role.sfn_role.arn
  definition = file("${path.module}/state_machine_definition.json")
}

