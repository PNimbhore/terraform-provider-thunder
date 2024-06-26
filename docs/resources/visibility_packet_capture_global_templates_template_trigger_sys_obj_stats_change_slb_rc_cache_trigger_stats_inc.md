---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_rc_cache_trigger_stats_inc Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_rc_cache_trigger_stats_inc: Configure stats to trigger packet capture on increment
  PLACEHOLDER
---

# thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_rc_cache_trigger_stats_inc (Resource)

`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_rc_cache_trigger_stats_inc`: Configure stats to trigger packet capture on increment

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_rc_cache_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_rc_cache_trigger_stats_inc" {

  name                  = "test"
  content_toobig        = 1
  content_toosmall      = 1
  entry_create_failures = 1
  rv_failure            = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `content_toobig` (Number) Enable automatic packet-capture for Policy Content Too Big
- `content_toosmall` (Number) Enable automatic packet-capture for Policy Content Too Small
- `entry_create_failures` (Number) Enable automatic packet-capture for Entry Create failures
- `rv_failure` (Number) Enable automatic packet-capture for Revalidation Failures
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


