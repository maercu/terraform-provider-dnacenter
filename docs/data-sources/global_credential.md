---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_global_credential Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Discovery.
  Returns global credential for the given credential sub typeReturns the credential sub type for the given Id
---

# dnacenter_global_credential (Data Source)

It performs read operation on Discovery.

- Returns global credential for the given credential sub type

- Returns the credential sub type for the given Id

## Example Usage

```terraform
data "dnacenter_global_credential" "example" {
  provider            = dnacenter
  credential_sub_type = "string"
  order               = "string"
  sort_by             = "string"
}

output "dnacenter_global_credential_example" {
  value = data.dnacenter_global_credential.example.items
}

data "dnacenter_global_credential" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_global_credential_example" {
  value = data.dnacenter_global_credential.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **credential_sub_type** (String) credentialSubType query parameter. Credential type as CLI / SNMPV2_READ_COMMUNITY / SNMPV2_WRITE_COMMUNITY / SNMPV3 / HTTP_WRITE / HTTP_READ / NETCONF
- **id** (String) id path parameter. Global Credential ID
- **order** (String) order query parameter.
- **sort_by** (String) sortBy query parameter.

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))
- **items** (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **response** (String)
- **version** (String)


<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- **comments** (String)
- **credential_type** (String)
- **description** (String)
- **id** (String)
- **instance_tenant_id** (String)
- **instance_uuid** (String)

