---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_network_mac_address_dynamic_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_network_mac_address_dynamic_oper: Operational Status for the object dynamic
  PLACEHOLDER
---

# thunder_network_mac_address_dynamic_oper (Data Source)

`thunder_network_mac_address_dynamic_oper`: Operational Status for the object dynamic

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_network_mac_address_dynamic_oper" "thunder_network_mac_address_dynamic_oper" {
}
output "get_network_mac_address_dynamic_oper" {
  value = ["${data.thunder_network_mac_address_dynamic_oper.thunder_network_mac_address_dynamic_oper}"]
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

- `age_time` (Number)
- `macoper` (Block List) (see [below for nested schema](#nestedblock--oper--macoper))

<a id="nestedblock--oper--macoper"></a>
### Nested Schema for `oper.macoper`

Optional:

- `age` (Number)
- `index` (Number)
- `mac_address` (String)
- `port` (Number)
- `type` (String)
- `vlan` (Number)

