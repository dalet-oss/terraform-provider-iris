# Terraform provider for Dalet Iris Network Appliance

This is a Terraform provider that lets you:
- provision DHCP static mappings on Dalet Iris instance

## Getting Started

In your `main.tf` file, specify the version you want to use:

```hcl
terraform {
  required_providers {
    iris = {
      source = "dalet-oss/iris"
    }
  }
}

provider "iris" {
  # Configuration options
}
```

And now run terraform init:

```
$ terraform init
```

### Provider configuration

```hcl
provider "iris" {
  uri      = "http://iris:port"
  token    = "iris_api_token"
}
```

### Resource configuration

```hcl
resource "iris_dhcp_subnet" "subnet1" {
  cidr      = "192.168.0.0/24"
  pool {
     first  = "192.168.0.10"
     last   = "192.168.0.20"
  }
  pool {
     first  = "192.168.0.30"
     last   = "192.168.0.40"
  }
}

resource "iris_dhcp_reservation" "dhcp1" {
  subnet    = iris_dhcp_subnet.subnet1.id
  mac       = "00:11:22:33:44:55"
  ipaddr    = "192.168.0.100"
  hostname  = "my_hostname"
}
```
## Authors

* Dalet (https://www.dalet.com/)
