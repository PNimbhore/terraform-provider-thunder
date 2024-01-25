---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_nat_icmpv6 Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_nat_icmpv6: ICMPv6 configuration for IPv6 NAT
  PLACEHOLDER
---

# thunder_cgnv6_nat_icmpv6 (Resource)

`thunder_cgnv6_nat_icmpv6`: ICMPv6 configuration for IPv6 NAT

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat_icmpv6" "thunder_cgnv6_nat_icmpv6" {
  respond_to_ping = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `respond_to_ping` (Number) Respond to ICMPv6 echo requests to NAT pool IPs (default: disabled)
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

