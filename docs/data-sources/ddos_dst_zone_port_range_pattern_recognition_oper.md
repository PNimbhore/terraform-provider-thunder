---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_dst_zone_port_range_pattern_recognition_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_dst_zone_port_range_pattern_recognition_oper: Operational Status for the object pattern-recognition
  PLACEHOLDER
---

# thunder_ddos_dst_zone_port_range_pattern_recognition_oper (Data Source)

`thunder_ddos_dst_zone_port_range_pattern_recognition_oper`: Operational Status for the object pattern-recognition

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_zone_port_range_pattern_recognition_oper" "thunder_ddos_dst_zone_port_range_pattern_recognition_oper" {
  zone_name        = "test"
  port_range_end   = 22
  port_range_start = 20
  protocol         = "tcp"

}
output "get_ddos_dst_zone_port_range_pattern_recognition_oper" {
  value = ["${data.thunder_ddos_dst_zone_port_range_pattern_recognition_oper.thunder_ddos_dst_zone_port_range_pattern_recognition_oper}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `port_range_end` (String) PortRangeEnd
- `port_range_start` (String) PortRangeStart
- `protocol` (String) Protocol
- `zone_name` (String) ZoneName

### Optional

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--oper))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--oper"></a>
### Nested Schema for `oper`

Optional:

- `filter_count` (Number)
- `filter_list` (Block List) (see [below for nested schema](#nestedblock--oper--filter_list))
- `filter_threshold` (Number)
- `peace_pkt_count` (Number)
- `state` (String)
- `timestamp` (String)
- `war_pkt_count` (Number)
- `war_pkt_percentage` (Number)

<a id="nestedblock--oper--filter_list"></a>
### Nested Schema for `oper.filter_list`

Optional:

- `filter_desc` (String)
- `filter_enabled` (Number)
- `filter_expr` (String)
- `hardware_filter` (Number)
- `processing_unit` (String)
- `sample_ratio` (Number)

