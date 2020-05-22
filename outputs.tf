output "public_subnets" {
  value = [aws_subnet.public.*.id]
  description = "public subnets"
}

output "private_subnets" {
  value = [aws_subnet.private.*.id]
  description = "private subnets"
}

output "vpc_id" {
  value = element(concat(aws_vpc.vpc.*.id, [""]), 0)
  description = "VPC ID"
}

output "vpc_default_sg" {
  value = element(
    concat(aws_default_security_group.vpc-default-sg.*.id, [""]),
    0,
  )
  description = "Default VPC security group"
}

output "net0ps_zone_id" {
  value = element(concat(aws_route53_zone.net0ps.*.zone_id, [""]), 0)
  description = "Private hosted zone ID"
}

output "private_zone_id" {
  value = element(concat(aws_route53_zone.net0ps.*.zone_id, [""]), 0)
  description = "Private hosted zone ID"
}

output "subdomain_zone_id" {
  value = element(concat(aws_route53_zone.subdomain.*.zone_id, [""]), 0)
  description = "Subdomain hosted zone ID"
}

output "public_subdomain_zone_id" {
  value = element(concat(aws_route53_zone.subdomain.*.zone_id, [""]), 0)
  description = "Subdomain hosted zone ID"
}

output "public_subdomain" {
  value = var.subdomain
  description = "Public subdomain name"
}

output "private_subdomain" {
  value = element(concat(aws_route53_zone.net0ps.*.name, [""]), 0)
  description = "Private subdomain name"
}

output "vpc_private_routing_table_id" {
  value = element(concat(aws_default_route_table.private.*.id, [""]), 0)
  description = "Private routing table ID"
}

output "vpc_public_routing_table_id" {
  value = element(concat(aws_route_table.public.*.id, [""]), 0)
  description = "Public routing table ID"
}

output "depends_id" {
  value = null_resource.dummy_dependency.id
  description = "dependency ID"
}

