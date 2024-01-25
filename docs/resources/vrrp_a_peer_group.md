---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_vrrp_a_peer_group Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_vrrp_a_peer_group: VRRP-A peer group
  PLACEHOLDER
---

# thunder_vrrp_a_peer_group (Resource)

`thunder_vrrp_a_peer_group`: VRRP-A peer group

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_peer_group" "thunder_vrrp_a_peer_group" {
  peer {
    ip_peer_address_cfg {
      ip_peer_address = "10.10.10.10"
    }
    ipv6_peer_address_cfg {
      ipv6_peer_address = "2001:db8:3333:4444:5555:6666:7777:8888"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `peer` (Block List, Max: 1) (see [below for nested schema](#nestedblock--peer))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--peer"></a>
### Nested Schema for `peer`

Optional:

- `ip_peer_address_cfg` (Block List) (see [below for nested schema](#nestedblock--peer--ip_peer_address_cfg))
- `ipv6_peer_address_cfg` (Block List) (see [below for nested schema](#nestedblock--peer--ipv6_peer_address_cfg))

<a id="nestedblock--peer--ip_peer_address_cfg"></a>
### Nested Schema for `peer.ip_peer_address_cfg`

Optional:

- `ip_peer_address` (String) IP Address


<a id="nestedblock--peer--ipv6_peer_address_cfg"></a>
### Nested Schema for `peer.ipv6_peer_address_cfg`

Optional:

- `ipv6_peer_address` (String) IPV6 address

