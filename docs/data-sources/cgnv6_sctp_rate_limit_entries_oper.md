---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_sctp_rate_limit_entries_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_sctp_rate_limit_entries_oper: Operational Status for the object rate-limit-entries
  PLACEHOLDER
---

# thunder_cgnv6_sctp_rate_limit_entries_oper (Data Source)

`thunder_cgnv6_sctp_rate_limit_entries_oper`: Operational Status for the object rate-limit-entries

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_sctp_rate_limit_entries_oper" "thunder_cgnv6_sctp_rate_limit_entries_oper" {
}
output "get_cgnv6_sctp_rate_limit_entries_oper" {
  value = ["${data.thunder_cgnv6_sctp_rate_limit_entries_oper.thunder_cgnv6_sctp_rate_limit_entries_oper}"]
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

- `entry_count` (Number)
- `rate_limit_entries_list` (Block List) (see [below for nested schema](#nestedblock--oper--rate_limit_entries_list))

<a id="nestedblock--oper--rate_limit_entries_list"></a>
### Nested Schema for `oper.rate_limit_entries_list`

Optional:

- `address` (String)
- `direction` (String)
- `pps` (Number)
- `rate_limit` (Number)

