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
resource "iris_dhcp_reservation" "dhcp1" {
  subnet    = "subnetId"
  mac       = "00:11:22:33:44:55"
  ipaddr    = "192.168.0.100"
  hostname  = "my_hostname"
}
```
## Authors

* Dalet (https://www.dalet.com/)
