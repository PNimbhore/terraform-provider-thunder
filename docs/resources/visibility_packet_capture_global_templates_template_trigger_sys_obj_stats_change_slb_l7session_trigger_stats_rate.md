---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_l7session_trigger_stats_rate Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_l7session_trigger_stats_rate: Configure stats to trigger packet capture on increment rate
  PLACEHOLDER
---

# thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_l7session_trigger_stats_rate (Resource)

`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_l7session_trigger_stats_rate`: Configure stats to trigger packet capture on increment rate

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_l7session_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_l7session_trigger_stats_rate" {

  name                  = "test"
  conn_not_exist        = 1
  data_cb_failed        = 1
  duration              = 60
  err_cb_failed         = 1
  err_event             = 1
  hps_fwdreq_fail       = 1
  server_conn_failed    = 1
  server_select_fail    = 1
  threshold_exceeded_by = 5
  wbuf_cb_failed        = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `conn_not_exist` (Number) Enable automatic packet-capture for Conn does not exist
- `data_cb_failed` (Number) Enable automatic packet-capture for Data event callback fail
- `duration` (Number) Time in seconds to look for the anomaly, default is 60
- `err_cb_failed` (Number) Enable automatic packet-capture for Err event callback failed
- `err_event` (Number) Enable automatic packet-capture for Err event from TCP
- `hps_fwdreq_fail` (Number) Enable automatic packet-capture for Fwd req fail
- `server_conn_failed` (Number) Enable automatic packet-capture for Server connection failed
- `server_select_fail` (Number) Enable automatic packet-capture for Server selection fail
- `threshold_exceeded_by` (Number) Set the threshold to the number of times greater than the previous duration to start the capture, default is 5
- `uuid` (String) uuid of the object
- `wbuf_cb_failed` (Number) Enable automatic packet-capture for Wbuf event callback failed

### Read-Only

- `id` (String) The ID of this resource.


