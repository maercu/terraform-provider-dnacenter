---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_business_sda_virtual_network_summary Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on SDA.
  Get Virtual Network Summary
---

# dnacenter_business_sda_virtual_network_summary (Data Source)

It performs read operation on SDA.

- Get Virtual Network Summary

## Example Usage

```terraform
data "dnacenter_business_sda_virtual_network_summary" "example" {
  provider            = dnacenter
  site_name_hierarchy = "string"
}

output "dnacenter_business_sda_virtual_network_summary_example" {
  value = data.dnacenter_business_sda_virtual_network_summary.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `site_name_hierarchy` (String) siteNameHierarchy query parameter. Complete fabric siteNameHierarchy Path

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `description` (String)
- `status` (String)
- `virtual_network_count` (Number)
- `virtual_network_summary` (List of Object) (see [below for nested schema](#nestedobjatt--item--virtual_network_summary))

<a id="nestedobjatt--item--virtual_network_summary"></a>
### Nested Schema for `item.virtual_network_summary`

Read-Only:

- `site_name_hierarchy` (String)
- `virtual_network_name` (String)

