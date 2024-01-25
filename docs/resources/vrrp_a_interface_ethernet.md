---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_vrrp_a_interface_ethernet Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_vrrp_a_interface_ethernet: VRRP-A interface ethernet
  PLACEHOLDER
---

# thunder_vrrp_a_interface_ethernet (Resource)

`thunder_vrrp_a_interface_ethernet`: VRRP-A interface ethernet

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_interface_ethernet" "thunder_vrrp_a_interface_ethernet" {
  both         = 1
  ethernet_val = 2
  no_heartbeat = 1
  user_tag     = "test"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `ethernet_val` (Number) Ethernet Interface

### Optional

- `both` (Number) both a router and server interface
- `no_heartbeat` (Number) do not send out heartbeat packet from this interface
- `router_interface` (Number) interface to upstream router
- `server_interface` (Number) interface to real server
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object
- `vlan_cfg` (Block List) (see [below for nested schema](#nestedblock--vlan_cfg))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--vlan_cfg"></a>
### Nested Schema for `vlan_cfg`

Optional:

- `vlan` (Number) VLAN ID

