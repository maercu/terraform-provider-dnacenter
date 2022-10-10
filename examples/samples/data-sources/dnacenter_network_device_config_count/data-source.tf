terraform {
  required_providers {
    dnacenter = {
      version = "1.0.8-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

data "dnacenter_network_device_config_count" "example" {
  provider = dnacenter
}

output "dnacenter_network_device_config_count_example" {
  value = data.dnacenter_network_device_config_count.example.item
}
