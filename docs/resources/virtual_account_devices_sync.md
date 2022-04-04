---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_virtual_account_devices_sync Resource - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs create operation on Device Onboarding (PnP).
          - Synchronizes the device info from the given smart account & virtual account with the PnP database. The response
          payload returns a list of synced devices
---

# dnacenter_virtual_account_devices_sync (Resource)

It performs create operation on Device Onboarding (PnP).
		- Synchronizes the device info from the given smart account & virtual account with the PnP database. The response
		payload returns a list of synced devices



<!-- schema generated by tfplugindocs -->


~>**Warning,**
This resource does not represent a real-world entity in Cisco DNA Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco DNA Center workflow. It is executed in DNACenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Schema

### Required

- **parameters** (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--parameters))

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))
- **last_updated** (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- **auto_sync_period** (Number)
- **cco_user** (String)
- **expiry** (Number)
- **last_sync** (Number)
- **profile** (Block List) (see [below for nested schema](#nestedblock--parameters--profile))
- **smart_account_id** (String)
- **sync_result** (Block List) (see [below for nested schema](#nestedblock--parameters--sync_result))
- **sync_result_str** (String)
- **sync_start_time** (Number)
- **sync_status** (String)
- **tenant_id** (String)
- **token** (String)
- **virtual_account_id** (String)

<a id="nestedblock--parameters--profile"></a>
### Nested Schema for `parameters.profile`

Optional:

- **address_fqdn** (String)
- **address_ip_v4** (String)
- **cert** (String)
- **make_default** (String)
- **name** (String)
- **port** (Number)
- **profile_id** (String)
- **proxy** (String)


<a id="nestedblock--parameters--sync_result"></a>
### Nested Schema for `parameters.sync_result`

Optional:

- **sync_list** (Block List) (see [below for nested schema](#nestedblock--parameters--sync_result--sync_list))
- **sync_msg** (String)

<a id="nestedblock--parameters--sync_result--sync_list"></a>
### Nested Schema for `parameters.sync_result.sync_list`

Optional:

- **device_sn_list** (List of String)
- **sync_type** (String)




<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **auto_sync_period** (Number)
- **cco_user** (String)
- **expiry** (Number)
- **last_sync** (Number)
- **profile** (List of Object) (see [below for nested schema](#nestedobjatt--item--profile))
- **smart_account_id** (String)
- **sync_result** (List of Object) (see [below for nested schema](#nestedobjatt--item--sync_result))
- **sync_result_str** (String)
- **sync_start_time** (Number)
- **sync_status** (String)
- **tenant_id** (String)
- **token** (String)
- **virtual_account_id** (String)

<a id="nestedobjatt--item--profile"></a>
### Nested Schema for `item.profile`

Read-Only:

- **address_fqdn** (String)
- **address_ip_v4** (String)
- **cert** (String)
- **make_default** (String)
- **name** (String)
- **port** (Number)
- **profile_id** (String)
- **proxy** (String)


<a id="nestedobjatt--item--sync_result"></a>
### Nested Schema for `item.sync_result`

Read-Only:

- **sync_list** (List of Object) (see [below for nested schema](#nestedobjatt--item--sync_result--sync_list))
- **sync_msg** (String)

<a id="nestedobjatt--item--sync_result--sync_list"></a>
### Nested Schema for `item.sync_result.sync_list`

Read-Only:

- **device_sn_list** (List of String)
- **sync_type** (String)

