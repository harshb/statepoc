token =
msg= 

set

{
"correlation_id": "123"
}

alias set='aws stepfunctions send-task-success --task-token $token --task-output "{\"key\": \"success\"}"'
alias set='aws stepfunctions send-task-success --task-token $token --task-output "{\"key\": $msg}"'

aws stepfunctions send-task-success --task-token $token --task-output "{\"key\": \"$msg\", \"correlation_id\": \"your_correlation_id_value\"}"