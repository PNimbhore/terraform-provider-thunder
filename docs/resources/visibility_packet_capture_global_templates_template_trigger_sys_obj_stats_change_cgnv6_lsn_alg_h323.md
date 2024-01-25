---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_alg_h323 Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_alg_h323: Configure triggers for cgnv6.lsn.alg.h323 object
  PLACEHOLDER
---

# thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_alg_h323 (Resource)

`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_alg_h323`: Configure triggers for cgnv6.lsn.alg.h323 object

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_alg_h323" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_alg_h323" {

  name = "test"
  trigger_stats_rate {
    threshold_exceeded_by = 5
    duration              = 60
    parse_error           = 1
    tcp_out_of_order_drop = 1
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

- `parse_error` (Number) Enable automatic packet-capture for Message Parse Error
- `tcp_out_of_order_drop` (Number) Enable automatic packet-capture for TCP Out-of-Order Drop
- `uuid` (String) uuid of the object


<a id="nestedblock--trigger_stats_rate"></a>
### Nested Schema for `trigger_stats_rate`

Optional:

- `duration` (Number) Time in seconds to look for the anomaly, default is 60
- `parse_error` (Number) Enable automatic packet-capture for Message Parse Error
- `tcp_out_of_order_drop` (Number) Enable automatic packet-capture for TCP Out-of-Order Drop
- `threshold_exceeded_by` (Number) Set the threshold to the number of times greater than the previous duration to start the capture, default is 5
- `uuid` (String) uuid of the object

