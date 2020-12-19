---
page_title: "dna_sda_fabric_port_assignment_for_access_point Resource - terraform-provider-dnacenter"
subcategory: "SDA"
description: |-
  The dna_sda_fabric_port_assignment_for_access_point resource allows you to assign a DNACenter SDA port for an accces point.
---

# Resource dna_sda_fabric_port_assignment_for_access_point

The dna_sda_fabric_port_assignment_for_access_point resource allows you to assign a DNACenter SDA port for an accces point.

## Example Usage

```hcl
resource "dna_sda_fabric_port_assignment_for_access_point" "response" {
  provider   = dnacenter
  site_name_hierarchy = assigment.site_name_hierarchy
  device_management_ip_address = assigment.device_management_ip_address
  interface_name = assigment.interface_name
  data_ip_address_pool_name = assigment.data_ip_address_pool_name
  voice_ip_address_pool_name = assigment.voice_ip_address_pool_name
  authenticate_template_name = assigment.authenticate_template_name
  scalable_group_name = assigment.scalable_group_name
}
```

## Argument Reference

- `site_name_hierarchy` - (Required) The assignment's site name hierarchy.
- `device_management_ip_address` - (Required) The assignment's device management ip address. If it's changed it forces the creation of a new resource.
- `interface_name` - (Required) The assignment's interface name. If it's changed it forces the creation of a new resource.
- `data_ip_address_pool_name` - (Required) The assignment's data ip address pool name.
- `voice_ip_address_pool_name` - (Required) The assignment's voice ip address pool name.
- `authenticate_template_name` - (Required) The assignment's authenticate template name.
- `scalable_group_name` - (Required) The assignment's scalable group name.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The assignment's updated time with format RFC850.