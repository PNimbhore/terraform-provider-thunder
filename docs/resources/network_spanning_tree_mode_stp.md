---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_network_spanning_tree_mode_stp Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_network_spanning_tree_mode_stp: Configure spanning tree protocol STP
  PLACEHOLDER
---

# thunder_network_spanning_tree_mode_stp (Resource)

`thunder_network_spanning_tree_mode_stp`: Configure spanning tree protocol STP

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_spanning_tree_mode_stp" "thunder_network_spanning_tree_mode_stp" {
  action     = 1
  priority   = 32768
  vlan_start = 426
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `action` (Number) enable spanning tree RSTP mode
- `priority` (Number) set bridge priority
- `uuid` (String) uuid of the object
- `vlan_start` (Number) VLAN ID

### Read-Only

- `id` (String) The ID of this resource.


