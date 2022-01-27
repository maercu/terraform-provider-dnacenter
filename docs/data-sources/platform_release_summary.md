---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_platform_release_summary Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Platform Configuration.
  Provides information such as API version, mandatory core packages for installation or upgrade, optional packages,
  Cisco DNA Center name and version, supported direct updates, and tenant ID.
---

# dnacenter_platform_release_summary (Data Source)

It performs read operation on Platform Configuration.

- Provides information such as API version, mandatory core packages for installation or upgrade, optional packages,
Cisco DNA Center name and version, supported direct updates, and tenant ID.

## Example Usage

```terraform
data "dnacenter_platform_release_summary" "example" {
  provider = dnacenter
}

output "dnacenter_platform_release_summary_example" {
  value = data.dnacenter_platform_release_summary.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **core_packages** (List of String)
- **installed_version** (String)
- **name** (String)
- **packages** (List of String)
- **supported_direct_updates** (List of String)
- **system_version** (String)
- **tenant_id** (String)

