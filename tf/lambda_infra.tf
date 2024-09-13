variable "env_name" {
  description = "Environment name"
}

data "aws_ecr_repository" "lore_command_ecr_repo" {
  name = "lore_command"
}

resource "aws_lambda_function" "lore_command_function" {
  function_name = "lore_command"
  timeout       = 5 # seconds
  image_uri     = "${data.aws_ecr_repository.lore_command_ecr_repo.repository_url}:latest"
  package_type  = "Image"

  architectures = ["arm64"]

  role = aws_iam_role.lore_command_function_role.arn

  environment {
    variables = {
      ENVIRONMENT = var.env_name
    }
  }
}

resource "aws_iam_role" "lore_command_function_role" {
  name = "lore_command"

  assume_role_policy = jsonencode({
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      },
    ]
  })
}