token =
msg= 

set

{
"template": "ecs"
}

alias set='aws stepfunctions send-task-success --task-token $token --task-output "{\"key\": \"success\"}"'
alias set='aws stepfunctions send-task-success --task-token $token --task-output "{\"key\": $msg}"'