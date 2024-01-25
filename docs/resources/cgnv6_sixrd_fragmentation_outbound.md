---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_sixrd_fragmentation_outbound Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_sixrd_fragmentation_outbound: sixrd fragmentation rules for outbound oversize packets (default: ipv6)
  PLACEHOLDER
---

# thunder_cgnv6_sixrd_fragmentation_outbound (Resource)

`thunder_cgnv6_sixrd_fragmentation_outbound`: sixrd fragmentation rules for outbound oversize packets (default: ipv6)

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_sixrd_fragmentation_outbound" "thunder_cgnv6_sixrd_fragmentation_outbound" {

  action = "drop"
  df_set = "send-icmp"

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `action` (String) 'drop': Drop Silently; 'ipv6': Use IPv6 Fragmentation for oversize packets (default); 'send-icmp': Send ICMP Type 3 Code 4 (Fragmentation Needed and DF Set); 'send-icmpv6': Send ICMP Type 2 Code 0 (Packet Too Big);
- `count1` (Number) Configure number of ICMP messages sent when DF set. Default is 1
- `df_set` (String) 'drop': Drop Silently; 'ipv6': Use IPv6 Fragmentation for oversize packets; 'send-icmp': Send ICMP Type 3 Code 4 (Fragmentation Needed and DF Set) (default); 'send-icmpv6': Send ICMP Type 2 Code 0 (Packet Too Big);
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

