---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_vcs_vcs_para Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_vcs_vcs_para: Virtual Chassis System Paramter
  PLACEHOLDER
---

# thunder_vcs_vcs_para (Resource)

`thunder_vcs_vcs_para`: Virtual Chassis System Paramter

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vcs_vcs_para" "thunder_vcs_vcs_para" {
  chassis_id                = 8
  config_info               = "93"
  config_seq                = "59"
  dead_interval             = 10
  failure_retry_count_value = 2
  force_wait_interval       = 5
  memory_stat_interval      = 30
  multicast_ip              = "224.0.1.210"
  multicast_port            = 41217
  slog_level                = 7
  slog_method               = 1
  ssl_enable                = 1
  tcp_channel_monitor       = 1
  time_interval             = 3
  transmit_fragment_size    = 6000
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `chassis_id` (Number) Chassis ID
- `config_info` (String) Configuration information (Configuration tag)
- `config_seq` (String) Configuration sequence number
- `dead_interval` (Number) The node will be considered dead as lack of hearbeats after this time (in unit of second, 10 by default)
- `dead_interval_mseconds` (Number) The node will be considered dead as lack of hearbeats after this time (milisecond) (in unit of msecond, default is 0)
- `failure_retry_count_value` (Number) 0-255, default is 2
- `floating_ip_cfg` (Block List) (see [below for nested schema](#nestedblock--floating_ip_cfg))
- `floating_ipv6_cfg` (Block List) (see [below for nested schema](#nestedblock--floating_ipv6_cfg))
- `force_wait_interval` (Number) The node will wait the specified time interval before it start aVCS (in unit of second (default is 5))
- `forever` (Number) VCS retry forever if fails to join the chassis
- `memory_stat_interval` (Number) Interval of aVCS memory statistics record (minutes)
- `multicast_ip` (String) Multicast (group) IP address (Multicast IP address)
- `multicast_ipv6` (String) Multicast (group) IPv6 address (Multicast IPv6 address)
- `multicast_port` (Number) Port used in multicast communication (Port number)
- `size` (Number) file size (MBytes) to transmit to monitor the TCP channel
- `slog_level` (Number) Set the level of slog for aVCS
- `slog_method` (Number) Set the print method of slog for aVCS
- `speed_limit` (Number) speed (KByte/s) limitation for the transmit monitor
- `ssl_enable` (Number) Enable SSL
- `tcp_channel_monitor` (Number) Enable vBlade TCP channel monitor
- `time_interval` (Number) how long between heartbeats (in unit of second, default is 3)
- `time_interval_mseconds` (Number) how long between heartbeats (mseconds) (in unit of milisecond, default is 0)
- `transmit_fragment_size` (Number) Set the fragment size (KByte) of the aVCS transmit
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--floating_ip_cfg"></a>
### Nested Schema for `floating_ip_cfg`

Optional:

- `floating_ip` (String) Floating IP address (IPv4 address)
- `floating_ip_mask` (String) Netmask


<a id="nestedblock--floating_ipv6_cfg"></a>
### Nested Schema for `floating_ipv6_cfg`

Optional:

- `floating_ipv6` (String) Floating IPv6 address

