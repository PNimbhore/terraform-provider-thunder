---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_global Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_global: Configure triggers for fw.global object
  PLACEHOLDER
---

# thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_global (Resource)

`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_global`: Configure triggers for fw.global object

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_global" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_global" {

  name = "test"
  trigger_stats_rate {
    threshold_exceeded_by     = 5
    duration                  = 60
    fullcone_creation_failure = 1
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

- `fullcone_creation_failure` (Number) Enable automatic packet-capture for Full-Cone Creation Failure
- `uuid` (String) uuid of the object


<a id="nestedblock--trigger_stats_rate"></a>
### Nested Schema for `trigger_stats_rate`

Optional:

- `duration` (Number) Time in seconds to look for the anomaly, default is 60
- `fullcone_creation_failure` (Number) Enable automatic packet-capture for Full-Cone Creation Failure
- `threshold_exceeded_by` (Number) Set the threshold to the number of times greater than the previous duration to start the capture, default is 5
- `uuid` (String) uuid of the object

