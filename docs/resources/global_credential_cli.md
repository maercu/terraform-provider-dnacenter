---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_global_credential_cli Resource - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It manages create, read and update operations on Discovery.
  Updates global CLI credentialsAdds global CLI credentialDeletes global credential for the given ID
---

# dnacenter_global_credential_cli (Resource)

It manages create, read and update operations on Discovery.

- Updates global CLI credentials

- Adds global CLI credential

- Deletes global credential for the given ID

## Example Usage

```terraform
resource "dnacenter_global_credential_cli" "example" {
  provider = dnacenter
  parameters {
    comments           = "string"
    credential_type    = "string"
    description        = "string"
    enable_password    = "string"
    id                 = "string"
    instance_tenant_id = "string"
    instance_uuid      = "string"
    password           = "string"
    username           = "string"
  }
}

output "dnacenter_global_credential_cli_example" {
  value     = dnacenter_global_credential_cli.example
  sensitive = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **parameters** (Block List, Min: 1, Max: 1) Array of RequestDiscoveryCreateCLICredentials (see [below for nested schema](#nestedblock--parameters))

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))
- **last_updated** (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- **comments** (String)
- **credential_type** (String)
- **description** (String)
- **enable_password** (String)
- **id** (String) The ID of this resource.
- **instance_tenant_id** (String)
- **instance_uuid** (String)
- **password** (String, Sensitive)
- **username** (String)


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **comments** (String)
- **credential_type** (String)
- **description** (String)
- **id** (String)
- **instance_tenant_id** (String)
- **instance_uuid** (String)

## Import

Import is supported using the following syntax:

```shell
terraform import dnacenter_global_credential.example "id:=string"
```