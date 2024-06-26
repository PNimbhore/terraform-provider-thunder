---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_object_templates_interface_ethernet_tmpl_trigger_stats_rate Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_object_templates_interface_ethernet_tmpl_trigger_stats_rate: Configure stats to triggers packet capture on increment
  PLACEHOLDER
---

# thunder_visibility_packet_capture_object_templates_interface_ethernet_tmpl_trigger_stats_rate (Resource)

`thunder_visibility_packet_capture_object_templates_interface_ethernet_tmpl_trigger_stats_rate`: Configure stats to triggers packet capture on increment

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_interface_ethernet_tmpl_trigger_stats_rate" "thunder_visibility_packet_capture_object_templates_interface_ethernet_tmpl_trigger_stats_rate" {

  name                  = "test"
  collisions            = 1
  crc                   = 1
  duration              = 60
  giants                = 1
  giants_output         = 1
  input_errors          = 1
  output_errors         = 1
  runts                 = 1
  threshold_exceeded_by = 5
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `collisions` (Number) Enable automatic packet-capture for Collisions
- `crc` (Number) Enable automatic packet-capture for CRC
- `duration` (Number) Time in seconds to look for the anomaly, default is 60
- `giants` (Number) Enable automatic packet-capture for Giants
- `giants_output` (Number) Enable automatic packet-capture for Output Giants
- `input_errors` (Number) Enable automatic packet-capture for Input errors
- `output_errors` (Number) Enable automatic packet-capture for Output errors
- `runts` (Number) Enable automatic packet-capture for Runts
- `threshold_exceeded_by` (Number) Set the threshold to the number of times greater than the previous duration to start the capture, default is 5
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


