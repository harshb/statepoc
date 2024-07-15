locals{
  dynamodb_table_name = "TokenTable"

  map_lambda_function_name = "MapGenerator"
  map_lambda_role_name = "map_lambda_exec_role"
  map_lambda_policy_name = "map_lambda_policy"

  send_token_lambda_function_name = "SendTokenToDynamoDB"
  save_token_lambda_exec_role_name = "save_token_lambda_exec_role"
  save_token_lambda_policy_name = "save_token_lambda_policy"

  state_machine_name = "MapDynamoDbStateMachine"
  state_machine_role_name = "map_dynamodb_sfn_execution_role"
  state_machine_policy_name = "map_dynamodb_sfn_policy"


}