---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_ddos_protection_ip_entries_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_ddos_protection_ip_entries_oper: Operational Status for the object ip-entries
  PLACEHOLDER
---

# thunder_cgnv6_ddos_protection_ip_entries_oper (Data Source)

`thunder_cgnv6_ddos_protection_ip_entries_oper`: Operational Status for the object ip-entries

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_ddos_protection_ip_entries_oper" "thunder_cgnv6_ddos_protection_ip_entries_oper" {

}
output "get_cgnv6_ddos_protection_ip_entries_oper" {
  value = ["${data.thunder_cgnv6_ddos_protection_ip_entries_oper.thunder_cgnv6_ddos_protection_ip_entries_oper}"]
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

- `all` (Number)
- `ddos_ip_entries_list` (Block List) (see [below for nested schema](#nestedblock--oper--ddos_ip_entries_list))
- `nat_pool` (String)
- `total_entries` (Number)
- `v4_netmask` (String)

<a id="nestedblock--oper--ddos_ip_entries_list"></a>
### Nested Schema for `oper.ddos_ip_entries_list`

Optional:

- `expiration` (Number)
- `hardware_index` (Number)
- `hints` (String)
- `hw_l3_drop_pps` (Number)
- `hw_l4_drop_pps` (Number)
- `in_blacklist` (Number)
- `in_hardware` (Number)
- `l4_entries_count` (Number)
- `sw_l3_drop_pps` (Number)
- `sw_l4_drop_pps` (Number)
- `sw_receive_pps` (Number)
- `total_pps` (Number)
- `v4_address` (String)

