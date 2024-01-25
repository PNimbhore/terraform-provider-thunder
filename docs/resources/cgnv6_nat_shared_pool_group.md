---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_nat_shared_pool_group Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_nat_shared_pool_group:
  PLACEHOLDER
---

# thunder_cgnv6_nat_shared_pool_group (Resource)

`thunder_cgnv6_nat_shared_pool_group`: 

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat_shared_pool_group" "thunder_cgnv6_nat_shared_pool_group" {
  members {
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `members` (Block List, Max: 1) (see [below for nested schema](#nestedblock--members))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--members"></a>
### Nested Schema for `members`

Optional:

- `uuid` (String) uuid of the object

