provider "aws" {
  s3_force_path_style         = true
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
  access_key                  = "This is not an actual access key."
  secret_key                  = "This is not an actual secret key."

  endpoints {
    ec2     = "http://localstack:4597"
    iam     = "http://localstack:4593"
    route53 = "http://localstack:4580"
    sts     = "http://localstack:4592"
  }
}

data "aws_caller_identity" "current" {}
data "aws_region" "current" {}
data "aws_availability_zones" "available" {}


module "vpc_enabled" {
  source = "../"

  enable             = true
  vpc_name           = terraform.workspace
  subdomain          = "foo.bar.baz"
  cidr               = "10.10.0.0/16"
  azs                = data.aws_availability_zones.available.names
  nat_az_number      = 1
  environment        = terraform.workspace
  replication_factor = 3
}

output "vpc_enabled_output" {
  value       = module.vpc_enabled.outputs
  description = "vpc_enabled output"
}
