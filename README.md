# terraform-variables-generator

Terraform versions support ![version](https://img.shields.io/badge/version-0.11.*-blue) ![version](https://img.shields.io/badge/version-0.12.*-blue) ![version](https://img.shields.io/badge/version-0.13.*-blue) ![Build Status](https://github.com/alexandrst88/terraform-variables-generator/workflows/release/badge.svg) [![Twitter](https://img.shields.io/twitter/url/https/twitter.com/AlexandrSt88.svg?style=social&label=Follow%20%40AlexandrSt88)](https://twitter.com/AlexandrSt88)


Simple Tool to Generate Variables file from Terraform Configuration. It will find all *.tf files in current directory, and generate variables.tf file. If you already have this file, it will ask to override it.

| Version | Supports |
|---------|----------|
| 0.11.*  |    yes   |
| 0.12.*  |    yes   |
| 0.13.*  |    yes   |


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

```hcl
resource "aws_vpc" "vpc" {
  cidr_block           = var.cidr
  enable_dns_hostnames = var.enable_dns_hostnames
  enable_dns_support   = var.enable_dns_support

  tags {
    Name = var.name
  }
}

resource "aws_internet_gateway" "vpc" {
  vpc_id = aws_vpc.vpc.id

  tags {
    Name = "${var.name}-igw"
  }
}
```

Will generate

```hcl
variable "ami" {
  description = ""
}

variable "instance_type" {
  description = ""
}

variable "cidr" {
  description = ""
}

variable "enable_dns_hostnames" {
  description = ""
}

variable "enable_dns_support" {
  description = ""
}

variable "name" {
  description = ""
}
```

## Tests

Run tests and linter

```bash
go test -v -race ./...
golint -set_exit_status $(go list ./...)
```

## TO DO

Move Locals and Variables to Single Interface
