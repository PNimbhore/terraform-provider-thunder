---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_vrrp_a_force_self_standby_persistent Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_vrrp_a_force_self_standby_persistent: HA VRRP-A Configured  Command to force the unit or a group to HA standby state
  PLACEHOLDER
---

# thunder_vrrp_a_force_self_standby_persistent (Resource)

`thunder_vrrp_a_force_self_standby_persistent`: HA VRRP-A Configured  Command to force the unit or a group to HA standby state

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_force_self_standby_persistent" "thunder_vrrp_a_force_self_standby_persistent" {

  vrid = 3
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `vrid` (Number) Specify one VRRP-A vrid to force into standby state

### Optional

- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

