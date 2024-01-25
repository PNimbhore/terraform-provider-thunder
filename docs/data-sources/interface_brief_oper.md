---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_interface_brief_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_interface_brief_oper: Operational Status for the object brief
  PLACEHOLDER
---

# thunder_interface_brief_oper (Data Source)

`thunder_interface_brief_oper`: Operational Status for the object brief

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_interface_brief_oper" "thunder_interface_brief_oper" {
}
output "get_interface_brief_oper" {
  value = ["${data.thunder_interface_brief_oper.thunder_interface_brief_oper}"]
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

- `interfaces` (Block List) (see [below for nested schema](#nestedblock--oper--interfaces))

<a id="nestedblock--oper--interfaces"></a>
### Nested Schema for `oper.interfaces`

Optional:

- `duplexity` (String)
- `encap_info` (String)
- `intf_name` (String)
- `ipv4_addr` (String) IP address
- `ipv4_addr_count` (Number)
- `ipv4_mask` (String) IP subnet mask
- `ipv6_addr` (String) IP address
- `ipv6_addr_count` (Number)
- `ipv6_prefix` (String) IP subnet mask
- `mac` (String)
- `port_num` (String)
- `speed` (String)
- `state` (String)
- `trunk_group` (String)
- `unnumbered_oper` (String)
- `vlan_info` (String)

