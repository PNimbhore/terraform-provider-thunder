---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_fixed_nat_port_mapping_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_fixed_nat_port_mapping_oper: Operational Status for the object port-mapping
  PLACEHOLDER
---

# thunder_cgnv6_fixed_nat_port_mapping_oper (Data Source)

`thunder_cgnv6_fixed_nat_port_mapping_oper`: Operational Status for the object port-mapping

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_fixed_nat_port_mapping_oper" "thunder_cgnv6_fixed_nat_port_mapping_oper" {

}
output "get_cgnv6_fixed_nat_port_mapping_oper" {
  value = ["${data.thunder_cgnv6_fixed_nat_port_mapping_oper.thunder_cgnv6_fixed_nat_port_mapping_oper}"]
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

- `inside_user_v4` (String)
- `inside_user_v6` (String)
- `mapping_list` (Block List) (see [below for nested schema](#nestedblock--oper--mapping_list))
- `nat_ip` (String)
- `nat_port` (Number)
- `partition` (String)

<a id="nestedblock--oper--mapping_list"></a>
### Nested Schema for `oper.mapping_list`

Optional:

- `assigned_to` (String)
- `icmp_port_end` (Number)
- `icmp_port_start` (Number)
- `nat_address` (String)
- `tcp_port_end` (Number)
- `tcp_port_start` (Number)
- `udp_port_end` (Number)
- `udp_port_start` (Number)

