---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_nat_shared_pool_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_nat_shared_pool_oper: Operational Status for the object shared-pool
  PLACEHOLDER
---

# thunder_cgnv6_nat_shared_pool_oper (Data Source)

`thunder_cgnv6_nat_shared_pool_oper`: Operational Status for the object shared-pool

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_nat_shared_pool_oper" "thunder_cgnv6_nat_shared_pool_oper" {
}
output "get_cgnv6_nat_shared_pool_oper" {
  value = ["${data.thunder_cgnv6_nat_shared_pool_oper.thunder_cgnv6_nat_shared_pool_oper}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--oper))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--oper"></a>
### Nested Schema for `oper`

Optional:

- `shared_pool_list` (Block List) (see [below for nested schema](#nestedblock--oper--shared_pool_list))

<a id="nestedblock--oper--shared_pool_list"></a>
### Nested Schema for `oper.shared_pool_list`

Optional:

- `end_address` (String)
- `netmask` (String)
- `pool_name` (String)
- `start_address` (String)
- `vird` (Number)

