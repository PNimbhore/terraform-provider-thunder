---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_rad_server_trigger_stats_inc Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_rad_server_trigger_stats_inc: Configure stats to trigger packet capture on increment
  PLACEHOLDER
---

# thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_rad_server_trigger_stats_inc (Resource)

`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_rad_server_trigger_stats_inc`: Configure stats to trigger packet capture on increment

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_rad_server_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_rad_server_trigger_stats_inc" {

  name                        = "test"
  ha_standby_dropped          = 1
  invalid_key                 = 1
  ipv6_prefix_length_mismatch = 1
  radius_request_dropped      = 1
  radius_table_full           = 1
  request_bad_secret_dropped  = 1
  request_ignored             = 1
  request_malformed_dropped   = 1
  request_no_key_vap_dropped  = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `ha_standby_dropped` (Number) Enable automatic packet-capture for HA Standby Dropped
- `invalid_key` (Number) Enable automatic packet-capture for Radius Request has Invalid Key Field
- `ipv6_prefix_length_mismatch` (Number) Enable automatic packet-capture for Framed IPV6 Prefix Length Mismatch
- `radius_request_dropped` (Number) Enable automatic packet-capture for RADIUS Request Dropped (Malformed Packet)
- `radius_table_full` (Number) Enable automatic packet-capture for RADIUS Request Dropped (Table Full)
- `request_bad_secret_dropped` (Number) Enable automatic packet-capture for RADIUS Request Bad Secret Dropped
- `request_ignored` (Number) Enable automatic packet-capture for RADIUS Request Table Full Dropped
- `request_malformed_dropped` (Number) Enable automatic packet-capture for RADIUS Request Malformed Dropped
- `request_no_key_vap_dropped` (Number) Enable automatic packet-capture for RADIUS Request No Key Attribute Dropped
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

