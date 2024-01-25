---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dns64_trigger_stats_inc Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dns64_trigger_stats_inc: Configure stats to trigger packet capture on increment
  PLACEHOLDER
---

# thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dns64_trigger_stats_inc (Resource)

`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dns64_trigger_stats_inc`: Configure stats to trigger packet capture on increment

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dns64_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dns64_trigger_stats_inc" {

  name          = "test"
  drop          = 1
  query_bad_pkt = 1
  resp_bad_pkt  = 1
  resp_bad_qr   = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `drop` (Number) Enable automatic packet-capture for Dropped
- `query_bad_pkt` (Number) Enable automatic packet-capture for Query Bad Packet
- `resp_bad_pkt` (Number) Enable automatic packet-capture for Response Bad Packet
- `resp_bad_qr` (Number) Enable automatic packet-capture for Response Bad Query
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

