---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_vrrp_a_common Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_vrrp_a_common: HA VRRP-A Global Commands
  PLACEHOLDER
---

# thunder_vrrp_a_common (Resource)

`thunder_vrrp_a_common`: HA VRRP-A Global Commands

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_common" "thunder_vrrp_a_common" {
  device_id            = 3
  set_id               = 3
  disable_default_vrid = 1
  action               = "enable"
  hello_interval       = 3
  preemption_delay     = 222
  dead_timer           = 20
  arp_retry            = 3
  ttl                  = 120
  hop_limit            = 122
  track_event_delay    = 33
  get_ready_time       = 120
  inline_mode_cfg {
    inline_mode = 1
  }
  restart_time = 99
  hostid_append_to_vrid {
    hostid_append_to_vrid_default = 1
  }
  forward_l4_packet_on_standby = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `action` (String) 'enable': enable vrrp-a; 'disable': disable vrrp-a;
- `arp_retry` (Number) Number of additional gratuitous ARPs sent out after HA failover (1-255, default is 4)
- `dead_timer` (Number) VRRP-A dead timer in terms of how many hello messages missed, default is 5 (2-255, default is 5)
- `device_id` (Number) Unique ID for each VRRP-A box (Device-id number)
- `disable_default_vrid` (Number) Disable default vrid
- `forward_l4_packet_on_standby` (Number) Enables Layer 2/3 forwarding of Layer 4 traffic on the Standby ACOS device
- `get_ready_time` (Number) set get ready time after ax starting up (60-1200, in unit of 100millisec, default is 60)
- `hello_interval` (Number) VRRP-A Hello Interval (1-255, in unit of 100millisec, default is 2)
- `hop_limit` (Number) VRRP-A packet IPv6 header hop-limit (hop-limit, default is 64)
- `hostid_append_to_vrid` (Block List, Max: 1) (see [below for nested schema](#nestedblock--hostid_append_to_vrid))
- `inline_mode_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--inline_mode_cfg))
- `preemption_delay` (Number) Delay before changing state from Active to Standby (1-255, in unit of 100millisec, default is 60)
- `restart_time` (Number) Time between restarting ports on standby system after transition
- `set_id` (Number) Set-ID for HA configuration (Set id from 1 to 15)
- `track_event_delay` (Number) Delay before changing state after up/down event (Units of 100 milliseconds (default 30))
- `ttl` (Number) VRRP-A packet IPv4 header TTL (TTL, default is 128)
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--hostid_append_to_vrid"></a>
### Nested Schema for `hostid_append_to_vrid`

Optional:

- `hostid_append_to_vrid_default` (Number) hostid append to vrid default
- `hostid_append_to_vrid_value` (Number) hostid append to vrid num


<a id="nestedblock--inline_mode_cfg"></a>
### Nested Schema for `inline_mode_cfg`

Optional:

- `inline_mode` (Number) Enable Layer 2 Inline Hot Standby Mode
- `preferred_port` (Number) Preferred ethernet Port
- `preferred_trunk` (Number) Preferred trunk Port


