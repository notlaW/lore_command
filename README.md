# lore_command

A set of serverless functions that act as a discord slash command. Written in golang, deployed with terraform

## Deployment Commands

### Terraform
```
$ cd ./tf
$ terrarform init
```

### ECR

```
$ terraform plan -target="module.ecr_repo"
$ terraform apply -target="module.ecr_repo"
```

### Docker
```
$ make repo-login
$ make build-container
$ make tag-latest
$ make push
```

### Lambda Infra

```
$ cd ./tf
$ terraform plan
$ terraform apply
```