resource "aws_dynamodb_table" "token_table" {
  name           = local.dynamodb_table_name
  billing_mode   = "PROVISIONED"
  read_capacity  = 5
  write_capacity = 5
  hash_key       = "TaskToken"

  attribute {
    name = "TaskToken"
    type = "S"
  }
}
