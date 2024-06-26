---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_health_stat_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_health_stat_stats: Statistics for the object health-stat
  PLACEHOLDER
---

# thunder_slb_health_stat_stats (Data Source)

`thunder_slb_health_stat_stats`: Statistics for the object health-stat

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_health_stat_stats" "thunder_slb_health_stat_stats" {
}
output "get_slb_health_stat_stats" {
  value = ["${data.thunder_slb_health_stat_stats.thunder_slb_health_stat_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `avg_jiffie` (Number) Average number of jiffies
- `close_socket` (Number) Number of closed sockets
- `config_health_rate` (Number) Config health rate
- `conn_imdt_succ` (Number) Number of connection immediete success
- `connect_failed` (Number) Number of failed connections
- `curr_health_rate` (Number) Current health rate
- `ext_health_rate` (Number) External health rate
- `ext_health_rate_val` (Number) External health rate value
- `max_jiffie` (Number) Maximum number of jiffies
- `min_jiffie` (Number) Minimum number of jiffies
- `num_burst` (Number) Number of burst
- `open_socket` (Number) Number of open sockets
- `open_socket_failed` (Number) Number of failed open sockets
- `recv_packet` (Number) Number of received packets
- `recv_packet_failed` (Number) Number of failed packet receives
- `retry_times` (Number) Retry times
- `running_time` (Number) Running time
- `send_packet` (Number) Number of packets sent
- `send_packet_failed` (Number) Number of packet send failures
- `sock_close_before_17` (Number) Number of sockets closed before l7
- `sock_close_without_notify` (Number) Number of sockets closed without notify
- `ssl_post_handshake_packet` (Number) Number of ssl post handshake packets before client sends request
- `status_down` (Number) Number of status downs
- `status_other` (Number) Number of other status
- `status_unkn` (Number) Number of status unknowns
- `status_up` (Number) Number of status ups
- `timeout` (Number) Timouet value
- `timeout_with_packet` (Number) Number of pin timeouts while socket has packets
- `total_number` (Number) Total number
- `unexpected_error` (Number) Number of unexpected errors


