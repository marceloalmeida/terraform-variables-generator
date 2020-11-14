# terraform-variables-generator

Simple Tool to Generate Variables file from Terraform Configuration. It will find all `*.tf` files in current directory, and generate variables.tf file. If you already have this file, it will ask to override it.

## Build

```bash
go build .
```

## Usage

```bash
./terraform-variables-generator
```

It will find all `*.tf` files in current directory, and generate variables.tf file. If you already have this file, it will ask to override it.

### Example

```text
resource "aws_vpc" "vpc" {
  cidr_block           = var.cidr
  enable_dns_hostnames = var.enable_dns_hostnames
  enable_dns_support   = var.enable_dns_support

  tags {
    Name = "${var.name}"
  }
}

resource "aws_internet_gateway" "vpc" {
  vpc_id = "${aws_vpc.vpc.id}"

  tags {
    Name = "${var.name}-igw"
  }
}
```

 Will generate

 ```hcl
 variable "ami" {}

variable "instance_type" {}

variable "cidr" {}

variable "enable_dns_hostnames" {}

variable "enable_dns_support" {}

variable "name" {}
 ```

## Tests

Run tests and linter

```bash
go test -v -race ./...
golint -set_exit_status $(go list ./...)
```
