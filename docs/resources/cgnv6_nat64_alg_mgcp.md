---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_nat64_alg_mgcp Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_nat64_alg_mgcp: NAT64 MGCP ALG (default: disabled)
  PLACEHOLDER
---

# thunder_cgnv6_nat64_alg_mgcp (Resource)

`thunder_cgnv6_nat64_alg_mgcp`: NAT64 MGCP ALG (default: disabled)

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat64_alg_mgcp" "thunder_cgnv6_nat64_alg_mgcp" {
  mgcp_enable = "enable"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `mgcp_enable` (String) 'enable': Enable NAT64 MGCP ALG;
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


