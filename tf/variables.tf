variable "region" {
  description = "The region Terraform deploys your instance"
  default = "us-east-1"
}

variable "profile"{
  type = string
  default = "adfs"
}

variable "step_function_execution_role_name" {
  description = "The name of the IAM role"
  default     = "your_role_name_here" // replace with your role name
}
