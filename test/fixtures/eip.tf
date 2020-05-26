resource "aws_eip" "external" {
  count = 5
  vpc   = true
}

output "external_elastic_ips" {
  value = aws_eip.external.*.id
}
