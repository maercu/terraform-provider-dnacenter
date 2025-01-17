
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.7-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_wireless_rf_profile" "example" {
  provider = dnacenter
  parameters {
    channel_width       = "20"
    default_rf_profile  = "false"
    enable_brown_field  = "false"
    enable_custom       = "true"
    enable_radio_type_a = "true"
    enable_radio_type_b = "true"
    enable_radio_type_c = "true"
    name                = "TestTerraform"
    radio_type_a_properties {

      data_rates           = "6,9,12,18,24,36,48,54"
      mandatory_data_rates = "6"
      max_power_level      = 30
      min_power_level      = -10
      parent_profile       = "LOW"
      power_threshold_v1   = -61
      radio_channels       = "36,40,44,48,52,56,60,64,149,153,157,161"
      rx_sop_threshold     = "LOW"
    }
    radio_type_b_properties {

      data_rates           = "9,12,18,24,36,48,54"
      mandatory_data_rates = "9"
      max_power_level      = 30
      min_power_level      = 7
      parent_profile       = "HIGH"
      power_threshold_v1   = -70
      radio_channels       = "1,6,11"
      rx_sop_threshold     = "MEDIUM"
    }
    radio_type_c_properties {
      parent_profile       = "HIGH"
      radio_channels       = "5,21,37,53,69,85,101,117,133,149,165,181,197,213,229"
      data_rates           = "6,9,12,18,24,36,48,54"
      mandatory_data_rates = "6,12,24"
      power_threshold_v1   = -70
      rx_sop_threshold     = "AUTO"
      min_power_level      = -10
      max_power_level      = 30
    }
  }
}
output "dnacenter_wireless_rf_profile_example" {
  value = dnacenter_wireless_rf_profile.example
}

