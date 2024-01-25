---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_interface_lif_ip_ospf_ospf_ip Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_interface_lif_ip_ospf_ospf_ip: IP address configuration for Open Shortest Path First for IPv4 (OSPF)
  PLACEHOLDER
---

# thunder_interface_lif_ip_ospf_ospf_ip (Resource)

`thunder_interface_lif_ip_ospf_ospf_ip`: IP address configuration for Open Shortest Path First for IPv4 (OSPF)

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_lif_ip_ospf_ospf_ip" "thunder_interface_lif_ip_ospf_ospf_ip" {

  ifname             = "test"
  authentication     = 1
  authentication_key = "1"
  cost               = 64948
  database_filter    = "all"
  dead_interval      = 40
  hello_interval     = 10
  ip_addr            = "10.10.10.10"
  message_digest_cfg {
    message_digest_key = 7
    md5_value          = "10"
  }
  mtu_ignore          = 1
  out                 = 1
  priority            = 1
  retransmit_interval = 115
  transmit_delay      = 10
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `ifname` (String) Ifname
- `ip_addr` (String) Address of interface

### Optional

- `authentication` (Number) Enable authentication
- `authentication_key` (String) Authentication password (key) (The OSPF password (key))
- `cost` (Number) Interface cost
- `database_filter` (String) 'all': Filter all LSA;
- `dead_interval` (Number) Interval after which a neighbor is declared dead (Seconds)
- `hello_interval` (Number) Time between HELLO packets (Seconds)
- `message_digest_cfg` (Block List) (see [below for nested schema](#nestedblock--message_digest_cfg))
- `mtu_ignore` (Number) Ignores the MTU in DBD packets
- `out` (Number) Outgoing LSA
- `priority` (Number) Router priority
- `retransmit_interval` (Number) Time between retransmitting lost link state advertisements (Seconds)
- `transmit_delay` (Number) Link state transmit delay (Seconds)
- `uuid` (String) uuid of the object
- `value` (String) 'message-digest': Use message-digest authentication; 'null': Use no authentication;

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--message_digest_cfg"></a>
### Nested Schema for `message_digest_cfg`

Optional:

- `md5_value` (String) The OSPF password (1-16)
- `message_digest_key` (Number) Message digest authentication password (key) (Key id)

