## 1.1.7-beta (May 11, 2023)
BUGFIXES:
*  Resource dnacenter_pnp_device_site_claim generates wrong POST BODY when multiple config_parameters sections are present #175 

## 1.1.6-beta (May 09, 2023)
BUGFIXES:
* Updating Registry documentation, supports 2.3.5.3 DNA Center API Version.

## 1.1.5-beta (May 09, 2023)
FEATURES:
**New Resource:** `resourceConfigurationTemplateVersion`

BUGFIXES:
* Updating README, supports 2.3.5.3 DNA Center API Version.

## 1.1.4-beta (May 09, 2023)
BUGFIXES:
* dnacenter_site resource does not save site_id in state any more, however site_id is required for creating other resources #168
* Failure when refreshing state of dnacenter_global_credential_cli and dnacenter_global_credential_netconf resources #161 
## 1.1.3-beta (May 08, 2023)
BUGFIXES:
* Failure when refreshing state of dnacenter_global_credential_cli and dnacenter_global_credential_netconf resources #161
* Failure when creating resource dnacenter_configuration_template #163
* Failure when executing AddVnInFabric #167
* dnacenter_site resource does not save site_id in state any more, however site_id is required for creating other resources #168

## 1.1.2-beta (May 03, 2023)
BUGFIXES:
* dnacenter_reserve_ip_subpool requires id parameter #159.
* Mandatory parameter wireless_profile_name for dnacenter_wireless_profile resource #160
* failure when refreshing state of dnacenter_global_credential_cli and dnacenter_global_credential_netconf resources #161
* dnacenter_sda_fabric_authentication_profile - Provider produced inconsistent result after apply #162
* Provider produced inconsistent result after apply while creating dnacenter_global_pool #164

## 1.1.1-beta (April 27, 2023)
BUGFIXES:
* Adding to `provider.go` file following resources:
  - `dnacenter_sensor_test_delete`
  - `dnacenter_sensor_test_create`
  - `dnacenter_golden_image`
  - `dnacenter_deploy_template_v1`
  - `dnacenter_global_credential_snmpv3`
  - `dnacenter_global_credential_snmpv2_write_community`
  - `dnacenter_global_credential_snmpv2_read_community`
  - `dnacenter_global_credential_netconf`
  - `dnacenter_global_credential_http_write`
  - `dnacenter_global_credential_http_read`
  - `dnacenter_global_credential_cli`
  - `dnacenter_tag_membership`
* Adding to `provider.go` file following data sources:
  - `dnacenter_license_smart_account_details`
  - `dnacenter_golden_tag_image_details`
