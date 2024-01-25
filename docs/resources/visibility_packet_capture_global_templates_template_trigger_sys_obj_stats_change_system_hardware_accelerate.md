---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_hardware_accelerate Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_hardware_accelerate: Configure triggers for system.hardware-accelerate object
  PLACEHOLDER
---

# thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_hardware_accelerate (Resource)

`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_hardware_accelerate`: Configure triggers for system.hardware-accelerate object

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_hardware_accelerate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_hardware_accelerate" {

  name = "test"
  trigger_stats_rate {
    threshold_exceeded_by        = 5
    duration                     = 60
    hw_fwd_prog_errors           = 1
    hw_fwd_flow_singlebit_errors = 1
    hw_fwd_flow_tag_mismatch     = 1
    hw_fwd_flow_seq_mismatch     = 1
    hw_fwd_flow_error_count      = 1
    hw_fwd_flow_unalign_count    = 1
    hw_fwd_flow_underflow_count  = 1
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

- `hw_fwd_flow_error_count` (Number) Enable automatic packet-capture for Hardware forward flow error count
- `hw_fwd_flow_seq_mismatch` (Number) Enable automatic packet-capture for Hardware forward sequence mismatch errors
- `hw_fwd_flow_singlebit_errors` (Number) Enable automatic packet-capture for Hardware forward singlebit Errors
- `hw_fwd_flow_tag_mismatch` (Number) Enable automatic packet-capture for Hardware forward tag mismatch errors
- `hw_fwd_flow_unalign_count` (Number) Enable automatic packet-capture for Hardware forward flow unalign count
- `hw_fwd_flow_underflow_count` (Number) Enable automatic packet-capture for Hardware forward flow underflow count
- `hw_fwd_prog_errors` (Number) Enable automatic packet-capture for Hardware forward programming Errors
- `uuid` (String) uuid of the object


<a id="nestedblock--trigger_stats_rate"></a>
### Nested Schema for `trigger_stats_rate`

Optional:

- `duration` (Number) Time in seconds to look for the anomaly, default is 60
- `hw_fwd_flow_error_count` (Number) Enable automatic packet-capture for Hardware forward flow error count
- `hw_fwd_flow_seq_mismatch` (Number) Enable automatic packet-capture for Hardware forward sequence mismatch errors
- `hw_fwd_flow_singlebit_errors` (Number) Enable automatic packet-capture for Hardware forward singlebit Errors
- `hw_fwd_flow_tag_mismatch` (Number) Enable automatic packet-capture for Hardware forward tag mismatch errors
- `hw_fwd_flow_unalign_count` (Number) Enable automatic packet-capture for Hardware forward flow unalign count
- `hw_fwd_flow_underflow_count` (Number) Enable automatic packet-capture for Hardware forward flow underflow count
- `hw_fwd_prog_errors` (Number) Enable automatic packet-capture for Hardware forward programming Errors
- `threshold_exceeded_by` (Number) Set the threshold to the number of times greater than the previous duration to start the capture, default is 5
- `uuid` (String) uuid of the object

