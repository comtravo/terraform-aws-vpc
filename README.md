# Comtravo's legacy Terraform AWS VPC module
# Do not use this legacy module

## Requirements

| Name | Version |
|------|---------|
| terraform | >= 0.12 |
| aws | ~> 2.0 |

## Providers

| Name | Version |
|------|---------|
| aws | ~> 2.0 |
| null | n/a |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| azs | Availability zones | `list(string)` | n/a | yes |
| cidr | CIDR | `string` | n/a | yes |
| enable | Enable or Disable the module | `bool` | n/a | yes |
| environment | environment | `string` | n/a | yes |
| replication\_factor | Number of subnets, routing tables, NAT gateways | `number` | n/a | yes |
| vpc\_name | VPC name | `string` | n/a | yes |
| depends\_id | For inter module dependencies | `string` | `""` | no |
| enable\_dns\_hostnames | Enable DNS hostnames | `bool` | `true` | no |
| enable\_dns\_support | Enable DNS support | `bool` | `true` | no |
| nat\_az\_number | Subnet number to deploy NAT gateway in | `number` | `0` | no |
| subdomain | Subdomain name | `string` | `""` | no |

## Outputs

| Name | Description |
|------|-------------|
| depends\_id | dependency ID |
| net0ps\_zone\_id | Private hosted zone ID |
| private\_subdomain | Private subdomain name |
| private\_subnets | private subnets |
| private\_zone\_id | Private hosted zone ID |
| public\_subdomain | Public subdomain name |
| public\_subdomain\_zone\_id | Subdomain hosted zone ID |
| public\_subnets | public subnets |
| subdomain\_zone\_id | Subdomain hosted zone ID |
| vpc\_default\_sg | Default VPC security group |
| vpc\_id | VPC ID |
| vpc\_private\_routing\_table\_id | Private routing table ID |
| vpc\_public\_routing\_table\_id | Public routing table ID |