* `dnacenter_site` documentation issues fixed.
## 1.1.0-beta (April 12, 2023)
`dnacenter-go-sdk` version changes from `github.com/cisco-en-programmability/dnacenter-go-sdk/v4/sdk` to `github.com/cisco-en-programmability/dnacenter-go-sdk/v5/sdk`.(Cisco DNA Center's v2.3.5.3 API)
FEATURES:
* **New Data Source:** `data_source_accesspoint_configuration_details_by_task_id`
* **New Data Source:** `data_source_authentication_policy_servers`
* **New Data Source:** `data_source_deploy_template`
* **New Data Source:** `data_source_device_reboot_apreboot`
* **New Data Source:** `data_source_dnac_packages`
* **New Data Source:** `data_source_eox_status_device`
* **New Data Source:** `data_source_eox_status_summary`
* **New Data Source:** `data_source_event_email_config`
* **New Data Source:** `data_source_event_snmp_config`
* **New Data Source:** `data_source_event_syslog_config`
* **New Data Source:** `data_source_global_credential_v2`
* **New Data Source:** `data_source_golden_tag_image`
* **New Data Source:** `data_source_integration_settings_instances_itsm`
* **New Data Source:** `data_source_lan_automation_log_by_serial_number`
* **New Data Source:** `data_source_license_device`
* **New Data Source:** `data_source_network_device_user_defined_field`
* **New Data Source:** `data_source_network_v2`
* **New Data Source:** `data_source_role_permissions`
* **New Data Source:** `data_source_roles`
* **New Data Source:** `data_source_service_provider_v2`
* **New Data Source:** `data_source_user`
* **New Data Source:** `data_source_users_external_servers`
* **New Data Source:** `data_source_wireless_accesspoint_configuration_summary`
* **New Resource:** `resource_credential_to_site_by_siteid_create_v2`
* **New Resource:** `resource_device_reboot_apreboot`
* **New Resource:** `resource_event_syslog_config`
* **New Resource:** `resource_execute_suggested_actions_commands`
* **New Resource:** `resource_global_credential_v2`
* **New Resource:** `resource_golden_tag_image`
* **New Resource:** `resource_integration_settings_instances_itsm`
* **New Resource:** `resource_network_device_user_defined_field`
* **New Resource:** `resource_network_v2`
* **New Resource:** `resource_pnp_device_unclaim`
* **New Resource:** `resource_service_provider_v2`
* **New Resource:** `resource_sp_profile_delete_v2`
* **New Resource:** `resource_user`
* **New Resource:** `resource_wireless_accespoint_configuration`

## 1.0.19-beta (Mar 03, 2023)
UPGRADE NOTES:
* The go version of the provider was updated to 1.20, this due to the new prerequisites of terraform gorealeser, in which it is detailed that you must have a GO version of 1.18 or higher. Here are [gorealeaser docs](https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-release-publish).

BUGFIXES:
* Provider crash on dnacenter_wireless_provision_device_update #139 removing header params to avoid crash    
* Provider crash on dnacenter_wireless_provision_access_point #140 removing header params to avoid crash
* dnacenter_lan_automation_create does not accept any values in parameters block #141 updating documentation.

## 1.0.18-beta (Feb 02, 2023)
BUGFIXES:
* Issue #131 => dnacenter_pnp_device_import does not save device id to terraform state [Fixed]

## 1.0.17-beta (Jan 20, 2023)
BUGFIXES:
* Removed error behavior when externally deleting a resource.
* Updating examples.

## 1.0.16-beta (Jan 17, 2023)
BUGFIXES:
* Changing `response` answer on `pnp_device_import` resource, parameter `id` turns to `type_id`.
* Changing `response` answer and `request` schema on `pnp_device` resource, parameter `id` turns to `type_id`.
* Allowing more than one `ssid_details` block on `wireless_profile` resource.
* Nil pointer exception fixed on `wireless_provision_ssid_create_provision` resource.
* Corrected key name from `items` to `item` in `response` schema of `wireless_rf_profile` resource.

## 1.0.15-beta (Jan 10, 2023)
BUGFIXES:
* `Offset` and `Limit` query params change to `int`.

## 1.0.14-beta (Nov 30, 2022)
BUGFIXES:
* Resource `dnacenter_reserve_ip_subpool` expand functions fixed.

IMPROVEMENTS:
* Examples added.

## 1.0.13-beta (Nov 17, 2022)
BUGFIXES:
* Resource `dnacenter_network_create` expand functions fixed.

## 1.0.12-beta (Nov 16, 2022)
BUGFIXES:
* Resource `dnacenter_sda_fabric_border_device` change following parameters from `boolean` to `string`:
    - `connected_to_internet`
    - `border_with_external_connectivity`
* Resource `dnacenter_sda_fabric_border_device` fixed to not allow array at `payload.external_connectivity_settings.l3_handoff.virtual_network` parameter.
* Resource `dnacenter_network_create` fixed.

IMPROVEMENTS:
* Examples added.
* `go.mod` and `go.sum` were updated.
## 1.0.11-beta (Nov 14, 2022)
BUGFIXES:
* Resource `dnacenter_sda_fabric_border_device` fixed to allow array at `payload.external_connectivity_settings.l3_handoff.virtual_network` parameter.

## 1.0.10-beta (Oct 21, 2022)
BUGFIXES:
* Repariring set `item` on site resource.

## 1.0.9-beta (Oct 18, 2022)
BUGFIXES:
* Added `time.Sleep` for `DELETE_SITE`.

## 1.0.8-beta (Oct 10, 2022)
BUGFIXES:
* Resource `dnacenter_site` fixed to update name and delete.
* Data source `dnacenter_site` fixed to make request by site_id.
  
## 1.0.7-beta (July 18, 2022)
BUGFIXES:
* Resource `dnacenter_discovery destroy method` fixed to delete only the managed resource.
* Resource `dnacenter_sda_fabric_border_device` item updated.
* Data Source `dnacenter_sda_fabric_border_device` item and flatten updated.

## 1.0.6-beta (July 12, 2022)
BUGFIXES:
* `site_name_hierarchy` added to `resource_sda_virtual_network_ip_pool` on `destroy`.

## 1.0.5-beta (July 12, 2022)
BUGFIXES:
* `site_name_hierarchy` was added to `resource_sda_virtual_network_ip_pool`.
* `*[]dnacentersdkgo.RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL3HandoffVirtualNetwork` turns to `*dnacentersdkgo.RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL3HandoffVirtualNetwork`

IMPROVEMENTS:
* Documentation was updated.
* `go.mod` and `go.sum` were updated.
  
## 1.0.4-beta (July 11, 2022)
BUGFIXES:
* `site_name_hierarchy` was added to `resource_sda_virtual_network_ip_pool`.
* `resource_event_subscription`  changes `resourceItem("parameters")` to `resourceItem("parameters.0.payload")`.
* `resource_event_subscription_rest` changes `resourceItem("parameters")` to `resourceItem("parameters.0.payload")`.
* `resource_sda_fabric_authentication_profile` changes `resourceItem("parameters")` to `resourceItem("parameters.0.payload")`.
* `resource_sda_fabric_border_device` changes `resourceItem("parameters")` to `resourceItem("parameters.0.payload")`.
* `resource_sda_fabric_control_plane_device` changes `resourceItem("parameters")` to `resourceItem("parameters.0.payload")`.
* `resource_sda_fabric_edge_device` changes `resourceItem("parameters")` to `resourceItem("parameters.0.payload")`.
* `resource_wireless_provision_access_point` changes `resourceItem("parameters")` to `resourceItem("parameters.0.payload")`.
* `resource_wireless_provision_device_update` changes `resourceItem("parameters")` to `resourceItem("parameters.0.payload")`.

## 1.0.3-beta (July 11, 2022)
BUGFIXES:
* `resource_sda_virtual_network` is now able to create the resource.
* `resource_sda_fabric_border_device` is now able to create the resource.
* `resource_sda_fabric_control_plane_device` is now able to create the resource.
* `resource_sda_fabric_edge_device` is now able to create the resource.
* `resource_sda_port_assignment_for_access_point` is now able to create the resource.
* `resource_sda_port_assignment_for_user_device` is now able to create the resource.
* `resource_sda_provision_device` is now able to create the resource.
* `resource_sda_virtual_network_ip_pool` is now able to create the resource.
* `resource_sda_virtual_network_v2` is now able to create the resource.

## 1.0.2-beta (July 10, 2022)
BUGFIXES:
* `resource_sda_virtual_network` is now able to create the resource.
* `resource_site` is now able to read resource information.

IMPROVEMENTS:
* An example for resource `dnacenter_sda_virtual_network` was added.

## 1.0.1-beta (July 08, 2022)
BUG FIXES:
* Add `site_name_hierarchy`,`fabric_name`,`fabric_type`,`fabric_domain_type` parameters to `item` schema of `data_source_sda_fabric_site`
* `payload` for array request added in following resources: `resource_event_subscription`,`resource_event_subscription_rest`,`resource_sda_fabric_authentication_profile`,`resource_sda_fabric_control_plane_device`,`resource_sda_fabric_edge_device`,`resource_sda_fabric_site`,`resource_sda_virtual_network` 

IMPROVEMENTS:
* In `resource_event_subscription` a `CustomCall` was added for verify creation state.
* * `dnacenter_sda_fabric_site` examples were added.

REMOVED:
* `dnacenter_sda_fabric` examples were removed.

## 1.0.0-beta (June 17, 2022)

NOTES:
`dnacenter-go-sdk` version changes from `github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk` to `github.com/cisco-en-programmability/dnacenter-go-sdk/v4/sdk`.(Cisco DNA Center's v2.3.3.0 API)

FEATURES: 
* **New Data Source:** `data_source_buildings_planned_access_points`
* **New Data Source:** `data_source_event_config_connector_types`
* **New Data Source:** `data_source_interface`
* **New Data Source:** `data_source_lan_automation_count`
* **New Data Source:** `data_source_lan_automation_log`
* **New Data Source:** `data_source_lan_automation_status`
* **New Data Source:** `data_source_network_device_custom_prompt`
* **New Data Source:** `data_source_network_device_interface_neighbor`
* **New Data Source:** `data_source_planned_access_points`
* **New Data Source:** `data_source_transit_peer_network`
* **New Resource:** `resource_network_device_custom_prompt`
* **New Resource:** `resource_transit_peer_network`
* **New Resource:** `resource_assign_device_to_site`
* **New Resource:** `resource_pnp_device_authorize`
* **New Resource:** `resource_interface_operation_create`
* **New Resource:** `resource_interface_update`
* **New Resource:** `resource_lan_automation_create`
* **New Resource:** `resource_lan_automation_delete`
* **New Resource:** `resource_syslog_config_create`
* **New Resource:** `resource_syslog_config_update`
* **New Resource:** `resource_event_email_config_update`
* **New Resource:** `resource_event_email_config_create`
* **New Resource:** `resource_event_webhook_create`
* **New Resource:** `resource_event_webhook_update`
* **New Resource:** `resource_file_import`
* **New Resource:** `resource_global_credential_delete`
* **New Resource:** `resource_global_credential_update`
* **New Resource:** `resource_network_create`
* **New Resource:** `resource_network_device_update_role`
* **New Resource:** `resource_network_update`
* **New Resource:** `resource_pnp_device_config_preview`
* **New Resource:** `resource_pnp_server_profile_update`
* **New Resource:** `resource_pnp_virtual_account_add`
* **New Resource:** `resource_pnp_virtual_account_deregister`
  
BREAKING CHANGES:
* Data Source `data_source_disasterrecovery_system_operationstatus` has been removed.
* Data Source `data_source_disasterrecovery_system_status` has been removed.
* Data Source `data_source_endpoint_analytics_profiling_rules` has been removed.
* Data Source `data_source_profiling_rules_count` has been removed.
* Data Source `data_source_sda_fabric` has been removed.
* Data Source `data_source_site_design_floormap` has been removed.
* Data Source `data_source_threat_detail` has been removed.
* Data Source `data_source_threat_detail_count` has been removed.
* Data Source `data_source_threat_summary` has been removed.
* Resource `resource_authentication_import_certificate` has been removed.
* Resource `resource_authentication_import_certificate_p12` has been removed.
* Resource `resource_endpoint_analytics_profiling_rules` has been removed.
* Resource `resource_profiling_rules_in_bulk_create` has been removed.
* Resource `resource_sda_fabric` has been removed.
* Resource `resource_site_design_floormap` has been removed.

## 0.3.0 (June 16, 2022)

NOTES: 
Stable version with `github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk` (Cisco DNA Center 2.2.3.3 API)

## 0.3.0-beta (April 04, 2022)

FEATURES:
* **New Resource:** `resource_deploy_template_v1`
* **New Resource:** `resource_deploy_template`
* **New Resource:** `resource_pnp_device_site_claim`
* **New Resource:** `resource_pnp_device_import`
* **New Resource:** `resource_image_distribution`
* **New Resource:** `resource_image_device_activation`
* **New Resource:** `resource_app_policy_intent_create`
* **New Resource:** `resource_associate_site_to_network_profile`
* **New Resource:** `resource_authentication_import_certificate`
* **New Resource:** `resource_authentication_import_certificate_p12`
* **New Resource:** `resource_business_sda_wireless_controller_create`
* **New Resource:** `resource_business_sda_wireless_controller_delete`
* **New Resource:** `resource_command_runner_run_command`
* **New Resource:** `resource_compliance`
* **New Resource:** `resource_configuration_template_clone`
* **New Resource:** `resource_configuration_template_export_project.`
* **New Resource:** `resource_configuration_template_export_template.`
* **New Resource:** `resource_configuration_template_import_project`
* **New Resource:** `resource_device_configurations_export`
* **New Resource:** `resource_device_replacement_deploy`
* **New Resource:** `resource_disassociate_site_to_network_profile`
* **New Resource:** `resource_itsm_integration_events_retry`
* **New Resource:** `resource_network_device_export`
* **New Resource:** `resource_network_device_sync.`
* **New Resource:** `resource_nfv_provision`
* **New Resource:** `resource_profiling_rules_in_bulk_create`
* **New Resource:** `resource_sensor_test_create`
* **New Resource:** `resource_sensor_test_delete`
* **New Resource:** `resource_sensor_test_run`
* **New Resource:** `resource_sensor_test_template_duplicate`
* **New Resource:** `resource_sensor_test_template_edit`
* **New Resource:** `resource_site_assign_credential`
* **New Resource:** `resource_template_preview`
* **New Resource:** `resource_virtual_account_devices_sync`
* **New Resource:** `resource_wireless_provision_access_point.`
* **New Resource:** `resource_wireless_provision_device_create.`
* **New Resource:** `resource_wireless_provision_device_update.`
* **New Resource:** `resource_wireless_provision_ssid_create_provision`
* **New Resource:** `resource_wireless_provision_ssid_delete_reprovisio`

## 0.2.0-beta (March 01, 2022)

* **New Resource:** `resource_tag_membership`
* **New Resource:** `resource_path_trace`
* **New Resource:** `resource_golden_image`
* **New Resource:** `resource_swim_image_file`
* **New Resource:** `resource_global_credential_cli`
* **New Resource:** `resource_global_credential_http_read`
* **New Resource:** `resource_global_credential_http_write`
* **New Resource:** `resource_global_credential_netconf`
* **New Resource:** `resource_global_credential_snmpv2_read`
* **New Resource:** `resource_global_credential_snmpv2_write`
* **New Resource:** `resource_global_credential_snmpv3`
* **New Resource:** `resource_site`
* **New Resource:** `resource_service_provider`

IMPROVEMENTS:
* dnacenter/resource_wireless_enterprise_ssid: Add trigger for update passphrase
* dnacenter/resource_network_device: Add trigger for update role of a device
* dnacenter/resource_configuration_template: Add trigger for new versio released

BUG FIXES:
* Remove `use_api_gateway` and `use_csrf_token` configuration parameters from `provider`

## 0.1.0-beta.2 (February 08, 2022)

NOTES:
* Data Sources of type 'action' have been removed. Removed data Sources of type 'action' have been classified as unsafe by the team.
* Next 0.2.0-beta version will transform some of them to resources.

BREAKING CHANGES:

* Data Source of type 'action' `dnacenter_app_policy_intent_create` has been removed
* Data Source of type 'action' `dnacenter_associate_site_to_network_profile` has been removed
* Data Source of type 'action' `dnacenter_authentication_import_certificate_p12` has been removed
* Data Source of type 'action' `dnacenter_authentication_import_certificate` has been removed
* Data Source of type 'action' `dnacenter_business_sda_wireless_controller_create` has been removed
* Data Source of type 'action' `dnacenter_business_sda_wireless_controller_delete` has been removed
* Data Source of type 'action' `dnacenter_cli_credential_create` has been removed
* Data Source of type 'action' `dnacenter_cli_credential_update` has been removed
* Data Source of type 'action' `dnacenter_command_runner_run_command` has been removed
* Data Source of type 'action' `dnacenter_compliance_check_run` has been removed
* Data Source of type 'action' `dnacenter_configuration_template_clone` has been removed
* Data Source of type 'action' `dnacenter_configuration_template_create` has been removed
* Data Source of type 'action' `dnacenter_configuration_template_deploy_v2` has been removed
* Data Source of type 'action' `dnacenter_configuration_template_deploy` has been removed
* Data Source of type 'action' `dnacenter_configuration_template_export_project` has been removed
* Data Source of type 'action' `dnacenter_configuration_template_export_template` has been removed
* Data Source of type 'action' `dnacenter_configuration_template_import_project` has been removed
* Data Source of type 'action' `dnacenter_configuration_template_import_template` has been removed
* Data Source of type 'action' `dnacenter_configuration_template_version_create` has been removed
* Data Source of type 'action' `dnacenter_device_configurations_export` has been removed
* Data Source of type 'action' `dnacenter_device_credential_create` has been removed
* Data Source of type 'action' `dnacenter_device_credential_delete` has been removed
* Data Source of type 'action' `dnacenter_device_credential_update` has been removed
* Data Source of type 'action' `dnacenter_device_replacement_deploy` has been removed
* Data Source of type 'action' `dnacenter_disassociate_site_to_network_profile` has been removed
* Data Source of type 'action' `dnacenter_discovery_range_delete` has been removed
* Data Source of type 'action' `dnacenter_global_credential_delete` has been removed
* Data Source of type 'action' `dnacenter_global_credential_update` has been removed
* Data Source of type 'action' `dnacenter_golden_image_create` has been removed
* Data Source of type 'action' `dnacenter_golden_tag_image_delete` has been removed
* Data Source of type 'action' `dnacenter_http_read_credential_create` has been removed
* Data Source of type 'action' `dnacenter_http_read_credential_update` has been removed
* Data Source of type 'action' `dnacenter_http_write_credential_create` has been removed
* Data Source of type 'action' `dnacenter_http_write_credential_update` has been removed
* Data Source of type 'action' `dnacenter_itsm_integration_events_retry` has been removed
* Data Source of type 'action' `dnacenter_license_device_deregistration` has been removed
* Data Source of type 'action' `dnacenter_license_device_registration` has been removed
* Data Source of type 'action' `dnacenter_license_virtual_account_change` has been removed
* Data Source of type 'action' `dnacenter_netconf_credential_create` has been removed
* Data Source of type 'action' `dnacenter_netconf_credential_update` has been removed
* Data Source of type 'action' `dnacenter_network_create` has been removed
* Data Source of type 'action' `dnacenter_network_device_export` has been removed
* Data Source of type 'action' `dnacenter_network_device_sync` has been removed
* Data Source of type 'action' `dnacenter_network_device_update_role` has been removed
* Data Source of type 'action' `dnacenter_network_update` has been removed
* Data Source of type 'action' `dnacenter_nfv_provision_details` has been removed
* Data Source of type 'action' `dnacenter_nfv_provision` has been removed
* Data Source of type 'action' `dnacenter_path_trace_create` has been removed
* Data Source of type 'action' `dnacenter_path_trace_delete` has been removed
* Data Source of type 'action' `dnacenter_pnp_device_claim_to_site` has been removed
* Data Source of type 'action' `dnacenter_pnp_device_claim` has been removed
* Data Source of type 'action' `dnacenter_pnp_device_config_preview` has been removed
* Data Source of type 'action' `dnacenter_pnp_device_import` has been removed
* Data Source of type 'action' `dnacenter_pnp_device_reset` has been removed
* Data Source of type 'action' `dnacenter_pnp_device_unclaim` has been removed
* Data Source of type 'action' `dnacenter_pnp_server_profile_update` has been removed
* Data Source of type 'action' `dnacenter_pnp_virtual_account_add` has been removed
* Data Source of type 'action' `dnacenter_pnp_virtual_account_deregister` has been removed
* Data Source of type 'action' `dnacenter_pnp_virtual_account_devices_sync` has been removed
* Data Source of type 'action' `dnacenter_profiling_rules_in_bulk_create` has been removed
* Data Source of type 'action' `dnacenter_sensor_create` has been removed
* Data Source of type 'action' `dnacenter_sensor_delete` has been removed
* Data Source of type 'action' `dnacenter_sensor_test_run` has been removed
* Data Source of type 'action' `dnacenter_sensor_test_template_duplicate` has been removed
* Data Source of type 'action' `dnacenter_sensor_test_template_edit` has been removed
* Data Source of type 'action' `dnacenter_service_provider_create` has been removed
* Data Source of type 'action' `dnacenter_service_provider_profile_delete` has been removed
* Data Source of type 'action' `dnacenter_service_provider_update` has been removed
* Data Source of type 'action' `dnacenter_site_assign_credential` has been removed
* Data Source of type 'action' `dnacenter_site_assign_device` has been removed
* Data Source of type 'action' `dnacenter_site_create` has been removed
* Data Source of type 'action' `dnacenter_site_delete` has been removed
* Data Source of type 'action' `dnacenter_site_update` has been removed
* Data Source of type 'action' `dnacenter_snmpv2_read_community_credential_create` has been removed
* Data Source of type 'action' `dnacenter_snmpv2_read_community_credential_update` has been removed
* Data Source of type 'action' `dnacenter_snmpv2_write_community_credential_create` has been removed
* Data Source of type 'action' `dnacenter_snmpv2_write_community_credential_update` has been removed
* Data Source of type 'action' `dnacenter_snmpv3_credential_create` has been removed
* Data Source of type 'action' `dnacenter_snmpv3_credential_update` has been removed
* Data Source of type 'action' `dnacenter_swim_import_local` has been removed
* Data Source of type 'action' `dnacenter_swim_import_via_url` has been removed
* Data Source of type 'action' `dnacenter_swim_trigger_activation` has been removed
* Data Source of type 'action' `dnacenter_swim_trigger_distribution` has been removed
* Data Source of type 'action' `dnacenter_tag_member_create` has been removed
* Data Source of type 'action' `dnacenter_tag_member_delete` has been removed
* Data Source of type 'action' `dnacenter_tag_membership` has been removed
* Data Source of type 'action' `dnacenter_template_preview` has been removed
* Data Source of type 'action' `dnacenter_wireless_provision_access_point` has been removed
* Data Source of type 'action' `dnacenter_wireless_provision_device_create` has been removed
* Data Source of type 'action' `dnacenter_wireless_provision_device_update` has been removed
* Data Source of type 'action' `dnacenter_wireless_provision_ssid_create_provision` has been removed
* Data Source of type 'action' `dnacenter_wireless_provision_ssid_delete_reprovision` has been removed
## 0.1.0-beta.1 (February 01, 2022)

BUG FIXES:
* dnacenter/data_source_system_performance_historical: Remove `t1` computed parameter from schema.

IMPROVEMENTS:
* dnacenter/data_source_pnp_device_claim_to_site: Add parameters `config_info, image_info, hostname` to schema.

## 0.1.0-beta (January 27, 2022)

IMPROVEMENTS:
* Removed comments.

## 0.1.0-alpha (January 27, 2022)

* New Release, supports Cisco DNA Center API 2.2.3.3.

BREAKING CHANGES:
* Drop or renamed previous data sources and resources.

FEATURES:
* **New Data Source:** `dnacenter_reserve_ip_subpool`
* **New Data Source:** `dnacenter_site_assign_device`
* **New Data Source:** `dnacenter_wireless_rf_profile`
* **New Data Source:** `dnacenter_wireless_provision_device_create`
* **New Data Source:** `dnacenter_wireless_provision_device_update`
* **New Data Source:** `dnacenter_wireless_provision_access_point`
* **New Data Source:** `dnacenter_wireless_profile`
* **New Data Source:** `dnacenter_template_preview`
* **New Data Source:** `dnacenter_configuration_template_create`
* **New Data Source:** `dnacenter_configuration_template_project`
* **New Data Source:** `dnacenter_tag_membership`
* **New Data Source:** `dnacenter_tag_member`
* **New Data Source:** `dnacenter_tag`
* **New Data Source:** `dnacenter_service_provider_profile_delete`
* **New Data Source:** `dnacenter_snmp_properties`
* **New Data Source:** `dnacenter_site_update`
* **New Data Source:** `dnacenter_site_delete`
* **New Data Source:** `dnacenter_site_create`
* **New Data Source:** `dnacenter_site`
* **New Data Source:** `dnacenter_service_provider_create`
* **New Data Source:** `dnacenter_service_provider`
* **New Data Source:** `dnacenter_service_provider_update`
* **New Data Source:** `dnacenter_sensor_test_template_duplicate`
* **New Data Source:** `dnacenter_sensor_test_run`
* **New Data Source:** `dnacenter_sensor`
* **New Data Source:** `dnacenter_threat_summary`
* **New Data Source:** `dnacenter_threat_detail_count`
* **New Data Source:** `dnacenter_authentication_import_certificate`
* **New Data Source:** `dnacenter_authentication_import_certificate_p12`
* **New Data Source:** `dnacenter_threat_detail`
* **New Data Source:** `dnacenter_pnp_virtual_account_deregister`
* **New Data Source:** `dnacenter_pnp_server_profile_update`
* **New Data Source:** `dnacenter_pnp_virtual_account_add`
* **New Data Source:** `dnacenter_pnp_global_settings`
* **New Data Source:** `dnacenter_pnp_virtual_account_devices_sync`
* **New Data Source:** `dnacenter_pnp_device_unclaim`
* **New Data Source:** `dnacenter_pnp_device_config_preview`
* **New Data Source:** `dnacenter_pnp_device_claim_to_site`
* **New Data Source:** `dnacenter_pnp_device_reset`
* **New Data Source:** `dnacenter_pnp_device_import`
* **New Data Source:** `dnacenter_pnp_device_claim`
* **New Data Source:** `dnacenter_site_design_floormap`
* **New Data Source:** `dnacenter_nfv_profile`
* **New Data Source:** `dnacenter_nfv_provision_details`
* **New Data Source:** `dnacenter_network_create`
* **New Data Source:** `dnacenter_network_update`
* **New Data Source:** `dnacenter_network`
* **New Data Source:** `dnacenter_network_device_sync`
* **New Data Source:** `dnacenter_network_device_export`
* **New Data Source:** `dnacenter_network_device_update_role`
* **New Data Source:** `dnacenter_command_runner_run_command`
* **New Data Source:** `dnacenter_device_configurations_export`
* **New Data Source:** `dnacenter_network_device_list`
* **New Data Source:** `dnacenter_network_device`
* **New Data Source:** `dnacenter_itsm_integration_events_failed`
* **New Data Source:** `dnacenter_itsm_integration_events_retry`
* **New Data Source:** `dnacenter_swim_import_via_url`
* **New Data Source:** `dnacenter_swim_import_local`
* **New Data Source:** `dnacenter_swim_trigger_distribution`
* **New Data Source:** `dnacenter_swim_trigger_activation`
* **New Data Source:** `dnacenter_global_pool`
* **New Data Source:** `dnacenter_snmpv3_credential_create`
* **New Data Source:** `dnacenter_snmpv3_credential_update`
* **New Data Source:** `dnacenter_snmpv2_write_community_credential_create`
* **New Data Source:** `dnacenter_snmpv2_write_community_credential_update`
* **New Data Source:** `dnacenter_snmpv2_read_community_credential_create`
* **New Data Source:** `dnacenter_snmpv2_read_community_credential_update`
* **New Data Source:** `dnacenter_netconf_credential_create`
* **New Data Source:** `dnacenter_netconf_credential_update`
* **New Data Source:** `dnacenter_http_write_credential_create`
* **New Data Source:** `dnacenter_http_write_credential_update`
* **New Data Source:** `dnacenter_http_read_credential_create`
* **New Data Source:** `dnacenter_http_read_credential_update`
* **New Data Source:** `dnacenter_cli_credential_create`
* **New Data Source:** `dnacenter_cli_credential_update`
* **New Data Source:** `dnacenter_global_credential_update`
* **New Data Source:** `dnacenter_global_credential_delete`
* **New Data Source:** `dnacenter_path_trace`
* **New Data Source:** `dnacenter_path_trace_create`
* **New Data Source:** `dnacenter_path_trace_delete`
* **New Data Source:** `dnacenter_event_subscription_syslog`
* **New Data Source:** `dnacenter_event_subscription_rest`
* **New Data Source:** `dnacenter_event_subscription_email`
* **New Data Source:** `dnacenter_event_subscription`
* **New Data Source:** `dnacenter_wireless_dynamic_interface`
* **New Data Source:** `dnacenter_wireless_psk_override`
* **New Data Source:** `dnacenter_wireless_enterprise_ssid`
* **New Data Source:** `dnacenter_discovery_range_delete`
* **New Data Source:** `dnacenter_discovery`
* **New Data Source:** `dnacenter_device_replacement_deploy`
* **New Data Source:** `dnacenter_device_replacement`
* **New Data Source:** `dnacenter_device_credential_create`
* **New Data Source:** `dnacenter_device_credential_update`
* **New Data Source:** `dnacenter_device_credential`
* **New Data Source:** `dnacenter_device_credential_delete`
* **New Data Source:** `dnacenter_reports`
* **New Data Source:** `dnacenter_site_assign_credential`
* **New Data Source:** `dnacenter_compliance_check_run`
* **New Data Source:** `dnacenter_wireless_provision_ssid_delete_reprovision`
* **New Data Source:** `dnacenter_wireless_provision_ssid_create_provision`
* **New Data Source:** `dnacenter_sda_multicast`
* **New Data Source:** `dnacenter_sda_virtual_network_v2`
* **New Data Source:** `dnacenter_sda_provision_device`
* **New Data Source:** `dnacenter_sda_virtual_network_ip_pool`
* **New Data Source:** `dnacenter_sda_virtual_network`
* **New Data Source:** `dnacenter_sda_port_assignment_for_user_device`
* **New Data Source:** `dnacenter_sda_port_assignment_for_access_point`
* **New Data Source:** `dnacenter_sda_fabric_site`
* **New Data Source:** `dnacenter_sda_fabric`
* **New Data Source:** `dnacenter_sda_fabric_edge_device`
* **New Data Source:** `dnacenter_sda_fabric_control_plane_device`
* **New Data Source:** `dnacenter_sda_fabric_border_device`
* **New Data Source:** `dnacenter_sda_fabric_authentication_profile`
* **New Data Source:** `dnacenter_nfv_provision`
* **New Data Source:** `dnacenter_sensor_test_template_edit`
* **New Data Source:** `dnacenter_applications`
* **New Data Source:** `dnacenter_application_sets`
* **New Data Source:** `dnacenter_pnp_workflow`
* **New Data Source:** `dnacenter_pnp_device`
* **New Data Source:** `dnacenter_event_artifact_count`
* **New Data Source:** `dnacenter_event_artifact`
* **New Data Source:** `dnacenter_user_enrichment_details`
* **New Data Source:** `dnacenter_topology_vlan_details`
* **New Data Source:** `dnacenter_topology_site`
* **New Data Source:** `dnacenter_topology_physical`
* **New Data Source:** `dnacenter_topology_layer_3`
* **New Data Source:** `dnacenter_topology_layer_2`
* **New Data Source:** `dnacenter_configuration_template`
* **New Data Source:** `dnacenter_configuration_template_version`
* **New Data Source:** `dnacenter_configuration_template_version_create`
* **New Data Source:** `dnacenter_configuration_template_deploy`
* **New Data Source:** `dnacenter_configuration_template_deploy_status`
* **New Data Source:** `dnacenter_configuration_template_clone`
* **New Data Source:** `dnacenter_configuration_template_deploy_v2`
* **New Data Source:** `dnacenter_configuration_template_export_project`
* **New Data Source:** `dnacenter_configuration_template_export_template`
* **New Data Source:** `dnacenter_configuration_template_import_project`
* **New Data Source:** `dnacenter_configuration_template_import_template`
* **New Data Source:** `dnacenter_network_device_chassis_details`
* **New Data Source:** `dnacenter_network_device_linecard_details`
* **New Data Source:** `dnacenter_network_device_stack_details`
* **New Data Source:** `dnacenter_network_device_supervisor_card_details`
* **New Data Source:** `dnacenter_network_device_inventory_insight_link_mismatch`
* **New Data Source:** `dnacenter_network_device_with_snmp_v3_des`
* **New Data Source:** `dnacenter_network_device_interface_poe`
* **New Data Source:** `dnacenter_system_health`
* **New Data Source:** `dnacenter_system_health_count`
* **New Data Source:** `dnacenter_system_performance`
* **New Data Source:** `dnacenter_system_performance_historical`
* **New Data Source:** `dnacenter_platform_nodes_configuration_summary`
* **New Data Source:** `dnacenter_platform_release_summary`
* **New Data Source:** `dnacenter_task_operation`
* **New Data Source:** `dnacenter_task_count`
* **New Data Source:** `dnacenter_task_tree`
* **New Data Source:** `dnacenter_task`
* **New Data Source:** `dnacenter_tag_member_type`
* **New Data Source:** `dnacenter_tag_count`
* **New Data Source:** `dnacenter_tag_member_count`
* **New Data Source:** `dnacenter_site_count`
* **New Data Source:** `dnacenter_site_health`
* **New Data Source:** `dnacenter_security_advisories_per_device`
* **New Data Source:** `dnacenter_security_advisories_ids_per_device`
* **New Data Source:** `dnacenter_security_advisories_summary`
* **New Data Source:** `dnacenter_security_advisories_devices`
* **New Data Source:** `dnacenter_security_advisories`
* **New Data Source:** `dnacenter_pnp_workflow_count`
* **New Data Source:** `dnacenter_pnp_virtual_accounts`
* **New Data Source:** `dnacenter_pnp_smart_account_domains`
* **New Data Source:** `dnacenter_pnp_virtual_account_sync_result`
* **New Data Source:** `dnacenter_pnp_device_history`
* **New Data Source:** `dnacenter_pnp_device_count`
* **New Data Source:** `dnacenter_topology_network_health`
* **New Data Source:** `dnacenter_network_device_register_for_wsa`
* **New Data Source:** `dnacenter_network_device_by_serial_number`
* **New Data Source:** `dnacenter_network_device_module_count`
* **New Data Source:** `dnacenter_network_device_module`
* **New Data Source:** `dnacenter_network_device_by_ip`
* **New Data Source:** `dnacenter_network_device_functional_capability`
* **New Data Source:** `dnacenter_network_device_count`
* **New Data Source:** `dnacenter_network_device_config_count`
* **New Data Source:** `dnacenter_network_device_config`
* **New Data Source:** `dnacenter_network_device_global_polling_interval`
* **New Data Source:** `dnacenter_network_device_lexicographically_sorted`
* **New Data Source:** `dnacenter_network_device_range`
* **New Data Source:** `dnacenter_network_device_wireless_lan`
* **New Data Source:** `dnacenter_network_device_vlan`
* **New Data Source:** `dnacenter_network_device_meraki_organization`
* **New Data Source:** `dnacenter_network_device_polling_interval`
* **New Data Source:** `dnacenter_network_device_summary`
* **New Data Source:** `dnacenter_network_device_poe`
* **New Data Source:** `dnacenter_network_device_equipment`
* **New Data Source:** `dnacenter_dna_command_runner_keywords`
* **New Data Source:** `dnacenter_site_membership`
* **New Data Source:** `dnacenter_issues`
* **New Data Source:** `dnacenter_issues_enrichment_details`
* **New Data Source:** `dnacenter_device_interface_ospf`
* **New Data Source:** `dnacenter_interface_network_device_detail`
* **New Data Source:** `dnacenter_interface_network_device`
* **New Data Source:** `dnacenter_interface_network_device_range`
* **New Data Source:** `dnacenter_device_interface_isis`
* **New Data Source:** `dnacenter_device_interface_by_ip`
* **New Data Source:** `dnacenter_device_interface_count`
* **New Data Source:** `dnacenter_device_interface`
* **New Data Source:** `dnacenter_swim_image_details`
* **New Data Source:** `dnacenter_global_credential`
* **New Data Source:** `dnacenter_file_namespaces`
* **New Data Source:** `dnacenter_file_namespace_files`
* **New Data Source:** `dnacenter_file`
* **New Data Source:** `dnacenter_event_count`
* **New Data Source:** `dnacenter_event`
* **New Data Source:** `dnacenter_event_subscription_count`
* **New Data Source:** `dnacenter_event_subscription_details_syslog`
* **New Data Source:** `dnacenter_event_subscription_details_rest`
* **New Data Source:** `dnacenter_event_subscription_details_email`
* **New Data Source:** `dnacenter_event_series_count`
* **New Data Source:** `dnacenter_event_series`
* **New Data Source:** `dnacenter_event_api_status`
* **New Data Source:** `dnacenter_discovery_count`
* **New Data Source:** `dnacenter_discovery_range`
* **New Data Source:** `dnacenter_discovery_summary`
* **New Data Source:** `dnacenter_discovery_device_count`
* **New Data Source:** `dnacenter_discovery_device`
* **New Data Source:** `dnacenter_discovery_device_range`
* **New Data Source:** `dnacenter_discovery_job_by_id`
* **New Data Source:** `dnacenter_discovery_jobs`
* **New Data Source:** `dnacenter_device_replacement_count`
* **New Data Source:** `dnacenter_device_health`
* **New Data Source:** `dnacenter_device_enrichment_details`
* **New Data Source:** `dnacenter_device_details`
* **New Data Source:** `dnacenter_reports_view_group_view`
* **New Data Source:** `dnacenter_reports_view_group`
* **New Data Source:** `dnacenter_reports_executions_download`
* **New Data Source:** `dnacenter_reports_executions`
* **New Data Source:** `dnacenter_compliance_device_status_count`
* **New Data Source:** `dnacenter_compliance_device_details_count`
* **New Data Source:** `dnacenter_compliance_device_details`
* **New Data Source:** `dnacenter_compliance_device_by_id_detail`
* **New Data Source:** `dnacenter_compliance_device`
* **New Data Source:** `dnacenter_compliance_device_by_id`
* **New Data Source:** `dnacenter_itsm_cmdb_sync_status`
* **New Data Source:** `dnacenter_client_proximity`
* **New Data Source:** `dnacenter_client_health`
* **New Data Source:** `dnacenter_client_enrichment_details`
* **New Data Source:** `dnacenter_client_detail`
* **New Data Source:** `dnacenter_sda_count`
* **New Data Source:** `dnacenter_sda_device_role`
* **New Data Source:** `dnacenter_sda_device`
* **New Data Source:** `dnacenter_nfv_provision_detail`
* **New Data Source:** `dnacenter_wireless_sensor_test_results`
* **New Data Source:** `dnacenter_applications_count`
* **New Data Source:** `dnacenter_application_sets_count`
* **New Data Source:** `dnacenter_applications_health`
* **New Data Source:** `dnacenter_event_series_audit_logs`
* **New Data Source:** `dnacenter_event_series_audit_logs_summary`
* **New Data Source:** `dnacenter_event_series_audit_logs_parent_records`
* **New Data Source:** `dnacenter_license_virtual_account_change`
* **New Data Source:** `dnacenter_license_device_count`
* **New Data Source:** `dnacenter_license_device_deregistration`
* **New Data Source:** `dnacenter_license_device_license_details`
* **New Data Source:** `dnacenter_license_device_license_summary`
* **New Data Source:** `dnacenter_license_device_registration`
* **New Data Source:** `dnacenter_license_term_details`
* **New Data Source:** `dnacenter_license_usage_details`
* **New Data Source:** `dnacenter_license_smart_account_details`
* **New Data Source:** `dnacenter_license_virtual_account_details`
* **New Data Source:** `dnacenter_app_policy`
* **New Data Source:** `dnacenter_app_policy_default`
* **New Data Source:** `dnacenter_app_policy_intent_create`
* **New Data Source:** `dnacenter_app_policy_queuing_profile`
* **New Data Source:** `dnacenter_app_policy_queuing_profile_count`
* **New Data Source:** `dnacenter_business_sda_hostonboarding_ssid_ippool`
* **New Data Source:** `dnacenter_business_sda_wireless_controller_create`
* **New Data Source:** `dnacenter_business_sda_wireless_controller_delete`
* **New Data Source:** `dnacenter_disasterrecovery_system_operationstatus`
* **New Data Source:** `dnacenter_disasterrecovery_system_status`
* **New Data Source:** `dnacenter_dnacaap_management_execution_status`
* **New Data Source:** `dnacenter_endpoint_analytics_profiling_rules`
* **New Data Source:** `dnacenter_profiling_rules_in_bulk_create`
* **New Data Source:** `dnacenter_profiling_rules_count`
* **New Data Source:** `dnacenter_device_family_identifiers_details`
* **New Data Source:** `dnacenter_golden_image_create`
* **New Data Source:** `dnacenter_golden_tag_image_details`
* **New Data Source:** `dnacenter_golden_tag_image_delete`
* **New Data Source:** `dnacenter_associate_site_to_network_profile`
* **New Data Source:** `dnacenter_disassociate_site_to_network_profile`
* **New Data Source:** `dnacenter_qos_device_interface`
* **New Data Source:** `dnacenter_qos_device_interface_info_count`
* **New Data Source:** `dnacenter_projects_details`
* **New Data Source:** `dnacenter_templates_details`
* **New Data Source:** `dnacenter_sensor_create`
* **New Data Source:** `dnacenter_sensor_delete`
* **New Data Source:** `dnacenter_tag_member_create`
* **New Data Source:** `dnacenter_tag_member_delete`
* **New Resource:** `dnacenter_reserve_ip_subpool`
* **New Resource:** `dnacenter_wireless_rf_profile`
* **New Resource:** `dnacenter_wireless_profile`
* **New Resource:** `dnacenter_configuration_template_project`
* **New Resource:** `dnacenter_tag`
* **New Resource:** `dnacenter_snmp_properties`
* **New Resource:** `dnacenter_pnp_global_settings`
* **New Resource:** `dnacenter_site_design_floormap`
* **New Resource:** `dnacenter_nfv_profile`
* **New Resource:** `dnacenter_network_device_list`
* **New Resource:** `dnacenter_network_device`
* **New Resource:** `dnacenter_global_pool`
* **New Resource:** `dnacenter_event_subscription_syslog`
* **New Resource:** `dnacenter_event_subscription_rest`
* **New Resource:** `dnacenter_event_subscription_email`
* **New Resource:** `dnacenter_event_subscription`
* **New Resource:** `dnacenter_wireless_dynamic_interface`
* **New Resource:** `dnacenter_wireless_enterprise_ssid`
* **New Resource:** `dnacenter_discovery`
* **New Resource:** `dnacenter_device_replacement`
* **New Resource:** `dnacenter_reports`
* **New Resource:** `dnacenter_sda_multicast`
* **New Resource:** `dnacenter_sda_virtual_network_v2`
* **New Resource:** `dnacenter_sda_provision_device`
* **New Resource:** `dnacenter_sda_virtual_network_ip_pool`
* **New Resource:** `dnacenter_sda_virtual_network`
* **New Resource:** `dnacenter_sda_port_assignment_for_user_device`
* **New Resource:** `dnacenter_sda_port_assignment_for_access_point`
* **New Resource:** `dnacenter_sda_fabric_site`
* **New Resource:** `dnacenter_sda_fabric`
* **New Resource:** `dnacenter_sda_fabric_edge_device`
* **New Resource:** `dnacenter_sda_fabric_control_plane_device`
* **New Resource:** `dnacenter_sda_fabric_border_device`
* **New Resource:** `dnacenter_sda_fabric_authentication_profile`
* **New Resource:** `dnacenter_applications`
* **New Resource:** `dnacenter_application_sets`
* **New Resource:** `dnacenter_pnp_workflow`
* **New Resource:** `dnacenter_pnp_device`
* **New Resource:** `dnacenter_configuration_template`
* **New Resource:** `dnacenter_app_policy_queuing_profile`
* **New Resource:** `dnacenter_business_sda_hostonboarding_ssid_ippool`
* **New Resource:** `dnacenter_endpoint_analytics_profiling_rules`
* **New Resource:** `dnacenter_qos_device_interface`

## 0.0.4 (October 29, 2021)

BUG FIXES:

* dnacenter/data_source_dna_discovery_snmp_property_add: Change comparisson to verify if is present.
* dnacenter/data_source_dna_pnp_device_claim: Change comparisson to verify if is present.
* dnacenter/data_source_dna_pnp_device_reset: Change comparisson to verify if is present.
* dnacenter/data_source_dna_pnp_device_sync_vaact: Change comparisson to verify if is present.
* dnacenter/data_source_dna_template_deploy: Change comparisson to verify if is present.
* dnacenter/resource_dna_cli_credential: Change comparisson to verify if is present.
* dnacenter/resource_dna_discovery: Change comparisson to verify if is present.
* dnacenter/resource_dna_http_read_credential: Change comparisson to verify if is present.
* dnacenter/resource_dna_http_write_credential: Change comparisson to verify if is present.
* dnacenter/resource_dna_netconf_credential: Change comparisson to verify if is present.
* dnacenter/resource_dna_network: Change comparisson to verify if is present.
* dnacenter/resource_dna_network_global_ip_pool: Change comparisson to verify if is present.
* dnacenter/resource_dna_pnp_device: Change comparisson to verify if is present.
* dnacenter/resource_dna_pnp_global_settings: Change comparisson to verify if is present.
* dnacenter/resource_dna_pnp_workflow: Change comparisson to verify if is present.
* dnacenter/resource_dna_sda_fabric_border_device: Change comparisson to verify if is present.
* dnacenter/resource_dna_sda_fabric_ip_pool_in_vn: Change comparisson to verify if is present.
* dnacenter/resource_dna_sda_fabric_ip_pool_in_vn: Change comparisson to verify if is present.
* dnacenter/resource_dna_site: Change comparisson to verify if is present.
* dnacenter/resource_dna_snmpv2_read_community_credential: Change comparisson to verify if is present.
* dnacenter/resource_dna_snmpv2_write_community_credential: Change comparisson to verify if is present.
* dnacenter/resource_dna_snmpv3_credential: Change comparisson to verify if is present.
* dnacenter/resource_dna_tag: Change comparisson to verify if is present.
* dnacenter/resource_dna_template: Change comparisson to verify if is present.

## 0.0.3 (February 01, 2021)

* Initial Release, supports Cisco DNA Center API 2.1.1.

FEATURES:
* **New Data Source:** `dna_application_set`
* **New Data Source:** `dna_applications`
* **New Data Source:** `dna_applications_count`
* **New Data Source:** `dna_command_runner_keywords`
* **New Data Source:** `dna_command_runner_run_command`
* **New Data Source:** `dna_discovery_all_delete`
* **New Data Source:** `dna_discovery_count`
* **New Data Source:** `dna_discovery_device`
* **New Data Source:** `dna_discovery_device_count`
* **New Data Source:** `dna_discovery_device_range`
* **New Data Source:** `dna_discovery_job`
* **New Data Source:** `dna_discovery_jobs`
* **New Data Source:** `dna_discovery_range`
* **New Data Source:** `dna_discovery_range_delete`
* **New Data Source:** `dna_discovery_snmp_property`
* **New Data Source:** `dna_discovery_snmp_property_add`
* **New Data Source:** `dna_discovery_summary`
* **New Data Source:** `dna_global_credentials`
* **New Data Source:** `dna_network`
* **New Data Source:** `dna_network_device_credential`
* **New Data Source:** `dna_network_global_ip_pool`
* **New Data Source:** `dna_network_service_provider_profile`
* **New Data Source:** `dna_pnp_device`
* **New Data Source:** `dna_pnp_device_claim`
* **New Data Source:** `dna_pnp_device_claim_site`
* **New Data Source:** `dna_pnp_device_config_preview`
* **New Data Source:** `dna_pnp_device_count`
* **New Data Source:** `dna_pnp_device_history`
* **New Data Source:** `dna_pnp_device_reset`
* **New Data Source:** `dna_pnp_device_saact`
* **New Data Source:** `dna_pnp_device_sync_result_vacct`
* **New Data Source:** `dna_pnp_device_sync_vaact`
* **New Data Source:** `dna_pnp_device_unclaim`
* **New Data Source:** `dna_pnp_global_settings`
* **New Data Source:** `dna_pnp_vacct`
* **New Data Source:** `dna_pnp_workflow`
* **New Data Source:** `dna_pnp_workflow_count`
* **New Data Source:** `dna_site`
* **New Data Source:** `dna_site_count`
* **New Data Source:** `dna_site_health`
* **New Data Source:** `dna_site_membership`
* **New Data Source:** `dna_tag`
* **New Data Source:** `dna_tag_count`
* **New Data Source:** `dna_tag_member`
* **New Data Source:** `dna_tag_member_count`
* **New Data Source:** `dna_tag_member_type`
* **New Data Source:** `dna_task`
* **New Data Source:** `dna_template`
* **New Data Source:** `dna_template_deploy`
* **New Data Source:** `dna_template_deploy_status`
* **New Data Source:** `dna_template_details`
* **New Data Source:** `dna_template_preview`
* **New Data Source:** `dna_template_project`
* **New Data Source:** `dna_template_version`
* **New Resource:** `dna_application_set`
* **New Resource:** `dna_applications`
* **New Resource:** `dna_cli_credential`
* **New Resource:** `dna_discovery`
* **New Resource:** `dna_http_read_credential`
* **New Resource:** `dna_http_write_credential`
* **New Resource:** `dna_netconf_credential`
* **New Resource:** `dna_network`
* **New Resource:** `dna_network_credential_site_assignment`
* **New Resource:** `dna_network_global_ip_pool`
* **New Resource:** `dna_network_service_provider_profile`
* **New Resource:** `dna_pnp_device`
* **New Resource:** `dna_pnp_global_settings`
* **New Resource:** `dna_pnp_workflow`
* **New Resource:** `dna_sda_fabric`
* **New Resource:** `dna_sda_fabric_authentication_profile`
* **New Resource:** `dna_sda_fabric_border_device`
* **New Resource:** `dna_sda_fabric_control_plane_device`
* **New Resource:** `dna_sda_fabric_edge_device`
* **New Resource:** `dna_sda_fabric_ip_pool_in_vn`
* **New Resource:** `dna_sda_fabric_port_assignment_for_access_point`
* **New Resource:** `dna_sda_fabric_port_assignment_for_user_device`
* **New Resource:** `dna_sda_fabric_site`
* **New Resource:** `dna_sda_fabric_virtual_network`
* **New Resource:** `dna_site`
* **New Resource:** `dna_snmpv2_read_community_credential`
* **New Resource:** `dna_snmpv2_write_community_credential`
* **New Resource:** `dna_snmpv3_credential`
* **New Resource:** `dna_tag`
* **New Resource:** `dna_template`
* **New Resource:** `dna_template_project`
