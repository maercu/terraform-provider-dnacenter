---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_tag_membership Resource - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It manages create, read and delete operations on Tag member.
  •   Adds members to the tag specified by id
  •   Removes Tag member from the tag specified by id
---

# dnacenter_tag_membership (Resource)

It manages create, read and delete operations on Tag member.

	•	Adds members to the tag specified by id
	•	Removes Tag member from the tag specified by id

## Example Usage

```terraform
resource "dnacenter_tag_membership" "example" {
  provider = dnacenter
  parameters {
    tag_id      = "string"
    member_type = "string"
    member_id   = "string"

  }
}

output "dnacenter_tag_membership_example" {
  value = dnacenter_tag_membership.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **parameters** (Block List, Min: 1, Max: 1) Array of RequestApplicationPolicyCreateApplicationSet (see [below for nested schema](#nestedblock--parameters))

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))
- **last_updated** (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Required:

- **member_id** (String) memberId path parameter. TagMember id to be removed from tag
- **member_type** (String)
- **tag_id** (String) id path parameter. Tag ID


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **instance_uuid** (String)

## Import

Import is supported using the following syntax:

```shell
terraform import dnacenter_tag_membership.example "id:=string"
```