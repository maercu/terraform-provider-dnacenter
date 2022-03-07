
resource "dnacenter_global_credential_http_read" "example" {
  provider = dnacenter
  parameters {
    secure             = "false"
    username           = "string"
    password           = "string"
    port               = 1
    description        = "string"
    comments           = "string"
    credential_type    = "string"
    instance_tenant_id = "string"
    instance_uuid      = "string"
    id                 = "string"
  }
}

output "dnacenter_global_credential_http_read_example" {
  value     = dnacenter_global_credential_http_read.example
  sensitive = true
}