terraform {
  required_providers {
    dnacenter = {
      version = "0.1.0-alpha"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

data "dnacenter_wireless_psk_override" "example" {
  provider = dnacenter
  payload {
    pass_phrase = "create"
    site        = "Global/San Francisco"
    ssid        = "test999_pop"
  }
}