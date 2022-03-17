
provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_deploy_template_v1" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    force_push_template             = "false"
    is_composite                    = "false"
    member_template_deployment_info = []
    target_info {
      host_name             = "string"
      id                    = "string"
      params                = {
        key1   ="value1"
        key2   ="value2"
      }
      resource_params       = ["string"]
      type                  = "string"
      versioned_template_id = "string"
    }
    template_id = "string"
  }
}

output "dnacenter_deploy_template_v1_example" {
  value = dnacenter_deploy_template_v1.example
}

