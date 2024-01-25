---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ipv6_fib_summary_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ipv6_fib_summary_oper: Operational Status for the object fib-summary
  PLACEHOLDER
---

# thunder_ipv6_fib_summary_oper (Data Source)

`thunder_ipv6_fib_summary_oper`: Operational Status for the object fib-summary

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ipv6_fib_summary_oper" "thunder_ipv6_fib_summary_oper" {
}
output "get_ipv6_fib_summary_oper" {
  value = ["${data.thunder_ipv6_fib_summary_oper.thunder_ipv6_fib_summary_oper}"]
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

- `connected_routes` (Number)
- `static_dynamic_paths` (Number)
- `static_dynamic_routes` (Number)
- `total_paths` (Number)
- `total_routes` (Number)

