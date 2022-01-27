---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_threat_summary Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs create operation on Devices.
  The Threat Summary for the Rogues and aWIPS
---

# dnacenter_threat_summary (Data Source)

It performs create operation on Devices.

- The Threat Summary for the Rogues and aWIPS

## Example Usage

```terraform
data "dnacenter_threat_summary" "example" {
  provider     = dnacenter
  end_time     = 1
  site_id      = ["string"]
  start_time   = 1
  threat_level = ["string"]
  threat_type  = ["string"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **end_time** (Number) End Time
- **id** (String) The ID of this resource.
- **site_id** (List of String) Site Id
- **start_time** (Number) Start Time
- **threat_level** (List of String) Threat Level
- **threat_type** (List of String) Threat Type

### Read-Only

- **items** (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- **threat_data** (List of Object) (see [below for nested schema](#nestedobjatt--items--threat_data))
- **timestamp** (Number)

<a id="nestedobjatt--items--threat_data"></a>
### Nested Schema for `items.threat_data`

Read-Only:

- **threat_count** (Number)
- **threat_level** (String)
- **threat_type** (String)

