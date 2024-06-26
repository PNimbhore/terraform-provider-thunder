---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_debug_ipv6_rip Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_debug_ipv6_rip: Routing Information Protocol (RIP) for IPv6
  PLACEHOLDER
---

# thunder_debug_ipv6_rip (Resource)

`thunder_debug_ipv6_rip`: Routing Information Protocol (RIP) for IPv6

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_ipv6_rip" "thunder_debug_ipv6_rip" {
  all    = 0
  detail = 0
  events = 0
  nsm    = 0
  packet = 0
  recv   = 0
  send   = 0
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `all` (Number) Enable all debugging
- `detail` (Number) Detailed information display
- `events` (Number) RIPng events
- `nsm` (Number) RIPng and NSM communication
- `packet` (Number) RIPng packets
- `recv` (Number) RIPng receive packets
- `send` (Number) RIPng send packets

### Read-Only

- `id` (String) The ID of this resource.


