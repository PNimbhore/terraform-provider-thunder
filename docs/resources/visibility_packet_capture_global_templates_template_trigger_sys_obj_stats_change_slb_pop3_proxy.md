---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_pop3_proxy Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_pop3_proxy: Configure triggers for slb.pop3-proxy object
  PLACEHOLDER
---

# thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_pop3_proxy (Resource)

`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_pop3_proxy`: Configure triggers for slb.pop3-proxy object

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_pop3_proxy" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_pop3_proxy" {

  name = "test"
  trigger_stats_rate {
    threshold_exceeded_by = 5
    duration              = 60
    svrsel_fail           = 1
    no_route              = 1
    snat_fail             = 1
    line_too_long         = 1
    invalid_start_line    = 1
    unsupported_command   = 1
    bad_sequence          = 1
    rsv_persist_conn_fail = 1
    smp_v6_fail           = 1
    smp_v4_fail           = 1
    insert_tuple_fail     = 1
    cl_est_err            = 1
    ser_connecting_err    = 1
    server_response_err   = 1
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `trigger_stats_inc` (Block List, Max: 1) (see [below for nested schema](#nestedblock--trigger_stats_inc))
- `trigger_stats_rate` (Block List, Max: 1) (see [below for nested schema](#nestedblock--trigger_stats_rate))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--trigger_stats_inc"></a>
### Nested Schema for `trigger_stats_inc`

Optional:

- `bad_sequence` (Number) Enable automatic packet-capture for Bad Sequence
- `cl_est_err` (Number) Enable automatic packet-capture for Client EST state erro
- `insert_tuple_fail` (Number) Enable automatic packet-capture for Serv Sel insert tuple fail
- `invalid_start_line` (Number) Enable automatic packet-capture for invalid start line
- `line_too_long` (Number) Enable automatic packet-capture for line too long
- `no_route` (Number) Enable automatic packet-capture for no route failure
- `rsv_persist_conn_fail` (Number) Enable automatic packet-capture for Serv Sel Persist fail
- `ser_connecting_err` (Number) Enable automatic packet-capture for Serv CTNG state error
- `server_response_err` (Number) Enable automatic packet-capture for Serv RESP state error
- `smp_v4_fail` (Number) Enable automatic packet-capture for Serv Sel SMPv4 fail
- `smp_v6_fail` (Number) Enable automatic packet-capture for Serv Sel SMPv6 fail
- `snat_fail` (Number) Enable automatic packet-capture for source nat failure
- `svrsel_fail` (Number) Enable automatic packet-capture for Server selection failure
- `unsupported_command` (Number) Enable automatic packet-capture for Unsupported cmd
- `uuid` (String) uuid of the object


<a id="nestedblock--trigger_stats_rate"></a>
### Nested Schema for `trigger_stats_rate`

Optional:

- `bad_sequence` (Number) Enable automatic packet-capture for Bad Sequence
- `cl_est_err` (Number) Enable automatic packet-capture for Client EST state erro
- `duration` (Number) Time in seconds to look for the anomaly, default is 60
- `insert_tuple_fail` (Number) Enable automatic packet-capture for Serv Sel insert tuple fail
- `invalid_start_line` (Number) Enable automatic packet-capture for invalid start line
- `line_too_long` (Number) Enable automatic packet-capture for line too long
- `no_route` (Number) Enable automatic packet-capture for no route failure
- `rsv_persist_conn_fail` (Number) Enable automatic packet-capture for Serv Sel Persist fail
- `ser_connecting_err` (Number) Enable automatic packet-capture for Serv CTNG state error
- `server_response_err` (Number) Enable automatic packet-capture for Serv RESP state error
- `smp_v4_fail` (Number) Enable automatic packet-capture for Serv Sel SMPv4 fail
- `smp_v6_fail` (Number) Enable automatic packet-capture for Serv Sel SMPv6 fail
- `snat_fail` (Number) Enable automatic packet-capture for source nat failure
- `svrsel_fail` (Number) Enable automatic packet-capture for Server selection failure
- `threshold_exceeded_by` (Number) Set the threshold to the number of times greater than the previous duration to start the capture, default is 5
- `unsupported_command` (Number) Enable automatic packet-capture for Unsupported cmd
- `uuid` (String) uuid of the object


