---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_icmp_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_icmp_stats: Statistics for the object icmp
  PLACEHOLDER
---

# thunder_cgnv6_icmp_stats (Data Source)

`thunder_cgnv6_icmp_stats`: Statistics for the object icmp

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_icmp_stats" "thunder_cgnv6_icmp_stats" {
}
output "get_cgnv6_icmp_stats" {
  value = ["${data.thunder_cgnv6_icmp_stats.thunder_cgnv6_icmp_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `icmp_bad_type` (Number) Bad Embedded ICMP Type
- `icmp_no_port_info` (Number) ICMP Port Info Not Included
- `icmp_no_session_drop` (Number) ICMP No Matching Session Drop
- `icmp_to_icmp` (Number) ICMP to ICMP Conversion
- `icmp_to_icmp_err` (Number) ICMP to ICMP Conversion Error
- `icmp_to_icmpv6` (Number) ICMP to ICMPv6 Conversion
- `icmp_to_icmpv6_err` (Number) ICMP to ICMPv6 Conversion Error
- `icmp_unknown_type` (Number) ICMP Unknown Type
- `icmpv6_bad_type` (Number) Bad Embedded ICMPv6 Type
- `icmpv6_no_port_info` (Number) ICMPv6 Port Info Not Included
- `icmpv6_no_session_drop` (Number) ICMPv6 No Matching Session Drop
- `icmpv6_to_icmp` (Number) ICMPv6 to ICMP Conversion
- `icmpv6_to_icmp_err` (Number) ICMPv6 to ICMP Conversion Error
- `icmpv6_to_icmpv6` (Number) ICMPv6 to ICMPv6 Conversion
- `icmpv6_to_icmpv6_err` (Number) ICMPv6 to ICMPv6 Conversion Error
- `icmpv6_unknown_type` (Number) ICMPv6 Unknown Type
- `known_drop46` (Number) NAT64 Reverse Known ICMP Drop
- `known_drop64` (Number) NAT64 Forward Known ICMPv6 Drop
- `midpoint_hop64` (Number) NAT64 Forward Unknown Source Drop
- `no_prefix_for_ipv446` (Number) NAT64 Reverse No Prefix Match for IPv4
- `unknown_drop46` (Number) NAT64 Reverse Known ICMPv6 Drop
- `unknown_drop64` (Number) NAT64 Forward Unknown ICMPv6 Drop


