variable "number_of_elastic_ips_to_create" {
  type = number
}

resource "aws_eip" "nat" {
  count = var.number_of_elastic_ips_to_create
  vpc   = true
}
