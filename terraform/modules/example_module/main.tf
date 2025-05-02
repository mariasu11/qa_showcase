resource "aws_instance" "example" {
  count         = var.instance_count
  ami           = "ami-0abcdef1234567890"
  instance_type = "t3.micro"
}
