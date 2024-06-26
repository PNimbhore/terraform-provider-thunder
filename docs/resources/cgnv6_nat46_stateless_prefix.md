---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_nat46_stateless_prefix Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_nat46_stateless_prefix: Enable Stateless NAT46 IPv6 source address
  PLACEHOLDER
---

# thunder_cgnv6_nat46_stateless_prefix (Resource)

`thunder_cgnv6_nat46_stateless_prefix`: Enable Stateless NAT46 IPv6 source address

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat46_stateless_prefix" "thunder_cgnv6_nat46_stateless_prefix" {

  ipv6_prefix = "46::/96"
  vrid        = 3

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `ipv6_prefix` (String) IPv6 prefix
- `uuid` (String) uuid of the object
- `vrid` (Number) VRRP-A vrid (Specify ha VRRP-A vrid)

### Read-Only

- `id` (String) The ID of this resource.


