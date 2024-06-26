---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_global_trigger_stats_rate Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_global_trigger_stats_rate: Configure stats to trigger packet capture on increment rate
  PLACEHOLDER
---

# thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_global_trigger_stats_rate (Resource)

`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_global_trigger_stats_rate`: Configure stats to trigger packet capture on increment rate

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_global_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_global_trigger_stats_rate" {

  name                      = "test"
  duration                  = 60
  fullcone_creation_failure = 1
  threshold_exceeded_by     = 5
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `duration` (Number) Time in seconds to look for the anomaly, default is 60
- `fullcone_creation_failure` (Number) Enable automatic packet-capture for Full-Cone Creation Failure
- `threshold_exceeded_by` (Number) Set the threshold to the number of times greater than the previous duration to start the capture, default is 5
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


