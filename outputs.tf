locals {
  outputs = {
    public_subnets        = aws_subnet.public
    private_subnets       = aws_subnet.private
    vpc                   = aws_vpc.vpc
    vpc_default_sg        = aws_default_security_group.vpc-default-sg
    private_zone          = aws_route53_zone.net0ps
    public_subdomain_zone = aws_route53_zone.subdomain
    private_routing_table = aws_default_route_table.private
    public_routing_table  = aws_route_table.public
    dummy_dependency      = null_resource.dummy_dependency
  }
}

output "public_subnets" {
  value       = var.enable ? aws_subnet.public.*.id : []
  description = "Public subnets"
}

output "private_subnets" {
  value       = var.enable ? aws_subnet.private.*.id : []
  description = "Private subnets"
}

output "vpc_id" {
  value       = var.enable ? aws_vpc.vpc[0].id : ""
  description = "VPC ID"
}

output "vpc_default_sg" {
  value       = var.enable ? aws_default_security_group.vpc-default-sg[0].id : ""
  description = "Default VPC security group"
}

output "net0ps_zone_id" {
  value       = var.enable ? aws_route53_zone.net0ps[0].zone_id : ""
  description = "Private hosted zone ID"
}

output "private_zone_id" {
  value       = var.enable ? aws_route53_zone.net0ps[0].zone_id : ""
  description = "Private hosted zone ID"
}

output "subdomain_zone_id" {
  value       = var.enable ? aws_route53_zone.subdomain[0].zone_id : ""
  description = "Subdomain hosted zone ID"
}

output "public_subdomain_zone_id" {
  value       = var.enable ? aws_route53_zone.subdomain[0].zone_id : ""
  description = "Subdomain hosted zone ID"
}

output "public_subdomain" {
  value       = var.enable ? var.subdomain : ""
  description = "Public subdomain name"
}

output "private_subdomain" {
  value       = var.enable ? aws_route53_zone.net0ps[0].name : ""
  description = "Private subdomain name"
}

output "vpc_private_routing_table_id" {
  value       = var.enable ? aws_default_route_table.private[0].id : ""
  description = "Private routing table ID"
}

output "vpc_public_routing_table_id" {
  value       = var.enable ? aws_route_table.public[0].id : ""
  description = "Public routing table ID"
}

output "depends_id" {
  value       = null_resource.dummy_dependency.id
  description = "Dependency ID"
}

output "outputs" {
  value       = local.outputs
  description = "All VPC outputs"
}
