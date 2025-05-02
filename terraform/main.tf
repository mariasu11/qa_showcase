provider "aws" {
  region = "us-west-2"
}

module "example" {
  source         = "./modules/example_module"
  instance_count = 2
}
