
terraform {
  required_providers {
    dnacenter = {
      version = "0.2.0-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

resource "dnacenter_compliance" "example" {
  provider = dnacenter
  parameters {
    trigger_full = true
    categories   = ["PSIRT"]
    device_uuids = ["3eb928b8-2414-4121-ac35-1247e5d666a4"]
  }
}

output "dnacenter_compliance_example" {
  value = dnacenter_compliance.example
}