---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_server_radius_trigger_stats_rate Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_server_radius_trigger_stats_rate: Configure stats to trigger packet capture on increment rate
  PLACEHOLDER
---

# thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_server_radius_trigger_stats_rate (Resource)

`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_server_radius_trigger_stats_rate`: Configure stats to trigger packet capture on increment rate

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_server_radius_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_server_radius_trigger_stats_rate" {

  name                  = "test"
  accounting_failure    = 1
  authen_failure        = 1
  authorize_failure     = 1
  duration              = 60
  job_start_error       = 1
  other_error           = 1
  polling_control_error = 1
  request_dropped       = 1
  response_error        = 1
  response_failure      = 1
  response_timeout      = 1
  threshold_exceeded_by = 5
  timeout_error         = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `accounting_failure` (Number) Enable automatic packet-capture for Accounting Failure
- `authen_failure` (Number) Enable automatic packet-capture for Total Authentication Failure
- `authorize_failure` (Number) Enable automatic packet-capture for Total Authorization Failure
- `duration` (Number) Time in seconds to look for the anomaly, default is 60
- `job_start_error` (Number) Enable automatic packet-capture for Total Job Start Error
- `other_error` (Number) Enable automatic packet-capture for Total Other Error
- `polling_control_error` (Number) Enable automatic packet-capture for Total Polling Control Error
- `request_dropped` (Number) Enable automatic packet-capture for Total Dropped Request
- `response_error` (Number) Enable automatic packet-capture for Total Error Response
- `response_failure` (Number) Enable automatic packet-capture for Total Failure Response
- `response_timeout` (Number) Enable automatic packet-capture for Total Timeout Response
- `threshold_exceeded_by` (Number) Set the threshold to the number of times greater than the previous duration to start the capture, default is 5
- `timeout_error` (Number) Enable automatic packet-capture for Total Timeout
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


